package main

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/structures"
)

func main() {
	var menuChoice int
	for {
		fmt.Println("======== Faites votre choix : ========")
		fmt.Println("1 - Commencer la partie")
		fmt.Println("2 - Quitter")
		fmt.Scan(&menuChoice)

		switch menuChoice {
		case 1:
			break

		case 2:
			os.Exit(0)
		}

	inv := []structures.Object{
		{"Potion", 3},
	}

	c1 := structures.InitCharacter("Harold", "Elfe", 1, 100, 100, inv)

	fmt.Println("======== Menu Personnage : ========")
	fmt.Println("1 - Afficher le personnage")
	fmt.Println("2 - Afficher l'inventaire")
	fmt.Println("3 - Prendre une potion")
	fmt.Scan(&menuChoice)

	switch menuChoice {
	case 1:
		character.DisplayInfo(c1)
	case 2:
		character.AccessInventory(c1)
	case 3:
		c1.ActualHp -= 50
		fmt.Println("HP avant potion :", c1.ActualHp)
		items.TakePot(c1)
		fmt.Println("HP après potion :", c1.ActualHp)
		character.AccessInventory(c1)
	}
}
