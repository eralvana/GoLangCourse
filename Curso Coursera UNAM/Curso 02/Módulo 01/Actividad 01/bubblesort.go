package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(slcint []int) {

	for i, _ := range slcint {
		min := IndexMinSlc(slcint[i:])
		for j := min; j > 0; j-- {
			Swap(slcint[i:], j-1)
		}
	}

	for index, value := range slcint {
		if index != len(slcint)-1 {
			fmt.Printf("%v, ", value)
			continue
		}
		fmt.Println("and the last one is", value)
	}
}

func Swap(slc []int, i int) {
	if i < 0 || i >= len(slc)-1 {
		fmt.Println("Your index is out of range")
	} else {
		value1 := slc[i]
		value2 := slc[i+1]

		slc[i] = value2
		slc[i+1] = value1

	}
}

func IndexMinSlc(slc []int) int {
	smallestNumber := slc[0]
	indmin := 0
	for index, value := range slc {
		if value < smallestNumber {
			smallestNumber = value
			indmin = index
		}
	}
	return indmin
}

func main() {
	var slc []int

	fmt.Println("Type in a sequence of up to 10 integers, when you are finished press enter:")
	input := bufio.NewReader(os.Stdin)
	inputstr, errinput := input.ReadString('\n')
	if errinput != nil {
		fmt.Println("There is an error.")
	}
	inputstr = strings.TrimRight(inputstr, "\n")
	slcinput := strings.Split(inputstr, " ")

	for _, v := range slcinput {
		item, erritem := strconv.Atoi(v)
		if erritem != nil {
			fmt.Println("There is an error")
			break
		}
		slc = append(slc, item)
	}

	if len(slcinput) > 10 {
		fmt.Println("You have entered more than 10 integers.")
	} else {
		fmt.Println("Your numbers ordered from least to greatest are")
		BubbleSort(slc)
	}
}
