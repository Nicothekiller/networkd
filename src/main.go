package main

func main() {
	test := vertice{
		id:         "1",
		fcapacidad: 2,
		factual:    1,
		next:       []*vertice{},
	}
	test.info()
}
