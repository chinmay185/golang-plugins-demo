package main

import (
	"fmt"
	"github.com/tcard/pluginunmarshal"
	"plugin"
)

type MetricPlugin interface {
	Name() string
	Desc() string
	Exec() (map[string]string)
}

func main() {
	pluginName := "plugins/cpu.so"

	//WithRawPluginApiCallAFunction(pluginName)
	//WithPluginMarshalCallAFunction(pluginName)

	WithRawPluginApiCallAStruct(pluginName)
}

func WithRawPluginApiCallAFunction(pluginName string) {
	fmt.Println("WithRawPluginApiCallAFunction")
	p, err := plugin.Open(pluginName)
	if err != nil {
		fmt.Errorf("cannot open plugin %s\n", pluginName)
	}
	methodName := "Add"
	add, err := p.Lookup(methodName)
	if err != nil {
		fmt.Errorf("cannot lookup method %s\n", methodName)
	}
	sum := add.(func(int, int) int)(1, 2)
	fmt.Printf("Sum from plugin is %d\n", sum)
}

func WithPluginMarshalCallAFunction(pluginName string) {
	fmt.Println("WithPluginMarshalCallAFunction")
	var p struct {
		Add func(a, b int) int
	}

	err := pluginunmarshal.Open(pluginName, &p)
	if err != nil {
		fmt.Errorf("could not open %s: %v", pluginName, err)
	}

	fmt.Printf("Sum from plugin is %d\n", p.Add(1, 2))
}

func WithRawPluginApiCallAStruct(pluginName string) {
	fmt.Println("WithRawPluginApiCallAStruct")
	p, err := plugin.Open(pluginName)
	if err != nil {
		fmt.Errorf("cannot open plugin %s\n", pluginName)
	}
	symName := "Cpu"
	cpu, err := p.Lookup(symName)
	if err != nil {
		fmt.Errorf("cannot lookup %s\n", symName)
	}

	m := cpu.(MetricPlugin)

	fmt.Printf("starting plugin name: %s\n", m.Name())
	fmt.Printf("short desc: %s\n", m.Desc())
	fmt.Printf("m values: %+v\n", m.Exec())
}
