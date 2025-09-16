package main

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/forgeron"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/marchand"
	"projet-red_PIZZA-BATTLE/structures"
)

func main() {
	var menuChoice int

	for {
		affichage.AffichageMenuPrincipal()
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			// Initialiser l'inventaire de départ
			HpPot := structures.Object{Name: "Potion de Vie"}
			inv := []structures.Object{
				{Name: HpPot.Name, Quantity: 3},
			}
			// Initialiser le skill de départ
			punch := structures.InitSkill("Coup de poing", 10)
			skillList := []structures.Skill{
				{Name: punch.Name, Damage: punch.Damage},
			}
			// Initialiser le personnage
			c1 := structures.InitCharacter(1, inv, 10, 100, skillList)
			character.CharacterCreation(c1)
			e1 := structures.InitEnemy("Giovanni", 100, 100, 5, "Facile")

			for {
				affichage.AffichageMenuPersonnage()
				fmt.Scan(&menuChoice)

				switch menuChoice {
				case 1:
					character.DisplayCInfo(c1)
				case 2:
					for {
						menuChoice = 0
						character.AccessInventory(c1)
						character.AccessEquipement(c1)
						character.AccessSkills(c1)
						affichage.AffichageMenuInventaire()
						fmt.Scan(&menuChoice)
						switch menuChoice {
						case 1:
							items.TakePot(c1)
						case 2:
							character.EquipEquipment(c1)
						case 3:
						}
						if menuChoice == 3 {
							menuChoice = 0
							break
						}
					}
				case 3:
					fmt.Printf("\nLe petit enfant Giovanni apparaît devant toi !\n-50 PV\n")
					character.DisplayEInfo(e1)

					fmt.Println("Tu lances une potion de poison !")
					items.ThrowPoisonPot(c1, e1)
					character.DisplayEInfo(e1)
				case 4:
					marchand.Marchand(c1)
				case 5:
					forgeron.Forgeron(c1)
				case 6:
				}
				if menuChoice == 6 {
					menuChoice = 0
					break
				}

			}
		case 2:
			os.Exit(0)
		}
	}
}
