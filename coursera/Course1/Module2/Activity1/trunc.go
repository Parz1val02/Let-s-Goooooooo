package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var numStr string
	for {
		fmt.Printf("Enter a floating point number> ")
		fmt.Scanf("%s", &numStr)
		numFloat, _ := strconv.ParseFloat(numStr, 32)
		numTrunc := math.Trunc(numFloat)
		fmt.Printf("Truncated version of %.2f: %.0f\n", numFloat, numTrunc)
	}
}
