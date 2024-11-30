package main

func main() {
	c := NewVertice("C", nil)
	b := NewVertice("B", []*Vertice{&c})
	a := NewVertice("A", []*Vertice{&b, &c})

	netw := NewNetwork([]*Vertice{&a, &b}, &a, &c, []*Arista{})
	netw.Info()

	println("-- Paths --")
	netw.FindPath(netw.inicio, "")
}
