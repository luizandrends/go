package main

type Animal interface {
	Sound() string
}

type Dog struct {}

type Cat struct {}

func (d *Dog) Sound() string {
	return "Au! Au!"
}

func (d *Cat) Sound() string {
	return "Meow! Meow!"
}

func takeAnimal(a Animal) {
	switch t := a.(type) {
	case *Dog:
		t.Sound()
	case *Cat:
		
	}

}

func main() {
}
