package main

import (
	"fmt"
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
	var size int
	fmt.Print("Enter the number of integers that will be sorted (max 10)> ")
	fmt.Scanf("%d", &size)
	if size < 10 {
		for i := 0; i < size; i++ {
			var num int
			fmt.Print("Type a number> ")
			fmt.Scanf("%d", &num)
			num_slice = append(num_slice, num)
		}
		BubbleSort(num_slice)
		fmt.Println(num_slice)
	} else {
		fmt.Println("Maximum 10 integers")
	}
}
