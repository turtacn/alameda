package responses

import (
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	"github.com/turtacn/api/alameda_api/v1alpha1/datahub/resources"
)

type ClusterExtended struct {
	*types.Cluster
}

func (p *ClusterExtended) ProduceCluster() *resources.Cluster {
	cluster := &resources.Cluster{
		ObjectMeta: NewObjectMeta(p.ObjectMeta),
	}
	return cluster
}
