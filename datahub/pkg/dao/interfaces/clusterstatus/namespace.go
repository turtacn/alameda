package clusterstatus

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewNamespaceDAO(config config.Config) types.NamespaceDAO {
	return influxdb.NewNamespaceWithConfig(*config.InfluxDB)
}
