package main

// to build this file as a plugin run go build -buildmode=plugin -o simple_plugin.so simple_plugin.go
func Add(a, b int) int {
	return a + b
}
