package main

import (
	"fmt"
	"time"
)

func fib(number float64, ch chan string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}
	ch <- fmt.Sprintf("Fib(%v): %v", number, x)
	return
}

func main() {
	size := 15
	ch := make(chan string, size)
	start := time.Now()

	for i := 1; i < 15; i++ {
		go fib(float64(i), ch)
	}

	for i := 1; i < 15; i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
