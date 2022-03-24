package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Input the name of the file that have the list of names and press enter:")
	x := bufio.NewReader(os.Stdin)
	y, _ := x.ReadString('\n')
	y = strings.TrimRight(y, "\n")
	data, error := ioutil.ReadFile(y)

	if error != nil {
		fmt.Println("There is a problem with the name file or the file")
	}

	s := string(data)
	//first, _ := s.ReadString('\n')
	//f, _ := first.ReadString(' ')
	//l, _ := first.ReadString('\n')
	//p := Person{fname: f, lname: l}
	//fmt.Println(p)
	fmt.Println(s)
}
