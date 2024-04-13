package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	arr := make([]int, 0)
	fmt.Print("Enter integers to sort> ")
	reader := bufio.NewReader(os.Stdin)
	request, _ := reader.ReadString('\n')
	request = strings.TrimSpace(request)

	integers := strings.Fields(request)
	n := len(integers)
	for _, item := range integers {
		i, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println("Error")
			break
		}
		arr = append(arr, i)
	}

	partSize := n / 4
	partitions := make([][]int, 4)
	for i := range partitions {
		start := i * partSize
		end := start + partSize
		if i == 3 {
			end = n
		}
		partitions[i] = arr[start:end]
	}

	var wg sync.WaitGroup
	wg.Add(4)

	for i := range partitions {
		go func(part []int) {
			defer wg.Done()
			fmt.Printf("Sorting> %v\n", part)
			sort.Ints(part)
		}(partitions[i])
	}

	wg.Wait()

	sorted := merge(partitions)

	fmt.Println("Sorted array> ", sorted)
}

func merge(arrs [][]int) []int {
	result := make([]int, 0, len(arrs[0])*len(arrs))
	for _, arr := range arrs {
		result = append(result, arr...)
	}
	sort.Ints(result)
	return result
}
