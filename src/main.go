package main

func main() {
	test := NewVertice("B", nil)

	vert2 := NewVertice("A", []*Vertice{&test})

	ars := NewArista(&vert2, &test, 4, 2)

	netw := NewNetwork([]*Vertice{&vert2, &test}, &vert2, &test, []*Arista{&ars})
	netw.info()
}
