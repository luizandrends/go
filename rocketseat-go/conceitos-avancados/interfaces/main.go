package main

import (
	"fmt"

	"github.com/luizandrends/interfaces/bar"
)

type Animal interface {
	Sound() string
}

type Dog struct {}

type Cat struct {}

func (Dog) Sound() string {
	return "Au! Au!"
}

func (Cat) Sound() string {
	return "Meow! Meow!"
}

func (Dog) Interface() {
	fmt.Println("Dog Interface Called")
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func main() {
	dog := Dog{}
	cat := Cat{}
	whatDoesThisAnimalSay(dog)
	whatDoesThisAnimalSay(cat)

	bar.TakeFoo(dog)
}
