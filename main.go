package main

import (
	"log"
	"net/http"
	"os"
	"plugin"
	"strings"

	"github.com/labstack/echo/v4"
)

type PluginInterface interface {
	Register(e *echo.Echo)
}

func RegisterPluginRoutes(e *echo.Echo) {
	de, err := os.ReadDir("./plugins")
	if err != nil {
		log.Fatal("cannot read plugins dir:", err)
	}

	for _, d := range de {
		name := "./plugins/" + d.Name()
		if !strings.HasSuffix(name, ".so") {
			continue
		}

		plugin, err := plugin.Open(name)
		if err != nil {
			log.Fatal("cannot open plugin:", err)
		}

		v, err := plugin.Lookup("Plugin")
		if err != nil {
			log.Fatal("lookup err:", err)
		}

		extractedPlugin := v.(PluginInterface)
		extractedPlugin.Register(e)
	}
}

func TODOHandler(c echo.Context) error {
	return c.String(http.StatusOK, "TODO Handler")
}

func main() {
	e := echo.New()

	e.GET("/api/foo", TODOHandler)
	e.GET("/api/bar", TODOHandler)
	e.GET("/api/baz", TODOHandler)

	RegisterPluginRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
