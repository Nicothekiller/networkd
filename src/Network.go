package main

type Network struct {
	totalVertices []*Vertice
	inicio        *Vertice
	final         *Vertice
	totalAristas  []*Arista
}

func NewNetwork(totalVertices []*Vertice, inicio, final *Vertice, totalAristas []*Arista) Network {
	return Network{
		totalVertices: totalVertices,
		inicio:        inicio,
		final:         final,
		totalAristas:  totalAristas,
	}
}

func (self *Network) Info() {
	println("-- Network info --")

	println("-- Inicio --")
	self.inicio.info()

	println("-- Final --")
	self.final.info()

	print("Vertices: ")
	for i := 0; i < len(self.totalVertices); i++ {
		print(self.totalVertices[i].id, " ")
	}
	println("\n")

	print("Aristas: ")
	for i := 0; i < len(self.totalAristas); i++ {
		print(self.totalAristas[i].id, " ")
	}
	println("\n")
}

func (self *Network) FindPath(iter *Vertice, path string) {
	path += iter.id
	if iter.id != self.final.id {
		for i := 0; i < len(iter.next); i++ {
			self.FindPath(iter.next[i], path)
		}
	} else {
		println("Found path:", path)
	}
}
