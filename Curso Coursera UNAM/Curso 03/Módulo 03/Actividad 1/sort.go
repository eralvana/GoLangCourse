package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Sort(slc []int) []int {
	input := make([]int, len(slc), len(slc))
	for i := 0; i < len(slc); i++ {
		input[i] = slc[i]
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	return input
}

func main() {
	var slc []int

	fmt.Println("Type a sequence of integers, when you're done press enter:")
	input := bufio.NewReader(os.Stdin)
	inputstr, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("There is an error.")
	}
	inputstr = strings.TrimRight(inputstr, "\n")
	slcinput := strings.Split(inputstr, " ")

	for _, v := range slcinput {
		item, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Some value is not a number.")
			break
		}
		slc = append(slc, item)
	}

	div := len(slc) / 4
	rest := len(slc) % 4
	map4slc := make(map[int][]int)

	if len(slc) <= 4 {
		fmt.Printf("Your slice is %v and sorted is %v\n", slc, Sort(slc))
	} else {
		switch rest {
		case 0:
			map4slc[0] = slc[:div]
			map4slc[1] = slc[div : 2*div]
			map4slc[2] = slc[2*div : 3*div]
			map4slc[3] = slc[3*div:]
		case 1:
			map4slc[0] = slc[:div+1]
			map4slc[1] = slc[div+1 : 2*div+1]
			map4slc[2] = slc[2*div+1 : 3*div+1]
			map4slc[3] = slc[3*div+1:]
		case 2:
			map4slc[0] = slc[:div+1]
			map4slc[1] = slc[div+1 : 2*div+2]
			map4slc[2] = slc[2*div+2 : 3*div+2]
			map4slc[3] = slc[3*div+2:]
		case 3:
			map4slc[0] = slc[:div+1]
			map4slc[1] = slc[div+1 : 2*div+2]
			map4slc[2] = slc[2*div+2 : 3*div+3]
			map4slc[3] = slc[3*div+3:]

		}

		channel := make(chan []int, 2)
		wg1 := sync.WaitGroup{}
		wg1.Add(2)
		go func() {
			channel <- Sort(map4slc[0])
			wg1.Done()
		}()

		go func() {
			channel <- Sort(map4slc[1])
			wg1.Done()
		}()
		wg1.Wait()

		slc1Part1 := <-channel
		slc1Part2 := <-channel

		slc1 := append(slc1Part1, slc1Part2...)

		wg2 := sync.WaitGroup{}
		wg2.Add(2)
		go func() {
			channel <- Sort(map4slc[2])
			wg2.Done()
		}()

		go func() {
			channel <- Sort(map4slc[3])
			wg2.Done()
		}()
		wg2.Wait()

		slc2Part1 := <-channel
		slc2Part2 := <-channel

		slc2 := append(slc2Part1, slc2Part2...)

		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			channel <- Sort(slc1)
			wg.Done()
		}()

		go func() {
			channel <- Sort(slc2)
			wg.Done()
		}()
		wg.Wait()

		slcPart1 := <-channel
		slcPart2 := <-channel

		slcsort := append(slcPart1, slcPart2...)

		fmt.Printf("Your slice is %v and sorted is %v\n", slc, Sort(slcsort))
	}
}
