package main

type Foo struct {
	Bar string
	Baz float64
	Quz []string
}

func main() {
	// colocar cursor em cima do type e apertar cmd + . para preencher
	f := Foo{
		Bar: "",
		Baz: 0,
		Quz: []string{},
	}
}
