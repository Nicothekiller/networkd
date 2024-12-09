package main

func main() {
	c := NewVertice("C", nil)
	b := NewVertice("B", []*Vertice{&c})
	d := NewVertice("D", []*Vertice{&b, &c})
	a := NewVertice("A", []*Vertice{&b, &c, &d})

	bc := NewArista(&b, &c, 6)
	db := NewArista(&d, &b, 4)
	dc := NewArista(&d, &c, 4)
	ab := NewArista(&a, &b, 4)
	ac := NewArista(&a, &c, 4)
	ad := NewArista(&a, &d, 6)

	netw := NewNetwork([]*Vertice{&a, &b, &c, &d}, &a, &c, []*Arista{&bc, &db, &dc, &ab, &ac, &ad})

	println("-- Paths --")
	netw.GenPath()

	// netw.Info()

	println("-- Despues --")
	netw.MaximizeFlow()
	netw.PathInfo()
}
