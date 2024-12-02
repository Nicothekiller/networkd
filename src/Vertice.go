package main

// struct vertice, has in id and pointers to the next vertices. It represents a vertice in a directed graph
type Vertice struct {
	id   string
	next []*Vertice
}

// constructor for vertice struct
func NewVertice(id string, next []*Vertice) Vertice {
	return Vertice{
		id:   id,
		next: next,
	}
}

// info method for vertice
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
