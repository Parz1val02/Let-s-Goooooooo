package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strconv"
)

//structs
type Point struct{
	x int32
	y int32
}
type Circle struct{
	radius float64
	center *Point
}
type Student struct{
	name	string
	grades	[]int
	age		int
}
//methods
func (s *Student) setAge(age int){
	s.age = age
}
func (s *Student) getAverageGrade()  float32{
	sum := 0
	for _, v := range s.grades{
		sum += v
	}
	return float32(sum)/float32(len(s.grades))
}
//interfaces
type circle struct{
	radius	float64
}
type rectangle struct{
	width	float64
	height	float64
}
func (r *rectangle) area () float64{
	return r.width * r.height
}
func (c *circle) area () float64{
	return  math.Pow(c.radius,2) * math.Pi
}
type shape interface{
	area() float64
}
func getArea(s shape) float64{
	return s.area()
}

func main(){
	fmt.Println("Welcome to my quiz game!")
	//Implicit declaration
	var name string = "Rodro"
	//var int var1 = -1
	//var uint var2 = 2 //unsigned integer
	//var float64 var3 = 4.2 //define precision for float, also available float32
	//var ga bool = true
	fmt.Println(name)
	//Explicit declaration
	name2 := "Kipi"
	age := 21
	money := 33.3
	fmt.Printf("Hello %v, how are you doing\n", name2)
	fmt.Printf("You are %d years old and have %f soles in the bank\n", age, money)

	var ga string
	var age2 uint
	fmt.Print("Enter your name> ")
	fmt.Scanf("%s",&ga)
	fmt.Printf("Hello %s, welcome to the game\n",ga)
	fmt.Print("Enter your age> ")
	fmt.Scanf("%d",&age2)

	if age2 > 10 {
		fmt.Println("Tas basado mano")
	} else if age2 == 10 {
		fmt.Println("Puede ser puede ser")
	}else {
		fmt.Println("Sorry mano uwu")
	}

	//Scanner object
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born> ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2023\n", 2023 - input)

	ans := 10
	switch ans {
	case 1,2,3,4,5:
		fmt.Println("siu")
	case 6,7,8,9,10:
		fmt.Println("gaaaa")
	default:
			
	}

	//arrays
	var arr[5]int
	arr[0]=100
	fmt.Println(arr)
	arr2 := [3]int{4,5,6}
	var sum int = 0
	for i:=0;i<len(arr2);i++{
		sum += arr2[i]
	}
	fmt.Println(sum)
	arr2D:=[2][3]int{{1,2,3},{3,2,1}}
	fmt.Println(arr2D)

	//slices
	var x [5]int = [5]int{1,2,3,4,5}
	var s []int	= x[1:3]
	fmt.Println(cap(s))
	fmt.Println(s[:cap(s)])

	var a []int = []int{5,6,7,8,9}
	b := append(a,6)
	fmt.Println(b)
	c := make([]int, 5)
	fmt.Printf("%T\n",c)

	for i, element := range(b){
		for j:=i+1;j<len(b);j++{
			element2 := b[j]
			if element2==element{
				fmt.Println(element)
			}
		}
	}

	//maps
	var mp map[string]int = map[string]int{
		"apple":5,
		"pear":6,
		"orange":9,
	}
	fmt.Println(mp["apple"])
	mp["banana"]=900
	fmt.Println(mp)
	delete(mp, "apple")
	fmt.Println(mp)	
	val, ok := mp["apple"]
	fmt.Println(val, ok)

	mp2 := make(map[string]int)
	fmt.Println(mp2)

	//Function closures
	test :=  func(x int) int{
		fmt.Println("Hello world!")
		return x*-1
	}
	test2(test)	
	ga1 := returnfunc("goodbye")
	ga1()

	//Pointers and dereference operators
	bluebeetle:= "18 de agosto. solo en cines"
	owo := &bluebeetle
	fmt.Println(bluebeetle, owo)
	*owo = "ga"
	fmt.Println(bluebeetle, owo)

	toChange:="hello"
	fmt.Println(toChange)
	cahngeValue(&toChange)
	fmt.Println(toChange)

	var pointer *string = &toChange
	fmt.Println(*pointer, pointer, &pointer)

	//Structs
	var p1 Point = Point{1,2}
	fmt.Println(p1.x, p1.y)
	
	p2 := &Point{y:1}
	fmt.Println(p2)
	changeX(p2)
	fmt.Println(p2)
	
	c1 := Circle{4.56, &p1}
	fmt.Println(c1.center.x)
	
	//Methods
	s1 := Student{"Rodrigo", []int{17,20,20,18}, 21}
	s1.setAge(20)
	fmt.Println(s1)
	fmt.Println(s1.getAverageGrade())

	//Interfaces
	cga := circle{4.5}
	r1 := rectangle{4,5}
	shapes := []shape{&cga, &r1}
	for _, shape := range shapes{
		fmt.Println(getArea(shape))
	}
}
 func test2(myfunc func(int) int){
	defer fmt.Println("Usando defer")
	fmt.Println(myfunc(7))
}
func returnfunc(x string) func(){
	return func() {fmt.Println(x)}
}
func cahngeValue(str *string){
	*str = "changed!"
}
func changeX(pt *Point){
	pt.x = 100
}
