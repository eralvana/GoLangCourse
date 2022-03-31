package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	mapInfo := make(map[string]Animal)
	mapInfo["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	mapInfo["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	mapInfo["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	fmt.Printf("Indicate, in that order, the name of the animal (cow, bird or snake) and what you would like to know about it (eat, move or speak) and press enter:>")
	input := bufio.NewReader(os.Stdin)
	info, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("There is an error.")
		return
	}
	info = strings.TrimRight(info, "\n")
	slcinfo := strings.Split(info, " ")

	animalinfo := slcinfo[0]

	for animalinfo != "X" {

		if len(slcinfo) != 2 {
			fmt.Println("There are not two words.")

		} else {

			animalAsk := mapInfo[animalinfo]

			switch slcinfo[1] {
			case "eat":
				fmt.Printf("%v eat %v\n", slcinfo[0], animalAsk.Eat())
			case "move":
				fmt.Printf("%v move %v\n", slcinfo[0], animalAsk.Move())
			case "speak":
				fmt.Printf("%v speak %v\n", slcinfo[0], animalAsk.Speak())
			}
		}

		fmt.Printf("Try again! Indicate, in that order, the name of the animal (cow, bird or snake) and what you would like to know about it (eat, move or speak) and press enter or type X and press enter to exit:>")
		input = bufio.NewReader(os.Stdin)
		info, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("There is an error.")
			return
		}
		info = strings.TrimRight(info, "\n")
		slcinfo = strings.Split(info, " ")

		animalinfo = slcinfo[0]
		continue
	}
}
