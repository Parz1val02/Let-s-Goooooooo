package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string
	jmap := make(map[string]string)
	fmt.Print("Enter a name> ")
	fmt.Scanf("%s", &name)
	jmap["name"] = name
	fmt.Print("Enter an address> ")
	fmt.Scanf("%s", &addr)
	jmap["address"] = addr
	barr, err := json.Marshal(jmap)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("JSON object: " + string(barr))
}
