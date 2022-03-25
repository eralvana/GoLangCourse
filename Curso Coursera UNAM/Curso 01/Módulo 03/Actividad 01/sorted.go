package main

import "fmt"

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
	ex1 := []int{6, 7, 1, 2}
	ex2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	ex3 := []int{1000, 2000, 1300, 999, 0, -3}
	ex4 := []int{2, 4, 6, 8, 10, 1, 3, 5, 7}

	fmt.Println(sortedSlice(ex1))
	fmt.Println(sortedSlice(ex2))
	fmt.Println(sortedSlice(ex3))
	fmt.Println(sortedSlice(ex4))
}
