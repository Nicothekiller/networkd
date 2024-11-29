package main

type Vertice struct {
	id   string
	next []*Vertice
}

func NewVertice(id string, next []*Vertice) Vertice {
	return Vertice{
		id:   id,
		next: next,
	}
}

func (self *Vertice) info() {
	println("id:", self.id)
	if self.next != nil {
		print("Next: ")
		for i := 0; i < len(self.next); i++ {
			print(self.next[i].id, " ")
		}
		println()
	}
	println()
}
