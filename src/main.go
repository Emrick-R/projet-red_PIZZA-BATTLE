package main

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/marchand"
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
			PotionDeVie := structures.Object{"Potion de Vie", 1}
			inv := []structures.Object{
				{PotionDeVie.Name, 3},
			}

			c1 := structures.InitCharacter("Harold", "Elfe", 1, 100, 100, inv)
			for {
				fmt.Println("======== Menu Personnage : ========")
				fmt.Println("1 - Afficher le personnage")
				fmt.Println("2 - Afficher l'inventaire")
				fmt.Println("3 - Marchand")
				fmt.Println("4 - Test de combat : Prendre une potion")
				fmt.Println("5 - RETOUR")
				fmt.Scan(&menuChoice)

				switch menuChoice {
				case 1:
					character.DisplayInfo(c1)
				case 2:
					character.AccessInventory(c1)
				case 3:
					marchand.Marchand(c1)
				case 4:
					c1.ActualHp -= 50
					fmt.Printf("\nPV avant potion : %d\n\n", c1.ActualHp)
					items.TakePot(c1)
					character.AccessInventory(c1)
				case 5:
					break
				}
				if menuChoice == 5 {
					menuChoice = 0
					break
				}
			}
		case 2:
			os.Exit(0)
		}
	}

}
