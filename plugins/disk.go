package main

import "github.com/shirou/gopsutil/disk"

type DiskMetric struct {
}

var Disk DiskMetric

func (*DiskMetric) Name() string {
	return "Disk metric"
}
func (*DiskMetric) Desc() string {
	return "Detailed disk metrics"
}
func (*DiskMetric) Exec() (map[string]interface{}) {
	partitions, _ := disk.Partitions(false)
	diskMetrics := map[string]interface{}{
		"info": partitions,
	}
	return diskMetrics
}
