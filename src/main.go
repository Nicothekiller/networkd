package main

func main() {
	c := NewVertice("C", nil)
	b := NewVertice("B", []*Vertice{&c})
	d := NewVertice("D", []*Vertice{&b, &c})
	a := NewVertice("A", []*Vertice{&b, &c, &d})

	netw := NewNetwork([]*Vertice{&a, &b, &c, &d}, &a, &c, []*Arista{})
	netw.Info()

	println("-- Paths --")
	netw.GenPath()
}
