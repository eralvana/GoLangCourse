package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Esto lo haremos para que cuando se tenga una entrada con espacios se borren y quede solamente como un arreglo de letras con espacios

	x := bufio.NewReader(os.Stdin)
	fmt.Println("Give me a string and press enter:")
	y, _ := x.ReadString('\n')

	//Dado que y ahora es un string que sC- considera espacios, verificamos que cumpla las condiciones. AquC- la condiciC3n del C-ndice de las entradas hay que restarle dos pues el C:ltimo es un cambio de lC-nea.
	if strings.ToLower(string(y[0])) == "i" && strings.ToLower(string(y[len(y)-2])) == "n" && 0 < strings.Index(strings.ToLower(y), "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
