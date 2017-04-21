package main

type DiskInterface struct {
}

var Disk DiskInterface

func (*DiskInterface) Name() string {
	return "Disk Metric"
}
func (*DiskInterface) Desc() string {
	return "Detailed disk usage metrics"
}
func (*DiskInterface) Exec() (map[string]string) {
	cpuMetrics := map[string]string{
		"disk":      "1",
		"vendor_id": "Sandisk",
		"capacity":  "500Gb",
	}
	return cpuMetrics
}
