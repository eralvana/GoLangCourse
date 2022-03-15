package main

import (
  "bufio"
	"fmt"
  "os"
	"strings"
)

func main() {
  //Esto lo haremos para que cuando se tenga una entrada con espacios se borren y quede solamente como un arreglo de letras sin espacios

  x := bufio.NewReader(os.Stdin)
  fmt.Println("Give me a string and press enter:")
  y, _ := x.ReadString('\n')

  //Dado que y ahora es un string que sí considera espacios, verificamos que cumpla las condiciones. Aquí la condición del índice de las entradas hay que restarle dos pues el último es un cambio de línea.
  if strings.ToLower(string(y[0])) == "i" && strings.ToLower(string(y[len(y)-2])) == "n" && 0 < strings.Index(strings.ToLower(y),"a") {
    fmt.Println("Found!")
  } else {
    fmt.Println("Not Found!")
  }
}
