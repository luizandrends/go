package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++{
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func sayWorld() {
	for i := 0; i < 3; i++{
		fmt.Println("World")
		time.Sleep(150 * time.Millisecond)
	}
}


func main() {
	go sayHello()
	go sayWorld()
	time.Sleep(1 * time.Second)
}