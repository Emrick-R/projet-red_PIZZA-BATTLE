package main

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/items"
)

func main() {

	//Nouveau menu
	menuChoice := 0
	fmt.Print("======== Faites votre choix :======== \n")
	fmt.Print("1 - Commencer la partie \n")
	fmt.Print("2 - Quitter \n")
	fmt.Scan(&menuChoice)

	switch menuChoice {
	case 1:
		// Création de l'inventaire
		inv := []inventory.Object{
			{"Potion", 3},
		}

		// Création du perso
		c1 := character.InitCharacter("Harold", "Elfe", 1, 100, 100, inv)

		//Nouveau menu
		menuChoice := 0
		fmt.Print("======== Faites votre choix :======== \n")
		fmt.Print("1 - Afficher le parsonage \n")
		fmt.Print("2 - Afficher l'inventaire \n")
		fmt.Print("3 - Retour \n")
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			c1.DisplayInfo() // affiche la fishe de perso

		case 2:
			c1.AccessInventory() // affiche inventaire

		case 3:
			//retour au menu précédent

			// Take pot test
			c1.ActualHp -= 50

			fmt.Println(c1.ActualHp)

			items.TakePot(c1)

			fmt.Println(c1.ActualHp)

			c1.AccessInventory()
		}
	}

}
