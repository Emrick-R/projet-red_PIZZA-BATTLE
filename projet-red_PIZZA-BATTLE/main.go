package main

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
)

func main() {
	c1 := character.InitCharacter("Harold", "Elfe", 1, 100, inventory.Object{"Potions", 3})
	fmt.Println(c1)

}

func IsDead() {
	return 0
}
