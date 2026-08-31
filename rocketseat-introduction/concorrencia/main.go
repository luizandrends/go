package main

import (
	"fmt"
	"time"
)

func exibirMsg() {
	fmt.Println("Olá de uma goroutine")
}

func main() {
	go exibirMsg()
	time.Sleep(1 * time.Second)
	fmt.Println("Olá main function!")
}