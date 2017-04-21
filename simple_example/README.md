## Running simple example

Go to simple_example directory and build the plugin with following command

```
cd simple_example
go build -buildmode=plugin -o simple_plugin.so simple_plugin.go
```
This will create a `simple_example.so` file. To load the plugin and run the code inside this plugin, execute following command

`go run simple_plugin_example.go`

Output will be as follows

```
﻿Calling a function using simple plugin API
Sum from plugin is 3
Calling a function using plugin marshal API
Sum from plugin is 3
```

[pluginmarshal](https://github.com/tcard/pluginunmarshal) package unmarshals Go plugins into structs. 