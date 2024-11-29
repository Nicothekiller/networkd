package main

type Arista struct {
	inicio    *Vertice
	final     *Vertice
	capacidad int
	actual    int
}

func NewArista(inicio, final *Vertice, capacidad, actual int) Arista {
	return Arista{
		inicio:    inicio,
		final:     final,
		capacidad: capacidad,
		actual:    actual,
	}
}

func (self *Arista) info() {
	println("Info Arista")

	println("Inicio:")
	self.inicio.info()
	println("Final:")
	self.final.info()

	println("Capacidad:", self.capacidad)
	println("Actual:", self.actual)
}
