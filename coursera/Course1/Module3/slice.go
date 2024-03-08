package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 0, 3)
	for {
		var str string
		fmt.Printf("Enter an integer> ")
		fmt.Scanf("%s", &str)
		if strings.EqualFold("X", str) {
			break
		}
		x, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		sli = append(sli, x)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
