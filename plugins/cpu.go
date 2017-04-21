package main

import (
	"github.com/shirou/gopsutil/cpu"
)

type CPUMetric struct {
}

var Cpu CPUMetric

func (*CPUMetric) Name() string {
	return "CPU metric"
}
func (*CPUMetric) Desc() string {
	return "Detailed CPU metrics"
}
func (*CPUMetric) Exec() (map[string]interface{}) {
	info, _ := cpu.Info()
	cpuMetrics := map[string]interface{}{
		"info": info,
	}
	return cpuMetrics
}
