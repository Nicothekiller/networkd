package main

func main() {
	test := NewVertice("B", nil)

	vert2 := NewVertice("A", []*Vertice{&test})

	ars := NewArista(&vert2, &test, 4, 2)
	ars.info()
}
