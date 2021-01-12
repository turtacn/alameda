package main

import (
	"github.com/turtacn/alameda/operator/pkg/probe"
)

func livenessProbe(cfg *probe.LivenessProbeConfig) {
	probe.LivenessProbe(cfg)
}

func readinessProbe(cfg *probe.ReadinessProbeConfig) {
	probe.ReadinessProbe(cfg)
}
