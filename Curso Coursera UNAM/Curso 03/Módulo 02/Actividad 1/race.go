package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	// chan to make sure operation has been made by one of goroutine
	exit := make(chan interface{})

	go (func() {
		// time.Sleep to generate time for race
		time.Sleep(time.Second)
		x += 2
		exit <- 1
	})()
	go (func() {
		time.Sleep(time.Second)
		x += 3
		exit <- 1
	})()
	<-exit
	//  Every time you run it'll be 3 or 4, because of which goroutine starts first
	fmt.Println(x)
}
