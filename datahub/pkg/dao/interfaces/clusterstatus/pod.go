package clusterstatus

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
)

func NewPodDAO(config config.Config) types.PodDAO {
	return influxdb.NewPodWithConfig(*config.InfluxDB)
}
