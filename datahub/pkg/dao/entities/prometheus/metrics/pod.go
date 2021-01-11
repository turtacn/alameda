package metrics

import (
	FormatTypes "github.com/turtacn/alameda/datahub/pkg/formatconversion/types"
)

type PodCPUUsageMillicoresEntity struct {
	NamespaceName string
	PodName       string
	Samples       []FormatTypes.Sample
}

type PodMemoryUsageBytesEntity struct {
	NamespaceName string
	PodName       string
	Samples       []FormatTypes.Sample
}
