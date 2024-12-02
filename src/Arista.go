package main

// Struct for arista. has an id that should be the id of the 2 vertices it points to combined.
// Has a pointer to start and to the end in the form of vertices. Also has capacity and current flow
type Arista struct {
	id        string
	inicio    *Vertice
	final     *Vertice
	capacidad int
	actual    int
}

// Constructor for arista. Doesnt have id parameter, it should be set automatically.
func NewArista(inicio, final *Vertice, capacidad, actual int) Arista {
	id := inicio.id + final.id
	return Arista{
		id:        id,
		inicio:    inicio,
		final:     final,
		capacidad: capacidad,
		actual:    actual,
	}
}

// Info method for arista
func (self *Arista) info() {
	println("Info Arista")

	println("Inicio:")
	self.inicio.info()
	println("Final:")
	self.final.info()

	println("Capacidad:", self.capacidad)
	println("Actual:", self.actual)
}
