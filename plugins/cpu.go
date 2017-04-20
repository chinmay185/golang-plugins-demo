package main

type CPU struct {
}

var Cpu CPU

func Add(a, b int) int {
	return a + b
}

func (*CPU) Name() string {
	return "CPU Metric"
}
func (*CPU) Desc() string {
	return "Detailed CPU metrics"
}
func (*CPU) Exec() (map[string]string) {
	cpuMetrics := map[string]string{
		"processor":  "0",
		"vendor_id":  "GenuineIntel",
		"cpu family": "6",
		"model":      "70",
		"model name": "Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz",
	}
	return cpuMetrics
}
