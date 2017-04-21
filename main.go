package main

import (
	"fmt"
	"plugin"
	"path/filepath"
	"io/ioutil"
	"strings"
	"github.com/gorilla/mux"
	"net/http"
)

type MetricPlugin interface {
	Name() string
	Desc() string
	Exec() (map[string]string)
}

const PluginDir = "plugins"
const PluginExtension = ".so"

func PluginsReloadHandler(w http.ResponseWriter, r *http.Request) {
	executePlugins(findPlugins())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Reloaded all plugins\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/plugins/reload", PluginsReloadHandler).Methods("PUT")
	http.ListenAndServe(":8000", r)

	executePlugins(findPlugins())
}

func executePlugins(plugins [][]string) {
	fmt.Printf("plugins: %+v\n", plugins)
	for _, ps := range plugins {
		runPlugin(ps[0], ps[1])
	}
}

func findPlugins() [][]string {
	allFiles, _ := ioutil.ReadDir(PluginDir)
	plugins := make([][]string, 0, len(allFiles))
	for _, f := range allFiles {
		if !f.IsDir() && strings.HasSuffix(f.Name(), PluginExtension) {
			pluginPath := filepath.Join(PluginDir, f.Name())
			symbolName := f.Name()[0:(len(f.Name()) - len(PluginExtension))]
			plugins = append(plugins, []string{pluginPath, symbolName})
		}
	}
	return plugins
}

func runPlugin(pluginPath, pluginName string) {
	fmt.Printf("running plugin %s\n", pluginName)
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Errorf("cannot open plugin %s\n", pluginPath)
	}
	symName := strings.Title(pluginName)
	cpu, err := p.Lookup(symName)
	if err != nil {
		fmt.Errorf("cannot lookup %s\n", symName)
	}

	m := cpu.(MetricPlugin)

	fmt.Printf("starting plugin name: %s\n", m.Name())
	fmt.Printf("short desc: %s\n", m.Desc())
	fmt.Printf("m values: %+v\n", m.Exec())
}
