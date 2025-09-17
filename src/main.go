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
				if menuChoice == 7 {
					menuChoice = 0
					// Retour au menu précédent (menu de démarrage)
					break
				}

			}
		case 2:
			// Easter Egg secret
			var returnChoice int

			// Effacer l'écran
			fmt.Print("\033[H\033[2J")

			fmt.Println("===============================================")
			fmt.Println("        *** EASTER EGG DÉCOUVERT! ***         ")
			fmt.Println("===============================================")

			fmt.Println("\n🍕 Pizza Battle 🍕")
			fmt.Println("Développé par les légendaires goateurs de pizza:")
			fmt.Println("Emrick Rivet & Harold François")

			fmt.Println("\n-----------------------------------------------")
			fmt.Println("RÉFÉRENCES CACHÉES DANS LE JEU:")
			fmt.Println("-----------------------------------------------")

			fmt.Println("\nÉtape 2 - Références musicales:")
			fmt.Println("ABBA")

			fmt.Println("\nÉtape 3 - Références cinématographiques:")
			fmt.Println("Steven Spielberg")

			fmt.Println("\n===============================================")
			fmt.Println("Appuyez sur 0 pour revenir au menu principal")
			fmt.Println("===============================================")

			// Attente de la saisie utilisateur pour retourner au menu
			for {
				fmt.Print("Votre choix: ")
				fmt.Scan(&returnChoice)
				if returnChoice == 0 {
					break
				} else {
					fmt.Println("Appuyez sur 0 pour revenir au menu principal")
				}
			}
		case 3:
			// Quitter le jeu
			os.Exit(0)
		default:
			// Choix invalide
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}
	}
}
