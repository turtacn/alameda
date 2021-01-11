package events

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/events/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/events/types"
)

func NewEventDAO(config config.Config) types.EventDAO {
	return influxdb.NewEventWithConfig(config.InfluxDB, config.RabbitMQ)
}
