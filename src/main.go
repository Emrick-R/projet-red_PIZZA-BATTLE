package main

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
)

func main() {

	choice := 0
	fmt.Print("======== Faites votre choix :======== \n")
	fmt.Print("1 - Commencer la partie \n")
	fmt.Print("2 - Quitter \n")
	fmt.Scan(&choice)

	switch choice {
	case 1:

		// Création de l'inventaire
		inv := []inventory.Object{
			{"Potion", 3},
		}

		// Création du perso
		c1 := character.InitCharacter("Harold", "Elfe", 1, 100, 100, inv)
		c1.DisplayInfo()
	case 2:
		fmt.Print("Ariverderci !")

	}
}
