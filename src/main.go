package main

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
)

func main() {
	// Création de l'inventaire
	inv := []inventory.Object{
		{"Potion", 3},
	}

	// Création du perso
	c1 := character.InitCharacter("Harold", "Elfe", 1, 100, 100, inv)

	// Affichage
	fmt.Println(c1)      // brut
	c1.DisplayInfo()     // via méthode
	c1.AccessInventory() // affiche inventaire

	c1.ActualPv = 50

	fmt.Print(c1.ActualPv)

}
