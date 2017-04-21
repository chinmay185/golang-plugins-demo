package main

import "github.com/shirou/gopsutil/disk"

type DiskMetric struct {
}

var Disk DiskMetric

func (*DiskMetric) Name() string {
	return "updated Disk metric"
}
func (*DiskMetric) Desc() string {
	return "updated detailed disk metrics"
}
func (*DiskMetric) Exec() (map[string]interface{}) {
	partitions, _ := disk.Partitions(false)
	diskMetrics := map[string]interface{}{
		"info": partitions,
	}
	return diskMetrics
}
