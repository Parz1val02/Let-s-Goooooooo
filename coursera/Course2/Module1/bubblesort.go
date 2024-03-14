package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(sli []int, index int) {
	temp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = temp
}

func BubbleSort(sli []int) {
	var swapped bool
	for i := 0; i < len(sli)-1; i++ {
		swapped = false
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	num_slice := make([]int, 0, 10)
	var integers string
	fmt.Print("Enter integers to be sorted separated by a comma i.e. 2,45,67> ")
	fmt.Scanf("%s", &integers)
	nums := strings.Split(integers, ",")
	if len(nums) < 10 {
		for _, item := range nums {
			i, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println("Error")
				break
			}
			num_slice = append(num_slice, i)
		}
		BubbleSort(num_slice)
		fmt.Println(num_slice)
	} else {
		fmt.Println("Maximum 10 integers")
	}
}
