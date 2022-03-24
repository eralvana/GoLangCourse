package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Fname string
	Lname string
}

func main() {
	fmt.Println("Input the name of the file that have the list of names and press enter:")
	x := bufio.NewReader(os.Stdin)
	y, erry := x.ReadString('\n')
	if erry != nil {
		fmt.Println("There is an error.")
	}
	y = strings.TrimRight(y, "\n")

	readFile, errRF := os.Open(y)
	if errRF != nil {
		fmt.Println("There is an error.")
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	slc := make(map[string][]string)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		wordinaline := strings.Fields(line)
		slc[line] = []string{wordinaline[0], wordinaline[1]}
	}
	readFile.Close()

	var slcstruc []Person

	for _, v := range slc {
		slcstruc = append(slcstruc, Person{
			Fname: v[0],
			Lname: v[1],
		})
	}
	fmt.Println("The slice of structurs is", slcstruc, "and the list is")
	for k, _ := range slc {
		fmt.Println(slc[k][0], slc[k][1])

	}
}
