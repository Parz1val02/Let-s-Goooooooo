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
