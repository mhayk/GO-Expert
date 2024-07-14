package main

import "fmt"

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received: %d\n", workerId, x)
	}
}

func main() {
	data := make(chan int)
	// threads quantity
	QtdWorkers := 1000000

	// go worker(1, data)
	// go worker(2, data)

	// initialize workers
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000000; i++ {
		data <- i
	}
}
