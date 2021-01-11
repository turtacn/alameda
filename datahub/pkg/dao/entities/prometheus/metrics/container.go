package metrics

import (
	DaoMetricTypes "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "github.com/turtacn/alameda/datahub/pkg/formatconversion/enumconv"
	FormatTypes "github.com/turtacn/alameda/datahub/pkg/formatconversion/types"
	"github.com/turtacn/alameda/datahub/pkg/kubernetes/metadata"
	InternalPromth "github.com/turtacn/alameda/internal/pkg/database/prometheus"
)

type ContainerCPUUsageMillicoresEntity struct {
	PrometheusEntity InternalPromth.Entity

	Namespace     string
	PodName       string
	ContainerName string
	Samples       []FormatTypes.Sample
}

// ContainerMetric Build ContainerMetric base on entity properties
func (e *ContainerCPUUsageMillicoresEntity) ContainerMetric() DaoMetricTypes.ContainerMetric {

	var (
		containerMetric DaoMetricTypes.ContainerMetric
	)

	containerMetric = DaoMetricTypes.ContainerMetric{
		ObjectMeta: DaoMetricTypes.ContainerMeta{
			ObjectMeta: metadata.ObjectMeta{
				Namespace: e.Namespace,
				Name:      e.ContainerName,
			},
			PodName: e.PodName,
		},
		Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeCPUUsageSecondsPercentage: e.Samples,
		},
	}

	return containerMetric
}

type ContainerMemoryUsageBytesEntity struct {
	PrometheusEntity InternalPromth.Entity

	Namespace     string
	PodName       string
	ContainerName string
	Samples       []FormatTypes.Sample
}

// ContainerMetric Build ContainerMetric base on entity properties
func (e *ContainerMemoryUsageBytesEntity) ContainerMetric() DaoMetricTypes.ContainerMetric {

	var (
		containerMetric DaoMetricTypes.ContainerMetric
	)

	containerMetric = DaoMetricTypes.ContainerMetric{
		ObjectMeta: DaoMetricTypes.ContainerMeta{
			ObjectMeta: metadata.ObjectMeta{
				Namespace: e.Namespace,
				Name:      e.ContainerName,
			},
			PodName: e.PodName,
		},
		Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeMemoryUsageBytes: e.Samples,
		},
	}

	return containerMetric
}
