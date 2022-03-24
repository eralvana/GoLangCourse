package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	name    string
	address string
}

func main() {
	// Primer instrucción para el usuario
	fmt.Println("Input your name and press enter:")
	x := bufio.NewReader(os.Stdin)

	y, _ := x.ReadString('\n')

	y = strings.TrimRight(y, "\n")

	fmt.Println("Input your address and press enter:")
	u := bufio.NewReader(os.Stdin)

	v, _ := u.ReadString('\n')

	v = strings.TrimRight(v, "\n")
	p := map[string]string{
		"name":    y,
		"address": v,
	}

	r, _ := json.Marshal(p)

	fmt.Println(string(r))
}
