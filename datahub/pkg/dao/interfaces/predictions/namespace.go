package predictions

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/types"
)

func NewNamespacePredictionsDAO(config config.Config) types.NamespacePredictionsDAO {
	return influxdb.NewNamespacePredictionsWithConfig(*config.InfluxDB)
}
