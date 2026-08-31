package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++{
			ch <- i
		}
		close(ch)
		fmt.Println("Escrita finalizada!")
	}()
	
	// <-ch
	// <-ch
	// valor := <-ch
	time.Sleep(time.Second * 1)
	for valor := range ch {
		fmt.Println("Leitura ", valor)
	}
}
