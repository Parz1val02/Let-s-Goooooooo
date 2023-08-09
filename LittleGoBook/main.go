package main

import(
    "fmt"
    "os"
)

func main(){
    //Command line arguments
    if len(os.Args) !=2 {
        os.Exit(1)
    }
    fmt.Println("It's over", os.Args[1])

    //Assign multiples variables
    name, power := "Goku", 9000
    fmt.Printf("%s's power is over %d\n", name, power)
}
