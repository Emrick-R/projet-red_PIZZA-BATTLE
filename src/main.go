package main

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/combat"
	"projet-red_PIZZA-BATTLE/forgeron"
	"projet-red_PIZZA-BATTLE/marchand"
	"projet-red_PIZZA-BATTLE/structures"
)

// main est le point d'entrée du programme
func main() {
	// Initialisation du choix du menu
	var menuChoice int

	// Boucle principale du jeu
	for {
		// Affichage du menu de démarrage
		affichage.AffichageMenuDemarrage()
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			// On lance une nouvelle partie

			// Initialisation des statistiques, de l'inventaire et des compétance de l'ennemi
			c1 := structures.InitCharacter()

			// Création du personnage par l'utilisateur
			character.CharacterCreation(c1)

			// Initialisation de l'ennemi
			e1 := structures.InitEnemy("Petit Giovanni", "Boss")

			// Boucle du menu principal
			for {
				affichage.AffichageMenuPrincipal()
				fmt.Scan(&menuChoice)

				switch menuChoice {
				case 1:
					// Afficher les infos du personnage
					character.DisplayCInfo(c1)
				case 2:
					// Menu Inventaire
					character.InventoryChoice(c1)
				case 3:
					// Combat en 1 contre 1
					combat.TurnCombat1v1(c1, e1)
				case 4:
					// Marchand
					marchand.Marchand(c1)
				case 5:
					// Forgeron
					forgeron.Forgeron(c1)
				case 6:
				default:
					// Choix invalide
					fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
				}
				// Reset de la variable menuChoice pour éviter les boucles infinies
				if menuChoice == 6 {
					menuChoice = 0
					// Retour au menu précédent (menu de démarrage)
					break
				}

			}
		case 2:
			// Quitter le jeu
			os.Exit(0)
		default:
			// Choix invalide
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}
	}
}
