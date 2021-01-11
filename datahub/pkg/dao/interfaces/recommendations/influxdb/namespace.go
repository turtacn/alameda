package influxdb

import (
	//DaoRecommendationTypes "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/recommendations/types"
	RepoInfluxRecommendation "github.com/turtacn/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InternalInflux "github.com/turtacn/alameda/internal/pkg/database/influxdb"
	ApiRecommendations "github.com/turtacn/api/alameda_api/v1alpha1/datahub/recommendations"
)

type NamespaceRecommendations struct {
	InfluxDBConfig InternalInflux.Config
}

func NewNamespaceRecommendationsWithConfig(config InternalInflux.Config) *NamespaceRecommendations {
	return &NamespaceRecommendations{InfluxDBConfig: config}
}

func (c *NamespaceRecommendations) CreateRecommendations(recommendations []*ApiRecommendations.NamespaceRecommendation) error {
	repository := RepoInfluxRecommendation.NewNamespaceRepository(&c.InfluxDBConfig)
	return repository.CreateRecommendations(recommendations)
}

func (c *NamespaceRecommendations) ListRecommendations(in *ApiRecommendations.ListNamespaceRecommendationsRequest) ([]*ApiRecommendations.NamespaceRecommendation, error) {
	repository := RepoInfluxRecommendation.NewNamespaceRepository(&c.InfluxDBConfig)
	return repository.ListRecommendations(in)
}
