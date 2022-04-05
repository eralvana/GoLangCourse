package main

import (
	"fmt"
	"time"
)

// This is an example of race condition 2 goroutines tries to read and write number and there is no access control.

var number int = 10
var value int = 40

func read() {
	for {
		var val int = number
		if val%10 == 0 {
			value = value - number + 1
		}
	}
}

func write() {
	for {
		number = number + 1
	}
}

func main() {
	for i := number; i < value; i++ {
		fmt.Println(number, value)
		go read()
		fmt.Println(number, value)
		go write()
		time.Sleep(time.Second)
	}
}
