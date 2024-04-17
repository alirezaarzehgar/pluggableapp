package main

import (
	"io"
	"log"
	"os"
	"plugin"

	"gopkg.in/yaml.v3"
)

type PluginInterface interface {
	Todo(data ...any)
}

func main() {
	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatal("cannot open file: %w", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("cannot read file:", err)
	}

	var conf struct {
		Plugins []string `yaml:"plugins"`
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal("cannot unmarshal data:", err)
	}

	for _, p := range conf.Plugins {
		plugin, err := plugin.Open(p)
		if err != nil {
			log.Fatal("err:", err)
		}

		v, err := plugin.Lookup("Plugin")
		if err != nil {
			log.Fatal("lookup err:", err)
		}

		extractedPlugin := v.(PluginInterface)
		extractedPlugin.Todo()
	}
}
