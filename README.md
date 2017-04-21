# Demo to showcase Go 1.8 [plugins](https://golang.org/pkg/plugin/)

Download the repo or use `go get github.com/chinmay185/golang-plugins-demo` to get the code locally.

## Instructions

#### Pre-requisites
* Go 1.8 or above
* Linux (for now)
* [Glide](https://github.com/Masterminds/glide)

Go to [simple_example](https://github.com/chinmay185/golang-plugins-demo/tree/master/simple_example) for a demo of calling a function exported from a plugin.

### Running CPU and Disk Metrics plugin
To run the plugins execute the following commands
```
go build -buildmode=plugin -o plugins/cpu.so plugins/cpu.go
go run main.go
```
This will execute cpu plugin and display some cpu stats on the console. To display the disk stats (without restarting the main program), execute following commands

```
go build -buildmode=plugin -o plugins/disk.so plugins/disk.go
﻿curl -X PUT http://localhost:8000/plugins/reload
```
You should now be able to see disk stats along with cpu stats.

### CPU and Disk Metrics plugin
- `plugins` directory contains two plugins, namely `cpu.go` and `disk.go`
- `main.go` defines a `MetricPlugin` interface
- Files in the `plugins` directory should export a struct) that satisfies the `MetricPlugin` interface
- Exported variable name follows Title case convention. ex. `cpu.go` exports `Cpu` and `disk.go` exports `Disk`. If you write a new `mem.go`, be sure to export `Mem` from that file.
- `main.go` finds all the plugins (i.e. all `.so` files under `plugins` directory) and execute them. It also starts a http web server which can be used to reload new plugins.
- For reloading the new plugins, add the corresponding `.go` files in `plugins` directory and build them using `go build -buildmode=plugin -o plugin.so plugin.go` command and then use the curl command mentioned above to reload all plugins 

## Some Go plugin limitations / gotchas
- Plugins are only supported on linux (for now)
- Plugin code needs to be compiled with the same version of Go that was used for compiling the main app (for ex. Go 1.9 plugins won't work on Go 1.8 app)
- Plugin has to be in the main package.﻿Packages not named main are ignored. For more info, run `go help buildmode`
- Plugins can't be reloaded (no hot-code swapping) Refer [this issue](https://github.com/golang/go/issues/17980) for more details
- Loading a copy of the already loaded plugin panics. Refer [here](https://github.com/golang/go/issues/19004)

## License
MIT