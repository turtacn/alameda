package predictions

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewPodPredictionsDAO(config config.Config) types.PodPredictionsDAO {
	return influxdb.NewPodPredictionsWithConfig(*config.InfluxDB)
}
