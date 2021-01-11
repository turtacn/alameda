package clusterstatus

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewNodeDAO(config config.Config) types.NodeDAO {
	return influxdb.NewNodeWithConfig(*config.InfluxDB)
}
