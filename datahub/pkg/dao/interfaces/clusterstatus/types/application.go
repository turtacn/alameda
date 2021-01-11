package types

import (
	"github.com/turtacn/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	"github.com/turtacn/alameda/datahub/pkg/kubernetes/metadata"
	"github.com/turtacn/alameda/internal/pkg/database/common"
	"github.com/turtacn/alameda/internal/pkg/database/influxdb"
)

type ApplicationDAO interface {
	CreateApplications([]*Application) error
	ListApplications(*ListApplicationsRequest) ([]*Application, error)
	DeleteApplications(*DeleteApplicationsRequest) error
}

type Application struct {
	ObjectMeta             *metadata.ObjectMeta
	AlamedaApplicationSpec *AlamedaApplicationSpec
	Controllers            []*Controller
}

type ListApplicationsRequest struct {
	common.QueryCondition
	ApplicationObjectMeta []*ApplicationObjectMeta
}

type DeleteApplicationsRequest struct {
	ApplicationObjectMeta []*ApplicationObjectMeta
}

type ApplicationObjectMeta struct {
	ObjectMeta  *metadata.ObjectMeta
	ScalingTool string
}

type AlamedaApplicationSpec struct {
	ScalingTool string
}

func NewApplication(entity *clusterstatus.ApplicationEntity) *Application {
	application := Application{}
	application.ObjectMeta = &metadata.ObjectMeta{}
	application.ObjectMeta.Name = entity.Name
	application.ObjectMeta.Namespace = entity.Namespace
	application.ObjectMeta.ClusterName = entity.ClusterName
	application.ObjectMeta.Uid = entity.Uid
	application.AlamedaApplicationSpec = NewAlamedaApplicationSpec(entity)
	application.Controllers = make([]*Controller, 0)
	return &application
}

func NewListApplicationsRequest() *ListApplicationsRequest {
	request := ListApplicationsRequest{}
	request.ApplicationObjectMeta = make([]*ApplicationObjectMeta, 0)
	return &request
}

func NewDeleteApplicationsRequest() *DeleteApplicationsRequest {
	request := DeleteApplicationsRequest{}
	request.ApplicationObjectMeta = make([]*ApplicationObjectMeta, 0)
	return &request
}

func NewApplicationObjectMeta(objectMeta *metadata.ObjectMeta, scalingTool string) *ApplicationObjectMeta {
	applicationObjectMeta := ApplicationObjectMeta{}
	applicationObjectMeta.ObjectMeta = objectMeta
	applicationObjectMeta.ScalingTool = scalingTool
	return &applicationObjectMeta
}

func NewAlamedaApplicationSpec(entity *clusterstatus.ApplicationEntity) *AlamedaApplicationSpec {
	spec := AlamedaApplicationSpec{}
	spec.ScalingTool = entity.ScalingTool
	return &spec
}

func (p *Application) BuildEntity() *clusterstatus.ApplicationEntity {
	entity := clusterstatus.ApplicationEntity{}

	entity.Time = influxdb.ZeroTime
	entity.Value = ""

	if p.ObjectMeta != nil {
		entity.Name = p.ObjectMeta.Name
		entity.Namespace = p.ObjectMeta.Namespace
		entity.ClusterName = p.ObjectMeta.ClusterName
		entity.Uid = p.ObjectMeta.Uid
	}

	if p.AlamedaApplicationSpec != nil {
		entity.ScalingTool = p.AlamedaApplicationSpec.ScalingTool
	}

	return &entity
}
