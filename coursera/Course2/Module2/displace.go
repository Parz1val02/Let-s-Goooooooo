package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v_0, s_0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v_0*t + s_0
	}
	return fn
}

func main() {
	var acceleration, init_velocity, init_displacement, time float64
	fmt.Print("Let's calculate displacement using a formula\n")
	fmt.Print("Enter the value for acceleration> ")
	fmt.Scanf("%f", &acceleration)
	fmt.Print("Enter the value for initial velocity> ")
	fmt.Scanf("%f", &init_velocity)
	fmt.Print("Enter the value for initial displacement> ")
	fmt.Scanf("%f", &init_displacement)
	fn := GenDisplaceFn(acceleration, init_velocity, init_displacement)

	fmt.Print("Enter the value of time in order to calculate the displacement> ")
	fmt.Scanf("%f", &time)
	fmt.Printf("Displacement after %.2f seconds> %.2f\n", time, fn(time))
}
