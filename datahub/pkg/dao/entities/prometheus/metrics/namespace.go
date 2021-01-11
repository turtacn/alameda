package metrics

import (
	DaoMetricTypes "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/types"
	FormatEnum "github.com/turtacn/alameda/datahub/pkg/formatconversion/enumconv"
	FormatTypes "github.com/turtacn/alameda/datahub/pkg/formatconversion/types"
	"github.com/turtacn/alameda/datahub/pkg/kubernetes/metadata"
	InternalPromth "github.com/turtacn/alameda/internal/pkg/database/prometheus"
)

type NamespaceCPUUsageMillicoresEntity struct {
	PrometheusEntity InternalPromth.Entity

	NamespaceName string
	Samples       []FormatTypes.Sample
}

// NamespaceMetric Build NamespaceMetric base on entity properties
func (e *NamespaceCPUUsageMillicoresEntity) NamespaceMetric() DaoMetricTypes.NamespaceMetric {

	m := DaoMetricTypes.NamespaceMetric{
		ObjectMeta: metadata.ObjectMeta{
			Name: e.NamespaceName,
		},
		Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeCPUUsageSecondsPercentage: e.Samples,
		},
	}
	return m
}

type NamespaceMemoryUsageBytesEntity struct {
	PrometheusEntity InternalPromth.Entity

	NamespaceName string
	Samples       []FormatTypes.Sample
}

// NamespaceMetric Build NamespaceMetric base on entity properties
func (e *NamespaceMemoryUsageBytesEntity) NamespaceMetric() DaoMetricTypes.NamespaceMetric {

	m := DaoMetricTypes.NamespaceMetric{
		ObjectMeta: metadata.ObjectMeta{
			Name: e.NamespaceName,
		}, Metrics: map[FormatEnum.MetricType][]FormatTypes.Sample{
			FormatEnum.MetricTypeMemoryUsageBytes: e.Samples,
		},
	}
	return m
}
