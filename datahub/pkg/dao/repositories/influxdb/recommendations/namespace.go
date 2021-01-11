package recommendations

import (
	EntityInfluxRecommend "github.com/turtacn/alameda/datahub/pkg/dao/entities/influxdb/recommendations"
	RepoInflux "github.com/turtacn/alameda/datahub/pkg/dao/repositories/influxdb"
	DBCommon "github.com/turtacn/alameda/internal/pkg/database/common"
	InternalInflux "github.com/turtacn/alameda/internal/pkg/database/influxdb"
	ApiRecommendations "github.com/turtacn/api/alameda_api/v1alpha1/datahub/recommendations"
	ApiResources "github.com/turtacn/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"strconv"
	"time"
)

type NamespaceRepository struct {
	influxDB *InternalInflux.InfluxClient
}

func NewNamespaceRepository(influxDBCfg *InternalInflux.Config) *NamespaceRepository {
	return &NamespaceRepository{
		influxDB: &InternalInflux.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

func (c *NamespaceRepository) CreateRecommendations(recommendations []*ApiRecommendations.NamespaceRecommendation) error {
	points := make([]*InfluxClient.Point, 0)
	for _, recommendation := range recommendations {
		recommendedType := recommendation.GetRecommendedType()

		if recommendedType == ApiRecommendations.ControllerRecommendedType_PRIMITIVE {
			recommendedSpec := recommendation.GetRecommendedSpec()

			tags := map[string]string{
				EntityInfluxRecommend.NamespaceClusterName: recommendation.GetObjectMeta().GetClusterName(),
				EntityInfluxRecommend.NamespaceName:        recommendation.GetObjectMeta().GetName(),
				EntityInfluxRecommend.NamespaceType:        ApiRecommendations.ControllerRecommendedType_PRIMITIVE.String(),
			}

			fields := map[string]interface{}{
				EntityInfluxRecommend.NamespaceCurrentReplicas: recommendedSpec.GetCurrentReplicas(),
				EntityInfluxRecommend.NamespaceDesiredReplicas: recommendedSpec.GetDesiredReplicas(),
				EntityInfluxRecommend.NamespaceCreateTime:      recommendedSpec.GetCreateTime().GetSeconds(),
				EntityInfluxRecommend.NamespaceKind:            recommendation.GetKind().String(),

				EntityInfluxRecommend.NamespaceCurrentCPURequest: recommendedSpec.GetCurrentCpuRequests(),
				EntityInfluxRecommend.NamespaceCurrentMEMRequest: recommendedSpec.GetCurrentMemRequests(),
				EntityInfluxRecommend.NamespaceCurrentCPULimit:   recommendedSpec.GetCurrentCpuLimits(),
				EntityInfluxRecommend.NamespaceCurrentMEMLimit:   recommendedSpec.GetCurrentMemLimits(),
				EntityInfluxRecommend.NamespaceDesiredCPULimit:   recommendedSpec.GetDesiredCpuLimits(),
				EntityInfluxRecommend.NamespaceDesiredMEMLimit:   recommendedSpec.GetDesiredMemLimits(),
				EntityInfluxRecommend.NamespaceTotalCost:         recommendedSpec.GetTotalCost(),
			}

			pt, err := InfluxClient.NewPoint(string(Namespace), tags, fields, time.Unix(recommendedSpec.GetTime().GetSeconds(), 0))
			if err != nil {
				scope.Error(err.Error())
			}

			points = append(points, pt)

		} else if recommendedType == ApiRecommendations.ControllerRecommendedType_K8S {
			recommendedSpec := recommendation.GetRecommendedSpecK8S()

			tags := map[string]string{
				EntityInfluxRecommend.NamespaceName: recommendation.GetObjectMeta().GetName(),
				EntityInfluxRecommend.NamespaceType: ApiRecommendations.ControllerRecommendedType_K8S.String(),
			}

			fields := map[string]interface{}{
				EntityInfluxRecommend.NamespaceCurrentReplicas: recommendedSpec.GetCurrentReplicas(),
				EntityInfluxRecommend.NamespaceDesiredReplicas: recommendedSpec.GetDesiredReplicas(),
				EntityInfluxRecommend.NamespaceCreateTime:      recommendedSpec.GetCreateTime().GetSeconds(),
				EntityInfluxRecommend.NamespaceKind:            recommendation.GetKind().String(),
			}

			pt, err := InfluxClient.NewPoint(string(Namespace), tags, fields, time.Unix(recommendedSpec.GetTime().GetSeconds(), 0))
			if err != nil {
				scope.Error(err.Error())
			}

			points = append(points, pt)
		}
	}

	err := c.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.Recommendation),
	})

	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (c *NamespaceRepository) ListRecommendations(in *ApiRecommendations.ListNamespaceRecommendationsRequest) ([]*ApiRecommendations.NamespaceRecommendation, error) {
	influxdbStatement := InternalInflux.Statement{
		Measurement:    Namespace,
		QueryCondition: DBCommon.BuildQueryConditionV1(in.GetQueryCondition()),
	}

	recommendationType := in.GetRecommendedType().String()
	kind := in.GetKind().String()

	for _, objMeta := range in.GetObjectMeta() {
		cluster := objMeta.GetClusterName()
		name := objMeta.GetName()

		keyList := []string{
			EntityInfluxRecommend.NamespaceClusterName,
			EntityInfluxRecommend.NamespaceName,
		}
		valueList := []string{
			cluster,
			name,
		}

		if recommendationType != ApiRecommendations.ControllerRecommendedType_CRT_UNDEFINED.String() {
			keyList = append(keyList, EntityInfluxRecommend.NamespaceType)
			valueList = append(valueList, recommendationType)
		}

		if kind != ApiResources.Kind_KIND_UNDEFINED.String() {
			keyList = append(keyList, EntityInfluxRecommend.NamespaceKind)
			valueList = append(valueList, kind)
		}

		tempCondition := influxdbStatement.GenerateCondition(keyList, valueList, "AND")
		influxdbStatement.AppendWhereClauseDirectly("OR", tempCondition)
	}

	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	results, err := c.influxDB.QueryDB(cmd, string(RepoInflux.Recommendation))
	if err != nil {
		return make([]*ApiRecommendations.NamespaceRecommendation, 0), err
	}

	influxdbRows := InternalInflux.PackMap(results)
	recommendations := c.getRecommendationsFromInfluxRows(influxdbRows)

	return recommendations, nil
}

func (c *NamespaceRepository) getRecommendationsFromInfluxRows(rows []*InternalInflux.InfluxRow) []*ApiRecommendations.NamespaceRecommendation {
	recommendations := make([]*ApiRecommendations.NamespaceRecommendation, 0)
	for _, influxdbRow := range rows {
		for _, data := range influxdbRow.Data {
			currentReplicas, _ := strconv.ParseInt(data[EntityInfluxRecommend.NamespaceCurrentReplicas], 10, 64)
			desiredReplicas, _ := strconv.ParseInt(data[EntityInfluxRecommend.NamespaceDesiredReplicas], 10, 64)
			createTime, _ := strconv.ParseInt(data[EntityInfluxRecommend.NamespaceCreateTime], 10, 64)

			t, _ := time.Parse(time.RFC3339, data[EntityInfluxRecommend.NamespaceTime])
			tempTime, _ := ptypes.TimestampProto(t)

			currentCpuRequests, _ := strconv.ParseFloat(data[EntityInfluxRecommend.NamespaceCurrentCPURequest], 64)
			currentMemRequests, _ := strconv.ParseFloat(data[EntityInfluxRecommend.NamespaceCurrentMEMRequest], 64)
			currentCpuLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.NamespaceCurrentCPULimit], 64)
			currentMemLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.NamespaceCurrentMEMLimit], 64)
			desiredCpuLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.NamespaceDesiredCPULimit], 64)
			desiredMemLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.NamespaceDesiredMEMLimit], 64)
			totalCost, _ := strconv.ParseFloat(data[EntityInfluxRecommend.NamespaceTotalCost], 64)

			var commendationType ApiRecommendations.ControllerRecommendedType
			if tempType, exist := data[EntityInfluxRecommend.NamespaceType]; exist {
				if value, ok := ApiRecommendations.ControllerRecommendedType_value[tempType]; ok {
					commendationType = ApiRecommendations.ControllerRecommendedType(value)
				}
			}

			var commendationKind ApiResources.Kind
			if tempKind, exist := data[EntityInfluxRecommend.NamespaceKind]; exist {
				if value, ok := ApiResources.Kind_value[tempKind]; ok {
					commendationKind = ApiResources.Kind(value)
				}
			}

			if commendationType == ApiRecommendations.ControllerRecommendedType_PRIMITIVE {
				tempRecommendation := &ApiRecommendations.NamespaceRecommendation{
					ObjectMeta: &ApiResources.ObjectMeta{
						ClusterName: data[string(EntityInfluxRecommend.NamespaceClusterName)],
						Name:        data[string(EntityInfluxRecommend.NamespaceName)],
					},
					Kind:            commendationKind,
					RecommendedType: commendationType,
					RecommendedSpec: &ApiRecommendations.ControllerRecommendedSpec{
						CurrentReplicas: int32(currentReplicas),
						DesiredReplicas: int32(desiredReplicas),
						Time:            tempTime,
						CreateTime: &timestamp.Timestamp{
							Seconds: createTime,
						},
						CurrentCpuRequests: currentCpuRequests,
						CurrentMemRequests: currentMemRequests,
						CurrentCpuLimits:   currentCpuLimits,
						CurrentMemLimits:   currentMemLimits,
						DesiredCpuLimits:   desiredCpuLimits,
						DesiredMemLimits:   desiredMemLimits,
						TotalCost:          totalCost,
					},
				}

				recommendations = append(recommendations, tempRecommendation)

			} else if commendationType == ApiRecommendations.ControllerRecommendedType_K8S {
				tempRecommendation := &ApiRecommendations.NamespaceRecommendation{
					ObjectMeta: &ApiResources.ObjectMeta{
						ClusterName: data[string(EntityInfluxRecommend.NamespaceClusterName)],
						Name:        data[string(EntityInfluxRecommend.NamespaceName)],
					},
					Kind:            commendationKind,
					RecommendedType: commendationType,
					RecommendedSpecK8S: &ApiRecommendations.ControllerRecommendedSpecK8S{
						CurrentReplicas: int32(currentReplicas),
						DesiredReplicas: int32(desiredReplicas),
						Time:            tempTime,
						CreateTime: &timestamp.Timestamp{
							Seconds: createTime,
						},
					},
				}

				recommendations = append(recommendations, tempRecommendation)
			}
		}
	}

	return recommendations
}
