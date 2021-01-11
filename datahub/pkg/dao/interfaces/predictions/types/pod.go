package types

import (
	"fmt"
	"github.com/turtacn/alameda/datahub/pkg/kubernetes/metadata"
	DBCommon "github.com/turtacn/alameda/internal/pkg/database/common"
	ApiPredictions "github.com/turtacn/api/alameda_api/v1alpha1/datahub/predictions"
)

// DAO DAO interface of prediction
type PodPredictionsDAO interface {
	CreatePredictions(PodPredictionMap) error
	ListPredictions(ListPodPredictionsRequest) (PodPredictionMap, error)
	FillPredictions(predictions []*ApiPredictions.PodPrediction, fillDays int64) error
}

// PodPrediction Prediction model to represent one pod's Prediction
type PodPrediction struct {
	ObjectMeta             metadata.ObjectMeta
	ContainerPredictionMap ContainerPredictionMap
}

// PodsPredictionMap Pods' Prediction map
type PodPredictionMap struct {
	MetricMap map[metadata.NamespacePodName]*PodPrediction
}

// ListPodPredictionsRequest ListPodPredictionsRequest
type ListPodPredictionsRequest struct {
	DBCommon.QueryCondition
	ObjectMeta   []metadata.ObjectMeta
	ModelId      string
	PredictionId string
	Granularity  int64
	FillDays     int64
}

func NewPodPrediction() *PodPrediction {
	podPrediction := &PodPrediction{}
	podPrediction.ContainerPredictionMap = NewContainerPredictionMap()
	return podPrediction
}

func NewPodPredictionMap() PodPredictionMap {
	podPredictionMap := PodPredictionMap{}
	podPredictionMap.MetricMap = make(map[metadata.NamespacePodName]*PodPrediction)
	return podPredictionMap
}

func NewListPodPredictionsRequest() ListPodPredictionsRequest {
	request := ListPodPredictionsRequest{}
	request.ObjectMeta = make([]metadata.ObjectMeta, 0)
	return request
}

// NamespacePodName Return identity of the pod Prediction
func (p *PodPrediction) NamespacePodName() metadata.NamespacePodName {
	return metadata.NamespacePodName(fmt.Sprintf("%s/%s", p.ObjectMeta.Namespace, p.ObjectMeta.Name))
}

// Merge Merge current PodPrediction with input PodPrediction
func (p *PodPrediction) Merge(in *PodPrediction) {
	for _, containerPrediction := range in.ContainerPredictionMap.MetricMap {
		p.ContainerPredictionMap.AddContainerPrediction(containerPrediction)
	}
}

// AddPodPrediction Add pod Prediction into PodsPredictionMap
func (p *PodPredictionMap) AddPodPrediction(podPrediction *PodPrediction) {
	namespacePodName := podPrediction.NamespacePodName()
	if existPodPrediction, exist := p.MetricMap[namespacePodName]; exist {
		existPodPrediction.Merge(podPrediction)
	} else {
		p.MetricMap[namespacePodName] = podPrediction
	}
}

// AddContainerPrediction Add container Prediction into PodsPredictionMap
func (p *PodPredictionMap) AddContainerPrediction(c *ContainerPrediction) {
	podPrediction := c.BuildPodPrediction()
	namespacePodName := podPrediction.NamespacePodName()
	if existedPodPrediction, exist := p.MetricMap[namespacePodName]; exist {
		existedPodPrediction.Merge(podPrediction)
	} else {
		p.MetricMap[namespacePodName] = podPrediction
	}
}
