package main

import "fmt"

type plugin struct {
}

func (h *plugin) Todo(data ...any) {
	fmt.Println("Hello, World!")
}

var Plugin plugin
