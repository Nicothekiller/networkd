package main

func main() {
	c := NewVertice("C", nil)
	b := NewVertice("B", []*Vertice{&c})
	d := NewVertice("D", []*Vertice{&b, &c})
	a := NewVertice("A", []*Vertice{&b, &c, &d})

	bc := NewArista(&b, &c, 4, 0)
	db := NewArista(&d, &b, 4, 0)
	dc := NewArista(&d, &c, 4, 0)
	ab := NewArista(&a, &b, 4, 0)
	ac := NewArista(&a, &c, 4, 0)
	ad := NewArista(&a, &d, 4, 0)

	netw := NewNetwork([]*Vertice{&a, &b, &c, &d}, &a, &c, []*Arista{&bc, &db, &dc, &ab, &ac, &ad})
	netw.Info()

	println("-- Paths --")
	netw.GenPath()
}
