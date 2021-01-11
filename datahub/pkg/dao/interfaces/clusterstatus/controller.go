package clusterstatus

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewControllerDAO(config config.Config) types.ControllerDAO {
	return influxdb.NewControllerWithConfig(*config.InfluxDB)
}
