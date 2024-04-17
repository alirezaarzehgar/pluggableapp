package main

import "fmt"

type plugin struct {
}

func (h *plugin) Todo(data ...any) {
	fmt.Println("Bye bye plugin!")
}

var Plugin plugin
