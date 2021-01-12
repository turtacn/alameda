package influxdb

import (
	RepoInfluxPlanning "github.com/turtacn/alameda/datahub/pkg/dao/repositories/influxdb/plannings"
	InternalInflux "github.com/turtacn/alameda/internal/pkg/database/influxdb"
	Log "github.com/turtacn/alameda/pkg/utils/log"
	ApiPlannings "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/plannings"
)

var (
	scope = Log.RegisterScope("planning_dao_implement", "planning dao implement", 0)
)

// Container Implements ContainerOperation interface
type ContainerPlannings struct {
	InfluxDBConfig InternalInflux.Config
}

func NewContainerPlanningsWithConfig(config InternalInflux.Config) *ContainerPlannings {
	return &ContainerPlannings{InfluxDBConfig: config}
}

// AddPodPlannings add pod plannings to database
func (c *ContainerPlannings) AddPodPlannings(in *ApiPlannings.CreatePodPlanningsRequest) error {
	containerRepository := RepoInfluxPlanning.NewContainerRepository(&c.InfluxDBConfig)
	return containerRepository.CreateContainerPlannings(in)
}

// ListPodPlannings list pod plannings
func (c *ContainerPlannings) ListPodPlannings(in *ApiPlannings.ListPodPlanningsRequest) ([]*ApiPlannings.PodPlanning, error) {
	containerRepository := RepoInfluxPlanning.NewContainerRepository(&c.InfluxDBConfig)
	return containerRepository.ListContainerPlannings(in)
}
