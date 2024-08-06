package main

import "fmt"

type Number interface {
	int | float64
}

func findMaxGeneric[Num Number](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func findMinGeneric[Num Number](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	maxInt := findMaxGeneric([]int{1, 32, 5, 8, 10, 11})
	fmt.Printf("max integer value: %v\n", maxInt)
	maxFloat := findMaxGeneric([]float64{1.5, 32.3, 5.6, 8.7, 10.4, 11.6})
	fmt.Printf("max integer value: %v\n", maxFloat)
	minInt := findMinGeneric([]int{1, 32, 5, 8, 10, 11})
	fmt.Printf("min integer value: %v\n", minInt)
	minFloat := findMinGeneric([]float64{1.5, 32.3, 5.6, 8.7, 10.4, 11.6})
	fmt.Printf("max integer value: %v\n", minFloat)
}
