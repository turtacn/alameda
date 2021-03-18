package v1alpha1

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/applications"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/metrics"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/plannings"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/predictions"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"google.golang.org/genproto/googleapis/rpc/status"

	"google.golang.org/genproto/googleapis/rpc/code"
)

func (s *ServiceV1alpha1) CreateApps(ctx context.Context, request *applications.CreateApplicationsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) ListApps(ctx context.Context, request *applications.ListApplicationsRequest) (*applications.ListApplicationsResponse, error) {
	return nil, nil
}
func (s *ServiceV1alpha1) DeleteApps(ctx context.Context, request *applications.DeleteApplicationsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) ReadData(ctx context.Context, request *data.ReadDataRequest) (*data.ReadDataResponse, error) {
	return nil, nil
}
func (s *ServiceV1alpha1) WriteData(ctx context.Context, request *data.WriteDataRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) DeleteData(ctx context.Context, request *data.DeleteDataRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) CreateMetrics(ctx context.Context, request *metrics.CreateMetricsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) ListMetrics(ctx context.Context, request *metrics.ListMetricsRequest) (*metrics.ListMetricsResponse, error) {
	return nil, nil
}
func (s *ServiceV1alpha1) CreatePlannings(ctx context.Context, request *plannings.CreatePlanningsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) ListPlannings(ctx context.Context, request *plannings.ListPlanningsRequest) (*plannings.ListPlanningsResponse, error) {
	return nil, nil
}
func (s *ServiceV1alpha1) CreatePredictions(ctx context.Context, request *predictions.CreatePredictionsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) ListPredictions(ctx context.Context, request *predictions.ListPredictionsRequest) (*predictions.ListPredictionsResponse, error) {
	return nil, nil
}
func (s *ServiceV1alpha1) CreateRecommendations(ctx context.Context, request *recommendations.CreateRecommendationsRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

func (s *ServiceV1alpha1) ListRecommendations(ctx context.Context, request *recommendations.ListRecommendationsRequest) (*recommendations.ListRecommendationsResponse, error) {
	return nil, nil
}
func (s *ServiceV1alpha1) CreateSchemas(ctx context.Context, request *schemas.CreateSchemasRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
func (s *ServiceV1alpha1) ListSchemas(ctx context.Context, request *schemas.ListSchemasRequest) (*schemas.ListSchemasResponse, error) {
	return nil, nil
}
func (s *ServiceV1alpha1) DeleteSchemas(ctx context.Context, request *schemas.DeleteSchemasRequest) (*status.Status, error) {
	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
