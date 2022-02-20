package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var guess int

	fmt.Println("Enter a number from 1 to 10")
	fmt.Scanln(&guess)

	rand.Seed(time.Now().UnixNano())

	min, max := 1, 10

	answer := rand.Intn(max-min+1) + min

	var unopened bool

	if guess == answer {
		fmt.Println("Congratulations! You've opened the chest!")
		treasureChest()
		unopened = false
	} else {
		fmt.Println("The chest remains locked..")
		unopened = true
	}

	for unopened {
		main()
		if unopened != true {
			break
		}
	}

}

func treasureChest() {

	gold := rand.Intn(10000)

	fmt.Printf("You've found %d gold pieces!", gold)

	a := []string{"[1]Deck of Many Things", "[2]Dragonscale Armor", "[3]Siryu's phone number"}
	fmt.Println(" The treasure chest also contains ", a)

	return
}
