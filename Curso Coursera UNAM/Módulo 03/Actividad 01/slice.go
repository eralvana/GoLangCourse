package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func changeLower(l []int, n int) []int {
	slcLower := make([]int, len(l), len(l))
	slcLower[0] = n
	for i := 1; i < len(l); i++ {
		slcLower[i] = l[i-1]
	}
	return slcLower
}

func intermediate(l []int, n int, v int) []int {
	slc := make([]int, len(l), len(l))

	for i := 0; i < n; i++ {
		slc[i] = l[i]
	}
	slc[n] = v
	for i := n + 1; i < len(l); i++ {
		slc[i] = l[i-1]
	}
	return slc
}

func sortedSlice(s []int) (ord []int) {
	ord = make([]int, len(s), len(s))
	ord[0] = s[0]

	for i := 1; i < len(s); i++ {

		switch {
		case s[i] < ord[0]:
			ord = changeLower(ord, s[i])
		case s[i] >= ord[i-1]:
			ord[i] = s[i]
		default:
			for j := 0; j <= i; j++ {
				if s[i] >= ord[j] {
					continue
				} else {
					ord = intermediate(ord, j, s[i])
					break
				}
			}
		}
	}
	return ord
}

func main() {
	// AquC- tendremos el arreglo de enteros que nos irC!n proporcionando
	var s []int

	// Primer instrucciC3n para el usuario
	fmt.Println("Give me some integer numbers and press enter between each, when you are finished write X and press enter:")
	x := bufio.NewReader(os.Stdin)

	for {
		y, err := x.ReadString('\n')
		if err != nil {
			continue
		}

		y = strings.TrimRight(y, "\n")

		if y == "X" {
			break
		}

		z, err := strconv.Atoi(y)
		if err != nil {
			fmt.Println("Sorry, this is not a number.")
			continue
		}

		s = append(s, z)
	}
	fmt.Printf("Your slice is %v and the sorted slice is %v\n", s, sortedSlice(s))
}
