# What is this project ?
This project is a simple representation of [https://pkg.go.dev/plugin](plugin) library in Go and a pluggable system implementation.
You can easily understand how it's work.

You can add plugin to `./plugins`. run make and add new `.so` to `config.yaml` file.

For example:
```bash
cat << _EOF_ > ./plugins/w8.go
package main

import (
        "fmt"
        "time"
)

type plugin struct {
}

func (h *plugin) Todo(data ...any) {
        time.Sleep(time.Second)
        fmt.Println("Ops! Thank you for your patience.")
}

var Plugin plugin
_EOF_

cat << _EOF_ > ./config.yaml
plugins:
 - ./plugins/hello.so
 - ./plugins/bye.so
 - ./plugins/w8.so
_EOF_

make
```