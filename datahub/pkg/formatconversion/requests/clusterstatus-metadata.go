package requests

import (
	"github.com/turtacn/alameda/datahub/pkg/kubernetes/metadata"
	"github.com/turtacn/api/alameda_api/v1alpha1/datahub/resources"
)

func NewObjectMeta(objectMeta *resources.ObjectMeta) metadata.ObjectMeta {
	meta := metadata.ObjectMeta{}
	if objectMeta != nil {
		meta.Name = objectMeta.GetName()
		meta.Namespace = objectMeta.GetNamespace()
		meta.NodeName = objectMeta.GetNodeName()
		meta.ClusterName = objectMeta.GetClusterName()
		meta.Uid = objectMeta.GetUid()
	}
	return meta
}
