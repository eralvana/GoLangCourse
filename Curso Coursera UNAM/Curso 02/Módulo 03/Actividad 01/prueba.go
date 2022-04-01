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
	sound      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.sound)
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", sound: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", sound: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", sound: "hsss"}

	animals := map[string]*Animal{"cow": &cow, "bird": &bird, "snake": &snake}

	for {
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := strings.ToLower(scanner.Text())
		animalName, action, found := strings.Cut(userInput, " ")

		if found && animals[animalName] != nil {
			switch strings.ToLower(action) {
			case "eat":
				animals[animalName].Eat()
			case "move":
				animals[animalName].Move()
			case "speak":
				animals[animalName].Speak()
			}
		}
	}
}
