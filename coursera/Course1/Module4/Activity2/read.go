package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filePath string
	fmt.Print("Enter the name of the text file to read> ")
	fmt.Scanf("%s", &filePath)
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fileLines := make([]name, 0)

	for fileScanner.Scan() {
		var nameLine name
		nombres := strings.Split(fileScanner.Text(), " ")
		nameLine.fname = nombres[0]
		nameLine.lname = nombres[1]
		fileLines = append(fileLines, nameLine)
	}

	readFile.Close()

	for i, line := range fileLines {
		fmt.Printf("Name #%d\n", i)
		fmt.Println("First name> " + line.fname)
		fmt.Println("Last name> " + line.lname)
	}
}
