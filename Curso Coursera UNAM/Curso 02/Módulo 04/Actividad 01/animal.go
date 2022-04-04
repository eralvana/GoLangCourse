package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Object struct {
	food       string
	locomotion string
	noise      string
}

func (ob Object) Eat() string {
	return ob.food
}

func (ob Object) Move() string {
	return ob.locomotion
}

func (ob Object) Speak() string {
	return ob.noise
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func main() {
	mapInfo := make(map[string]Object)
	mapInfo["cow"] = Object{food: "grass", locomotion: "walk", noise: "moo"}
	mapInfo["bird"] = Object{food: "worms", locomotion: "fly", noise: "peep"}
	mapInfo["snake"] = Object{food: "mice", locomotion: "slither", noise: "hsss"}

	fmt.Println("Indicate, in that order,")
	fmt.Println("1) What do you want?: create new animal (type new) or make a request (type request)")
	fmt.Println("2) The name of the animal that you want create or the name of the animal to make a request (cow, bird or snake)")
	fmt.Println("3) The properties to copy to the new animal (cow, bird or speak) or what you would like to know about it (eat, move or speak)")
	fmt.Println("Press enter:>")
	input := bufio.NewReader(os.Stdin)
	info, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("There is an error.")
		return
	}
	info = strings.TrimRight(info, "\n")
	slcinfo := strings.Split(info, " ")

	for slcinfo[0] != "X" {

		if len(slcinfo) != 3 {
			fmt.Println("There are not three words.")

		} else if slcinfo[0] == "new" {

			mapInfo[string(slcinfo[1])] = mapInfo[slcinfo[2]]

		} else if slcinfo[0] == "request" {

			animalAsk := mapInfo[string(slcinfo[1])]

			switch slcinfo[2] {
			case "eat":
				fmt.Printf("%v eat %v\n", slcinfo[1], animalAsk.Eat())
			case "move":
				fmt.Printf("%v move %v\n", slcinfo[1], animalAsk.Move())
			case "speak":
				fmt.Printf("%v speak %v\n", slcinfo[1], animalAsk.Speak())
			}
		} else {
			fmt.Println("The instruction is wrong.")
		}

		fmt.Println("Indicate, in that order,")
		fmt.Println("1) What do you want?: create new animal (type new) or make a request (type request)")
		fmt.Println("2) The name of the animal that you want create or the name of the animal to make a request (cow, bird or snake)")
		fmt.Println("3) The properties to copy to the new animal (cow, bird or speak) or what you would like to know about it (eat, move or speak)")
		fmt.Println("Press enter:>")
		input = bufio.NewReader(os.Stdin)
		info, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("There is an error.")
			return
		}
		info = strings.TrimRight(info, "\n")
		slcinfo = strings.Split(info, " ")

		continue
	}
}
