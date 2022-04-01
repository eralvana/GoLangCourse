package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return (a/2)*t*t + v*t + s
	}
}

func main() {
	var acceleration float64
	fmt.Println("Define acceleration please?")
	fmt.Scan(&acceleration)

	var velocity float64
	fmt.Println("Define velocity please?")
	fmt.Scan(&velocity)

	var displacement float64
	fmt.Println("Define displacement please?")
	fmt.Scan(&displacement)

	var time float64
	fmt.Println("Define time please?")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println("The result is ", fn(time))
}
