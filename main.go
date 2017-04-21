package main

import (
	"fmt"
	"plugin"
	"path/filepath"
	"io/ioutil"
	"strings"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type MetricPlugin interface {
	Name() string
	Desc() string
	Exec() (map[string]interface{})
}

const PluginDir = "plugins"
const PluginExtension = ".so"

func main() {
	fmt.Println("Running existing plugins from `plugins` directory")

	findAndExecPlugins()

	fmt.Println("Add new plugins in `plugins` directory that satisfy MetricPlugin interface and `curl -X PUT http://localhost:8000/plugins/reload` to reload the plugins")
	r := mux.NewRouter()
	r.HandleFunc("/plugins/reload", PluginsReloadHandler).Methods("PUT")
	http.ListenAndServe(":8000", r)
}

func PluginsReloadHandler(w http.ResponseWriter, r *http.Request) {
	plugins := findAndExecPlugins()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "reloaded %v\n", plugins)
}

func findAndExecPlugins() [][]string {
	plugins := findPlugins()
	executePlugins(plugins)
	return plugins
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

func executePlugins(plugins [][]string) {
	fmt.Printf("plugins: %+v\n", plugins)
	for _, ps := range plugins {
		runPlugin(ps[0], ps[1])
	}
}

func runPlugin(pluginPath, pluginName string) {
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

	fmt.Printf("plugin: %s\n", m.Name())
	fmt.Printf("short desc: %s\n", m.Desc())
	output := m.Exec()
	mo := output["info"]

	jsonOut, err := json.MarshalIndent(mo, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("output:", string(jsonOut))
}
