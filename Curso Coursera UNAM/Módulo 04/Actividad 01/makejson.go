package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string
	Address string
}

func main() {
	// Primer instrucción para el usuario
	fmt.Println("Input your name and press enter:")
	x := bufio.NewReader(os.Stdin)
	y, erry := x.ReadString('\n')
	if erry != nil {
		fmt.Println("There is an error.")
		return
	}
	y = strings.TrimRight(y, "\n")

	fmt.Println("Input your address and press enter:")
	u := bufio.NewReader(os.Stdin)
	v, errv := u.ReadString('\n')
	if errv != nil {
		fmt.Println("There is an error.")
		return
	}
	v = strings.TrimRight(v, "\n")

	p := Person{
		Name:    y,
		Address: v,
	}

	r, errr := json.Marshal(p)
	if errr != nil {
		fmt.Println("There is an error.")
		return
	}

	fmt.Println(string(r))
}
