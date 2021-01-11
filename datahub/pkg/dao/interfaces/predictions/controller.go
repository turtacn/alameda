package predictions

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewControllerPredictionsDAO(config config.Config) types.ControllerPredictionsDAO {
	return influxdb.NewControllerPredictionsWithConfig(*config.InfluxDB)
}
