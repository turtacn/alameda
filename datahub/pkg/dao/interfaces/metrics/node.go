package metrics

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	DaoClusterStatus "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/prometheus"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/metrics/types"
)

func NewNodeMetricsReaderDAO(config config.Config) types.NodeMetricsDAO {
	switch config.Apis.Metrics.Source {
	case "influxdb":
		return influxdb.NewNodeMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewNodeMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	default:
		return prometheus.NewNodeMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	}
}

func NewNodeMetricsWriterDAO(config config.Config) types.NodeMetricsDAO {
	switch config.Apis.Metrics.Target {
	case "influxdb":
		return influxdb.NewNodeMetricsWithConfig(*config.InfluxDB)
	case "prometheus":
		return prometheus.NewNodeMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	default:
		return prometheus.NewNodeMetricsWithConfig(*config.Prometheus, DaoClusterStatus.NewNodeDAO(config), config.ClusterUID)
	}
}
