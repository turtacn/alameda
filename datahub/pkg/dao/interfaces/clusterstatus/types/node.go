package types

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/turtacn/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	"github.com/turtacn/alameda/datahub/pkg/kubernetes/metadata"
	"github.com/turtacn/alameda/internal/pkg/database/common"
	"github.com/turtacn/alameda/internal/pkg/database/influxdb"
)

// Node provides node measurement operations
type NodeDAO interface {
	CreateNodes([]*Node) error
	ListNodes(*ListNodesRequest) ([]*Node, error)
	DeleteNodes(*DeleteNodesRequest) error
}

type Node struct {
	ObjectMeta      *metadata.ObjectMeta
	CreateTime      *timestamp.Timestamp
	Capacity        *Capacity
	AlamedaNodeSpec *AlamedaNodeSpec
}

type ListNodesRequest struct {
	common.QueryCondition
	ObjectMeta []*metadata.ObjectMeta
}

type DeleteNodesRequest struct {
	ObjectMeta []*metadata.ObjectMeta
}

type Capacity struct {
	CpuCores                 int64
	MemoryBytes              int64
	NetworkMegabitsPerSecond int64
}

type AlamedaNodeSpec struct {
	Provider *Provider
}

type Provider struct {
	Provider     string
	InstanceType string
	Region       string
	Zone         string
	Os           string
	Role         string
	InstanceId   string
	StorageSize  int64
}

func NewNode(entity *clusterstatus.NodeEntity) *Node {
	node := Node{}
	node.ObjectMeta = &metadata.ObjectMeta{}
	node.ObjectMeta.Name = entity.Name
	node.ObjectMeta.ClusterName = entity.ClusterName
	node.ObjectMeta.Uid = entity.Uid
	node.CreateTime = &timestamp.Timestamp{Seconds: entity.CreateTime}
	node.Capacity = NewCapacity(entity)
	node.AlamedaNodeSpec = NewAlamedaNodeSpec(entity)
	return &node
}

func NewListNodesRequest() *ListNodesRequest {
	request := ListNodesRequest{}
	request.ObjectMeta = make([]*metadata.ObjectMeta, 0)
	return &request
}

func NewDeleteNodesRequest() *DeleteNodesRequest {
	request := DeleteNodesRequest{}
	request.ObjectMeta = make([]*metadata.ObjectMeta, 0)
	return &request
}

func NewCapacity(entity *clusterstatus.NodeEntity) *Capacity {
	capacity := Capacity{}
	capacity.CpuCores = entity.CPUCores
	capacity.MemoryBytes = entity.MemoryBytes
	capacity.NetworkMegabitsPerSecond = entity.NetworkMbps
	return &capacity
}

func NewAlamedaNodeSpec(entity *clusterstatus.NodeEntity) *AlamedaNodeSpec {
	spec := AlamedaNodeSpec{}
	spec.Provider = NewProvider(entity)
	return &spec
}

func NewProvider(entity *clusterstatus.NodeEntity) *Provider {
	provider := Provider{}
	provider.Provider = entity.IOProvider
	provider.InstanceType = entity.IOInstanceType
	provider.Region = entity.IORegion
	provider.Zone = entity.IOZone
	provider.Os = entity.IOOS
	provider.Role = entity.IORole
	provider.InstanceId = entity.IOInstanceID
	provider.StorageSize = entity.IOStorageSize
	return &provider
}

func (p *Node) BuildEntity() *clusterstatus.NodeEntity {
	entity := clusterstatus.NodeEntity{}

	entity.Time = influxdb.ZeroTime

	if p.ObjectMeta != nil {
		entity.Name = p.ObjectMeta.Name
		entity.ClusterName = p.ObjectMeta.ClusterName
		entity.Uid = p.ObjectMeta.Uid
	}

	if p.CreateTime != nil {
		entity.CreateTime = p.CreateTime.GetSeconds()
	}

	if p.Capacity != nil {
		entity.CPUCores = p.Capacity.CpuCores
		entity.MemoryBytes = p.Capacity.MemoryBytes
		entity.NetworkMbps = p.Capacity.NetworkMegabitsPerSecond
	}

	if p.AlamedaNodeSpec != nil {
		if p.AlamedaNodeSpec.Provider != nil {
			entity.IOProvider = p.AlamedaNodeSpec.Provider.Provider
			entity.IOInstanceType = p.AlamedaNodeSpec.Provider.InstanceType
			entity.IORegion = p.AlamedaNodeSpec.Provider.Region
			entity.IOZone = p.AlamedaNodeSpec.Provider.Zone
			entity.IOOS = p.AlamedaNodeSpec.Provider.Os
			entity.IORole = p.AlamedaNodeSpec.Provider.Role
			entity.IOInstanceID = p.AlamedaNodeSpec.Provider.InstanceId
			entity.IOStorageSize = p.AlamedaNodeSpec.Provider.StorageSize
		}
	}

	return &entity
}
