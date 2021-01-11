package v1alpha1

import (
	DaoMetrics "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics"
	FormatRequest "github.com/turtacn/alameda/datahub/pkg/formatconversion/requests"
	FormatResponse "github.com/turtacn/alameda/datahub/pkg/formatconversion/responses"
	AlamedaUtils "github.com/turtacn/alameda/pkg/utils"
	ApiMetrics "github.com/turtacn/api/alameda_api/v1alpha1/datahub/metrics"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (s *ServiceV1alpha1) CreateControllerMetrics(ctx context.Context, in *ApiMetrics.CreateControllerMetricsRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateControllerMetrics grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExtended := FormatRequest.CreateControllerMetricsRequestExtended{CreateControllerMetricsRequest: *in}
	if err := requestExtended.Validate(); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: err.Error(),
		}, nil
	}

	metricDAO := DaoMetrics.NewControllerMetricsWriterDAO(*s.Config)
	err := metricDAO.CreateMetrics(ctx, requestExtended.ProduceMetrics())
	if err != nil {
		scope.Errorf("failed to create application metrics: %+v", err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListControllerMetrics(ctx context.Context, in *ApiMetrics.ListControllerMetricsRequest) (*ApiMetrics.ListControllerMetricsResponse, error) {
	scope.Debug("Request received from ListControllerMetrics grpc function: " + AlamedaUtils.InterfaceToString(in))

	requestExtended := FormatRequest.ListControllerMetricsRequestExtended{Request: in}
	if err := requestExtended.Validate(); err != nil {
		return &ApiMetrics.ListControllerMetricsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: err.Error(),
			},
		}, nil
	}
	requestExtended.SetDefaultWithMetricsDBType(s.Config.Apis.Metrics.Source)

	metricsDao := DaoMetrics.NewControllerMetricsReaderDAO(*s.Config)
	metricMap, err := metricsDao.ListMetrics(ctx, requestExtended.ProduceRequest())
	if err != nil {
		return &ApiMetrics.ListControllerMetricsResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}
	i := 0
	datahubControllerMetrics := make([]*ApiMetrics.ControllerMetric, len(metricMap.MetricMap))
	for _, metric := range metricMap.MetricMap {
		m := FormatResponse.ControllerMetricExtended{ControllerMetric: *metric}.ProduceMetrics()
		datahubControllerMetrics[i] = &m
		i++
	}

	return &ApiMetrics.ListControllerMetricsResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		ControllerMetrics: datahubControllerMetrics,
	}, nil
}
