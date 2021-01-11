package recommendations

import (
	"github.com/turtacn/alameda/internal/pkg/database/influxdb"
)

const (
	Container   influxdb.Measurement = "container"
	Controller  influxdb.Measurement = "controller"
	Application influxdb.Measurement = "application"
	Namespace   influxdb.Measurement = "namespace"
	Node        influxdb.Measurement = "node"
	Cluster     influxdb.Measurement = "cluster"
)
