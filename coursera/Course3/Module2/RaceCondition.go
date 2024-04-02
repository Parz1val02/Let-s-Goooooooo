package main

import (
	"fmt"
	"time"
)

var bestLinuxDistro = []byte("               ")

func main() {
	fmt.Println("Which is the best Linux distro?>")

	go writeText("Arch  ")
	go writeText("Ubuntu")

	time.Sleep(1 * time.Second)

	fmt.Println(string(bestLinuxDistro))
}

func writeText(newOrder string) {
	orderAsBytes := []byte(newOrder)
	for index, b := range orderAsBytes {
		bestLinuxDistro[index] = b
		time.Sleep(10 * time.Millisecond)
	}
}

/*
Explanation>
A race condition is a situation in which the outcome of a program depends on non-deterministic interleaving.
In other words, it happens when two or more threads communicate by accessing shared data and try to change it at the same time.
Because the scheduler swaps between threads at any time, the order in which the threads will attempt to access the shared data is unknown.
Therefore, the end result of the program is unpredictable.
In Go, two goroutines attempt to execute the same function with different arguments at almost the same time, generating a non-deterministic outcome.
*/
