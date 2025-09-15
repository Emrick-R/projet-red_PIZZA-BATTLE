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
		fmt.Println("1 - Commencer une nouvelle partie")
		fmt.Println("2 - Quitter")
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			HpPot := structures.Object{Name: "Potion de Vie"}
			inv := []structures.Object{
				{Name: HpPot.Name, Quantity: 3},
			}
			punch := structures.InitSkill("Coup de poing", 10)
			skillList := []structures.Skill{
				{Name: punch.Name, Damage: punch.Damage},
			}
			c1 := structures.InitCharacter("", "Elfe", 1, 100, 100, inv, 10, 100, skillList)
			character.CharacterCreation(c1)
			e1 := structures.InitEnemy("Giovanni", 100, 100)

			for {
				fmt.Println("\n======== Menu Personnage : ========")
				fmt.Println("1 - Afficher le personnage")
				fmt.Println("2 - Afficher l'inventaire")
				fmt.Println("3 - Marchand")
				fmt.Println("4 - Test de combat : Utiliser une potion de poison")
				fmt.Println("5 - RETOUR")
				fmt.Scan(&menuChoice)

				switch menuChoice {
				case 1:
					character.DisplayCInfo(c1)
				case 2:
					for {
						menuChoice = 0
						fmt.Println("======== Menu Inventaire : ========")
						character.AccessInventory(c1)
						fmt.Println("1 - Utiliser une potion")
						fmt.Println("2 - RETOUR")
						fmt.Scan(&menuChoice)
						switch menuChoice {
						case 1:
							items.TakePot(c1)
						case 2:
						}
						if menuChoice == 2 {
							menuChoice = 0
							break
						}
					}
				case 3:
					marchand.Marchand(c1)
				case 4:
					fmt.Printf("\nLe petit enfant Giovanni apparaît devant toi !\n-50 PV\n")
					character.DisplayEInfo(e1)

					fmt.Println("Tu lance une potion de poison !")
					items.ThrowPoisonPot(c1, e1)
					character.DisplayEInfo(e1)

				case 5:
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
