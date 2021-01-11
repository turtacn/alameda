package clusterstatus

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewClusterDAO(config config.Config) types.ClusterDAO {
	return influxdb.NewClusterWithConfig(*config.InfluxDB)
}
