package main

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/combat"
	"projet-red_PIZZA-BATTLE/forgeron"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/marchand"
	"projet-red_PIZZA-BATTLE/structures"
)

func main() {
	var menuChoice int

	for {
		affichage.AffichageMenuDemarrage()
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
			c1 := structures.InitCharacter(1, inv, 10, 100, skillList, 100, 0, 100)
			// Création du personnage par l'utilisateur
			character.CharacterCreation(c1)
			// Initialiser l'ennemi
			e1 := structures.InitEnemy("Petit Giovanni", 100, 100, 5, "Facile", 100)

			for {
				affichage.AffichageMenuPrincipal()
				fmt.Scan(&menuChoice)

				switch menuChoice {
				case 1:
					// Afficher les infos du personnage
					character.DisplayCInfo(c1)
				case 2:
					// Menu Inventaire
					for {
						menuChoice = 0
						character.AccessInventory(c1)
						character.AccessEquipement(c1)
						character.AccessSkills(c1)
						// Menu Inventaire
						affichage.AffichageMenuInventaire()
						fmt.Scan(&menuChoice)
						switch menuChoice {
						case 1:
							// Utiliser une potion
							items.TakePot(c1)
						case 2:
							// Equiper un équipement
							character.EquipEquipment(c1)
						case 3:
							// Retour
						}
						if menuChoice == 3 {
							menuChoice = 0
							break
						}
					}
				case 3:
					// Test de combat (1v1)
					combat.TurnCombat1v1(c1, e1)
					combat.Combat(c1, e1)
				case 4:
					// Marchand
					marchand.Marchand(c1)
				case 5:
					// Forgeron
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
