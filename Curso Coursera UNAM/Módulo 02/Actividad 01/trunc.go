package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Give me an rational number and press enter:")
	fmt.Scan(&x)
	fmt.Printf("Your number truncated is %v\n", int(x))
}
