package scores

import (
	"github.com/turtacn/alameda/datahub/pkg/config"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/scores/influxdb"
	"github.com/turtacn/alameda/datahub/pkg/dao/interfaces/scores/types"
)

func NewScoreDAO(config config.Config) types.ScoreDAO {
	return influxdb.NewScoreWithConfig(*config.InfluxDB)
}
