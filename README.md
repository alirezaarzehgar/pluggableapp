# What is this project ?
This project is a simple representation of [plugin](https://pkg.go.dev/plugin) library in Go and a pluggable system implementation.
You can easily understand how it's work.

You can add plugin to `./plugins`. Create new go source code on this folder and then run `make run`. Application will register your routes on plugin.

For example:
```bash
cat << _EOF_ > ./plugins/w8.go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type plugin struct {
}

func TODOHandler(c echo.Context) error {
	return c.String(http.StatusOK, "blahblah Plugin TODO Handler")
}

func (h *plugin) Register(e *echo.Echo) {
	p := e.Group("/plugin/blahblah")
	p.GET("/foo", TODOHandler)
	p.GET("/bar", TODOHandler)
	p.GET("/baz", TODOHandler)
}

var Plugin plugin
_EOF_

make
```
