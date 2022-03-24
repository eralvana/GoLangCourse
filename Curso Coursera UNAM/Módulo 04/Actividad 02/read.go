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

	var slc []Person

	for fileScanner.Scan() {
		line := fileScanner.Text()
		wordinaline := strings.Fields(line)
		slc = append(slc, Person{
			Fname: wordinaline[0],
			Lname: wordinaline[1],
		})
	}
	readFile.Close()

	fmt.Println("The slice of structurs is", slc, "and the list is")
	for k := range slc {
		fmt.Println(slc[k].Fname, slc[k].Lname)

	}
}
