package main

// Network struct. Based on a directed graph, has all vertices, start, end, all edges and paths.
// All paths should be set automatically by findPath / GenPath
type Network struct {
	totalVertices []*Vertice
	inicio        *Vertice
	final         *Vertice
	totalAristas  []*Arista
	paths         [][]*Vertice
}

// Constructor for Network struct.
func NewNetwork(totalVertices []*Vertice, inicio, final *Vertice, totalAristas []*Arista) Network {
	return Network{
		totalVertices: totalVertices,
		inicio:        inicio,
		final:         final,
		totalAristas:  totalAristas,
	}
}

// Info method for network func. Displays info.
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

// Method for finding a path from a vertice to the end of the network. Needs start and where to save pointers to vertices.
func (self *Network) findPath(iter *Vertice, path []*Vertice) {
	path = append(path, iter)

	if iter.id != self.final.id {
		for i := 0; i < len(iter.next); i++ {
			self.findPath(iter.next[i], path)
		}
	} else {
		print("Found path: ")
		for i := 0; i < len(path); i++ {
			print(path[i].id, " ")
		}
		print("\n")

		self.paths = append(self.paths, path)
	}
}

// Wrapper method to find all paths from start to end on the network. Uses findPath internally.
func (self *Network) GenPath() {
	self.findPath(self.inicio, []*Vertice{})
}
