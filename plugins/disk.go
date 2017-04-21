package main

import "github.com/shirou/gopsutil/disk"

type diskMetric struct {
}

// exported variable to be used from main program
var Disk diskMetric

func (*diskMetric) Name() string {
	return "Disk metric"
}
func (*diskMetric) Desc() string {
	return "Detailed disk metrics"
}
func (*diskMetric) Exec() (map[string]interface{}) {
	partitions, _ := disk.Partitions(false)
	diskMetrics := map[string]interface{}{
		"info": partitions,
	}
	return diskMetrics
}
