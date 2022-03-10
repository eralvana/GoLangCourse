package main

import "fmt"

func main() {
  var x int = 1
  var y int
  var ip *int

  fmt.Println(x,y,ip)

  ip = &x

  fmt.Println(x,y,ip)

  y = *ip

	fmt.Println(x,y,ip)

  ptr := new(int)

  fmt.Println(ptr)

  *ptr = 3

  fmt.Println(ptr,*ptr)

  fmt.Println(x+y+*ptr)
}
