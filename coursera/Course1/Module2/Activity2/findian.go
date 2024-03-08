package main

import (
	"fmt"
	s "strings"
)

func main() {
	var x string
	fmt.Printf("Enter a string> ")
	fmt.Scanf("%s", &x)
	x = s.ToLower(x)
	if s.Contains(x, "a") && s.HasPrefix(x, "i") && s.HasSuffix(x, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
