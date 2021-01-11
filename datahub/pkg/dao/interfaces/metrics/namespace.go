package metrics

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	DaoClusterStatus "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/prometheus"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/types"
)

func NewNamespaceMetricsReaderDAO(config config.Config) types.NamespaceMetricsDAO {
	switch config.Apis.Metrics.Source {
	case "influxdb":
		return influxdb.NewNamespaceMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewNamespaceMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNamespaceDAO(config), config.ClusterUID)
	default:
		return prometheus.NewNamespaceMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNamespaceDAO(config), config.ClusterUID)
	}
}

func NewNamespaceMetricsWriterDAO(config config.Config) types.NamespaceMetricsDAO {
	switch config.Apis.Metrics.Target {
	case "influxdb":
		return influxdb.NewNamespaceMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewNamespaceMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNamespaceDAO(config), config.ClusterUID)
	default:
		return prometheus.NewNamespaceMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNamespaceDAO(config), config.ClusterUID)
	}
}
