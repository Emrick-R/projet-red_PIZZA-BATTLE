package forgeron

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/structures"
)

// Fonction utilitaire pour obtenir le minimum entre deux entiers
func min(a, b int) int {
	// Retourne le plus petit des deux entiers
	if a < b {
		return a
	}
	return b
}

// Forgeron permet au joueur de fabriquer des objets s'il a les matériaux nécessaires
func Forgeron(c *structures.Character) {
	// Initialisation des objets nécessaires à la fabrication et des objets fabriqués
	// Objets nécessaires
	wolfFur := structures.Object{Name: "Peau de Loup", Quantity: 1}
	trollSkin := structures.Object{Name: "Peau de Troll", Quantity: 1}
	boarLeather := structures.Object{Name: "Cuir de Sanglier", Quantity: 1}
	ravenFeather := structures.Object{Name: "Plume de Corbeau", Quantity: 1}

	// Objets fabriqués
	chapeauAventurier := structures.Object{Name: "Chapeau de l'aventurier", Quantity: 1}
	tuniqueAventurier := structures.Object{Name: "Tunique de l'aventurier", Quantity: 1}
	bottesAventurier := structures.Object{Name: "Bottes de l'aventurier", Quantity: 1}

	// Initialisation des variables pour les choix du joueur
	var forgeron_choice int
	var craft_confirmation int

	// Effacer l'écran
	fmt.Print("\033[H\033[2J")

	// Boucle principale du forgeron
	for {
		affichage.Separator()
		fmt.Println("⚒️ Bienvenue chez le Forgeron !")
		affichage.Separator()
		fmt.Println("Que veux tu frabriquer ?")
		fmt.Println("1 - 👒 Chapeau de l'aventurier : + 10 PV Max (1 Plume de Corbeau + 1 Cuir de Sanglier)")
		fmt.Println("2 - 👕 Tunique de l'aventurier : + 25 PV Max (2 Fourrures de Loup + 1 Peau de Troll)")
		fmt.Println("3 - 👖 Bottes de l'aventurier : + 15 PV Max (1 Fourrure de Loup + 1 Cuir de Sanglier)")
		fmt.Println("4 - ⬅️  RETOUR")
		fmt.Print("👉 Ton choix : ")
		fmt.Scan(&forgeron_choice)

		// Initialisation des compteurs pour chaque matériau
		var countRavenFeather, countBoarLeather, countWolfFur, countTrollSkin int

		// Comptage des matériaux dans l'inventaire du personnage
		for i := range c.Inventory {
			switch c.Inventory[i].Name {
			case "Plume de Corbeau":
				countRavenFeather = c.Inventory[i].Quantity
			case "Cuir de Sanglier":
				countBoarLeather = c.Inventory[i].Quantity
			case "Peau de Loup":
				countWolfFur = c.Inventory[i].Quantity
			case "Peau de Troll":
				countTrollSkin = c.Inventory[i].Quantity
			}
		}

		// Gestion des choix du joueur pour la fabrication
		switch forgeron_choice {
		case 1:
			// Fabrication du Chapeau de l'aventurier
			// Calcul du nombre d'objets fabriquables en fonction des matériaux disponibles
			itemcraftable := min(countRavenFeather, countBoarLeather)
			//Vérification si le joueur a les matériaux nécessaires
			if itemcraftable < 1 {
				fmt.Printf("\n❌ Tu n'as pas les objets requis pour fabriquer le Chapeau de l'Aventurier\n\n")
				// Il n'a pas les matériaux nécessaires, retour au menu
				break
			} else {
				// Il a les matériaux nécessaires, affiche le nombre d'objets fabriquables
				fmt.Printf("\n✅ Tu peux fabriquer %dx Chapeau de l'Aventurier, veux-tu en fabriquer 1 ?\n", itemcraftable)
			}

			// Demande de confirmation au joueur avant de fabriquer l'objet
			// 1 = Oui, 2 = Non
			fmt.Println("1 - Oui je suis sûr !")
			fmt.Println("2 - Non je ne suis pas sûr, retour !")
			fmt.Print("👉 Ton choix : ")
			fmt.Scan(&craft_confirmation)
			switch craft_confirmation {
			case 1:
				// Fabrication du chapeau
				fmt.Println("\n✅ Tu viens de fabriquer", itemcraftable, "x Chapeau de l'Aventurier")
				// Retrait des matériaux de l'inventaire
				inventory.RemoveInventory(c, ravenFeather)
				inventory.RemoveInventory(c, boarLeather)
				// Ajout du chapeau à l'inventaire
				inventory.AddInventory(c, chapeauAventurier)
			case 2:
				// Retour au menu précédent sans fabriquer l'objet
				fmt.Printf("\nRetour au menu précédent\n")
			}

		case 2:
			// Fabrication de la Tunique de l'aventurier
			//besoin de 2 fourrures de loup + 1 peau de troll
			itemcraftable := min(countWolfFur/2, countTrollSkin)
			if itemcraftable < 1 {
				fmt.Printf("\n❌ Tu n'as pas les objets requis pour Tunique de l'aventurier de l'Aventurier\n\n")
				break
			} else {
				fmt.Printf("\n✅ Tu peux fabriquer %dx Chapeau de l'Aventurier, veux-tu en fabriquer 1 ?\n", itemcraftable)
			}

			fmt.Println("1 - Oui je suis sûr !")
			fmt.Println("2 - Non je ne suis pas sûr, retour !")
			fmt.Print("👉 Ton choix : ")
			fmt.Scan(&craft_confirmation)
			switch craft_confirmation {
			case 1:
				// Fabrication de la tunique, retrait des matériaux et ajout à l'inventaire
				fmt.Printf("\n✅ Tu viens de fabriquer %d x Tunique de l'Aventurier\n\n", itemcraftable)
				inventory.RemoveInventory(c, wolfFur)
				inventory.RemoveInventory(c, wolfFur) //Retire 2 fourrures de loup
				inventory.RemoveInventory(c, trollSkin)
				inventory.AddInventory(c, tuniqueAventurier)
			case 2:
				// Retour au menu précédent sans fabriquer l'objet
				fmt.Printf("\nRetour au menu précédent\n\n")
			}

		case 3:
			// Fabrication des Bottes de l'aventurier
			itemcraftable := min(countWolfFur, countBoarLeather)
			if itemcraftable < 1 {
				fmt.Printf("\n❌ Tu n'as pas les objets requis pour Bottes de l'aventurier de l'Aventurier\n\n")
				break
			} else {
				fmt.Printf("\n✅ Tu peux fabriquer %dx Chapeau de l'Aventurier, veux-tu en fabriquer 1 ?\n", itemcraftable)
			}

			fmt.Println("1 - Oui je suis sûr !")
			fmt.Println("2 - Non je ne suis pas sûr, retour !")
			fmt.Print("👉 Ton choix : ")
			fmt.Scan(&craft_confirmation)
			switch craft_confirmation {
			case 1:
				// Fabrication des bottes, retrait des matériaux et ajout à l'inventaire
				fmt.Printf("\n✅ Tu viens de fabriquer %d x Bottes de l'Aventurier\n\n", itemcraftable)
				inventory.RemoveInventory(c, wolfFur)
				inventory.RemoveInventory(c, boarLeather)
				inventory.AddInventory(c, bottesAventurier)
			case 2:
				// Retour au menu précédent sans fabriquer l'objet
				fmt.Printf("\nRetour au menu précédent\n")
			}

		case 4:
			// Retour au menu précédent
		}
		if forgeron_choice == 4 {
			forgeron_choice = 0

			// Effacer l'écran
			fmt.Print("\033[H\033[2J")

			break
		}
	}
}
