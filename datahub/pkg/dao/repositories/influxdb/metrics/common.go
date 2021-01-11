package metrics

import (
	"github.com/turtacn/alameda/internal/pkg/database/common"
	"github.com/turtacn/alameda/internal/pkg/database/influxdb"
)

var aggregateFuncToInfluxDBFunc = map[common.AggregateFunction]influxdb.Functions{
	common.None:        influxdb.Last,
	common.MaxOverTime: influxdb.Max,
	common.AvgOverTime: influxdb.Mean,
}
