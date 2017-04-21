package main

import (
	"fmt"
	"github.com/tcard/pluginunmarshal"
	"plugin"
)

func main() {

	pluginName := "simple_plugin.so"

	WithRawPluginApiCallAFunction(pluginName)
	WithPluginMarshalCallAFunction(pluginName)
}

func WithRawPluginApiCallAFunction(pluginName string) {
	fmt.Println("Calling a function using simple plugin API")
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
	fmt.Println("Calling a function using plugin marshal API")
	var p struct {
		Add func(a, b int) int
	}

	err := pluginunmarshal.Open(pluginName, &p)
	if err != nil {
		fmt.Errorf("could not open %s: %v", pluginName, err)
	}

	fmt.Printf("Sum from plugin is %d\n", p.Add(1, 2))
}
