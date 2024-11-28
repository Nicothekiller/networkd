package main

import "fmt"

type vertice struct {
	id         string
	fcapacidad int
	factual    int
	next       []*vertice
}

func (self *vertice) info() {
	println("id: ", self.id)
	println("fcapacidad: ", self.fcapacidad)
	println("factual: ", self.factual)
	fmt.Println("factual: ", self.next)
}
