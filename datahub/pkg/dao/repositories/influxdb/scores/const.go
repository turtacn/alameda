package scores

import (
	InternalInflux "github.com/turtacn/alameda/internal/pkg/database/influxdb"
)

const (
	// SimulatedSchedulingScore Measurement name of simulated scheduling score in influxdb
	SimulatedSchedulingScore InternalInflux.Measurement = "simulated_scheduling_score"
)
