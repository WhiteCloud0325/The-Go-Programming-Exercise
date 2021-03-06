package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
