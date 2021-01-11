package influxdb

import (
	//DaoRecommendationTypes "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/recommendations/types"
	RepoInfluxRecommendation "github.com/turtacn/alameda/datahub/pkg/dao/repositories/influxdb/recommendations"
	InternalInflux "github.com/turtacn/alameda/internal/pkg/database/influxdb"
	ApiRecommendations "github.com/turtacn/api/alameda_api/v1alpha1/datahub/recommendations"
)

type ClusterRecommendations struct {
	InfluxDBConfig InternalInflux.Config
}

func NewClusterRecommendationsWithConfig(config InternalInflux.Config) *ClusterRecommendations {
	return &ClusterRecommendations{InfluxDBConfig: config}
}

func (c *ClusterRecommendations) CreateRecommendations(recommendations []*ApiRecommendations.ClusterRecommendation) error {
	repository := RepoInfluxRecommendation.NewClusterRepository(&c.InfluxDBConfig)
	return repository.CreateRecommendations(recommendations)
}

func (c *ClusterRecommendations) ListRecommendations(in *ApiRecommendations.ListClusterRecommendationsRequest) ([]*ApiRecommendations.ClusterRecommendation, error) {
	repository := RepoInfluxRecommendation.NewClusterRepository(&c.InfluxDBConfig)
	return repository.ListRecommendations(in)
}
