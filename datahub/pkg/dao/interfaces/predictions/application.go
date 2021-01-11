package predictions

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewApplicationPredictionsDAO(config config.Config) types.ApplicationPredictionsDAO {
	return influxdb.NewApplicationPredictionsWithConfig(*config.InfluxDB)
}
