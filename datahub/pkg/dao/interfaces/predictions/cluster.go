package predictions

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewClusterPredictionsDAO(config config.Config) types.ClusterPredictionsDAO {
	return influxdb.NewClusterPredictionsWithConfig(*config.InfluxDB)
}
