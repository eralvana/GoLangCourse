package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(acc, inivel, inidis float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5)*acc*t*t + inivel*t + inidis
	}
	return fn
}

func main() {
	fmt.Println("Input the value of acceleration and press enter:")
	accelerationinput := bufio.NewReader(os.Stdin)
	acceleration, erracceleration := accelerationinput.ReadString('\n')
	if erracceleration != nil {
		fmt.Println("There is an error.")
		return
	}
	acceleration = strings.TrimRight(acceleration, "\n")
	accval, erraccval := strconv.ParseFloat(acceleration, 5)
	if erraccval != nil {
		fmt.Println("Invalid value.")
	}

	fmt.Println("Input the value of initial velocity and press enter:")
	initialvelocityinput := bufio.NewReader(os.Stdin)
	initialvelocity, errinitialvelocity := initialvelocityinput.ReadString('\n')
	if errinitialvelocity != nil {
		fmt.Println("There is an error.")
		return
	}
	initialvelocity = strings.TrimRight(initialvelocity, "\n")
	inivelval, errinivelval := strconv.ParseFloat(initialvelocity, 5)
	if errinivelval != nil {
		fmt.Println("Invalid value.")
	}

	fmt.Println("Input the value of initial displacement and press enter:")
	initialdisplacementinput := bufio.NewReader(os.Stdin)
	initialdisplacement, errinitialdisplacement := initialdisplacementinput.ReadString('\n')
	if errinitialdisplacement != nil {
		fmt.Println("There is an error.")
		return
	}
	initialdisplacement = strings.TrimRight(initialdisplacement, "\n")
	inidisval, errinidisval := strconv.ParseFloat(initialdisplacement, 5)
	if errinidisval != nil {
		fmt.Println("Invalid value.")
	}

	fn := GenDisplaceFn(accval, inivelval, inidisval)

	fmt.Printf("Thank you! Your function to calculate displacement at the time t is calculated. Now, input the value of time and press enter:")
	timeinput := bufio.NewReader(os.Stdin)
	time, errtime := timeinput.ReadString('\n')
	if errtime != nil {
		fmt.Println("There is an error.")
		return
	}
	time = strings.TrimRight(time, "\n")
	timevalue, errtimevalue := strconv.ParseFloat(time, 5)
	if errtimevalue != nil {
		fmt.Println("Invalid value.")
	}
	fmt.Println(fn(timevalue))
}
