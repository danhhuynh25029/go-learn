package main

import (
	"fmt"
)

func worker(name int, a, b chan int) {
	for i := range a {
		fmt.Println("worker %d, job %d", name, i)
		b <- i * 2
	}
}
func main() {
	a := make(chan int, 3)
	b := make(chan int, 3)
	for i := 0; i < 2; i++ {
		go worker(i, a, b)
	}
	for i := 0; i <= 2; i++ {
		a <- i
	}
	defer close(a)
	for i := 0; i <= 2; i++ {
		fmt.Println(<-b)
	}
}
