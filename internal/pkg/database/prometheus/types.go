package prometheus

import (
	DBCommon "github.com/turtacn/alameda/internal/pkg/database/common"
)

type nodeMetricsFetchingFunction func(nodeName string, options ...DBCommon.Option) ([]Entity, error)
