package main

import (
	"github.com/shirou/gopsutil/cpu"
)

type cpuMetric struct {
}

// exported variable to be used from main program
var Cpu cpuMetric

func (*cpuMetric) Name() string {
	return "CPU metric"
}
func (*cpuMetric) Desc() string {
	return "Detailed CPU metrics"
}
func (*cpuMetric) Exec() (map[string]interface{}) {
	info, _ := cpu.Info()
	cpuMetrics := map[string]interface{}{
		"info": info,
	}
	return cpuMetrics
}
