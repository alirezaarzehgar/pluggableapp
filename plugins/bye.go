package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type plugin struct {
}

func TODOHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Plugin TODO Handler")
}

func (h *plugin) Register(e *echo.Echo) {
	p := e.Group("/plugin/bye")
	p.GET("/foo", TODOHandler)
	p.GET("/bar", TODOHandler)
	p.GET("/baz", TODOHandler)
}

var Plugin plugin
