package forgeron

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/structures"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Forgeron(c *structures.Character) {
	wolfFur := structures.Object{Name: "Peau de Loup", Quantity: 1}
	trollSkin := structures.Object{Name: "Peau de Troll", Quantity: 1}
	boarLeather := structures.Object{Name: "Cuir de Sanglier", Quantity: 1}
	ravenFeather := structures.Object{Name: "Plume de Corbeau", Quantity: 1}

	var forgeron_choice int
	var craft_confirmation int
	chapeauAventurier := structures.Object{Name: "Chapeau de l'aventurier", Quantity: 1}
	tuniqueAventurier := structures.Object{Name: "Tunique de l'aventurier", Quantity: 1}
	bottesAventurier := structures.Object{Name: "Bottes de l'aventurier", Quantity: 1}
	for {
		fmt.Println("======== Forgeron : ========")
		fmt.Println("Bonjour je suis le Forgeron, quel est votre choix ?")
		fmt.Println("1 - Fabrication : Chapeau de l'aventurier (1 Plume de Corbeau + 1 Cuir de Sanglier)")
		fmt.Println("2 - Fabrication : Tunique de l'aventurier (2 Fourrures de Loup + 1 Peau de Troll)")
		fmt.Println("3 - Fabrication : Bottes de l'aventurier (1 Fourrure de Loup + 1 Cuir de Sanglier)")
		fmt.Println("4 - RETOUR")
		fmt.Scan(&forgeron_choice)

		var countRavenFeather, countBoarLeather, countWolfFur, countTrollSkin int

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

		switch forgeron_choice {
		case 1:
			itemcraftable := min(countRavenFeather, countBoarLeather)
			if itemcraftable < 1 {
				fmt.Println("Tu n'as pas les items requis pour fabriquer le Chapeau de l'Aventurier")
			} else {
				fmt.Println("Tu peux fabriquer", itemcraftable, "x Chapeau de l'Aventurier")
			}

			if inventory.CheckMaxInventory(c) {
				fmt.Scan(&craft_confirmation)
				fmt.Println("1 - Oui je suis sûr !")
				fmt.Println("2 - Non je ne suis pas sûr, retour !")
				switch craft_confirmation {
				case 1:
					inventory.AddInventory(c, chapeauAventurier)
					fmt.Println("Tu viens de fabriquer", itemcraftable, "x Chapeau de l'Aventurier")
					inventory.RemoveInventory(c, ravenFeather)
					inventory.RemoveInventory(c, boarLeather)
				case 2:
					fmt.Println("Retour au menu précédent")
					break
				}
			} else {
				fmt.Printf("\nTu n'as pas de places dans ton inventaire, il est plein...\n\n")
				return
			}

		case 2:
			itemcraftable := min(countWolfFur/2, countTrollSkin) //besoin de 2 fourrures de loup
			if itemcraftable < 1 {
				fmt.Println("Tu n'as pas les items requis pour Tunique de l'aventurier de l'Aventurier")
			} else {
				fmt.Println("Tu peux fabriquer", itemcraftable, "x Tunique de l'aventurier")
			}

			if inventory.CheckMaxInventory(c) {
				fmt.Scan(&craft_confirmation)
				fmt.Println("1 - Oui je suis sûr !")
				fmt.Println("2 - Non je ne suis pas sûr, retour !")
				switch craft_confirmation {
				case 1:
					inventory.AddInventory(c, tuniqueAventurier)
					fmt.Println("Tu viens de fabriquer", itemcraftable, "x Tunique de l'Aventurier")
					inventory.RemoveInventory(c, wolfFur)
					inventory.RemoveInventory(c, trollSkin)
				case 2:
					fmt.Println("Retour au menu précédent")
					break
				}
			} else {
				fmt.Printf("\nTu n'as pas de places dans ton inventaire, il est plein...\n\n")
			}

			return

		case 3:
			itemcraftable := min(countWolfFur, countBoarLeather)
			if itemcraftable < 1 {
				fmt.Println("Vous n'avez pas les items requis pour craft les Bottes de l'Aventurier")
			} else {
				fmt.Println("Vous pouvez craft", itemcraftable, "x Bottes de l'Aventurier")
			}
			switch craft_confirmation {
			case 1:
				fmt.Println("1 - Oui je suis sûr !")
				// Enlevez les items de l'inventaires et craft le ou les items

			case 2:
				fmt.Println("1 - Non je ne suis pas sûr, retour !")

			}
		case 4:
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer")
		}
	}
}
