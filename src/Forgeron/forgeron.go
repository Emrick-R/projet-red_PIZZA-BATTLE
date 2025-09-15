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
	for {
		fmt.Println("======== Forgeron : ========")
		fmt.Println("Bonjour je suis le Forgeron, quel est votre choix ?")
		fmt.Println("1 - Craft : Chapeau de l'aventurier (1 Plume de Corbeau + 1 Cuir de Sanglier)")
		fmt.Println("2 - Craft : Tunique de l'aventurier (2 Fourrures de Loup + 1 Peau de Troll)")
		fmt.Println("3 - Craft : Bottes de l'aventurier (1 Fourrure de Loup + 1 Cuir de Sanglier)")
		fmt.Println("4 - RETOUR")
		fmt.Scan(&forgeron_choice)

		var countRavenFeather, countBoarLeather, countWolfFur, countTrollSkin int

		for i := range c.Inventory {
			switch c.Inventory[i].Name {
			case "RavenFeather":
				countRavenFeather = c.Inventory[i].Quantity
			case "BoarLeather":
				countBoarLeather = c.Inventory[i].Quantity
			case "WolfFur":
				countWolfFur = c.Inventory[i].Quantity
			case "TrollSkin":
				countTrollSkin = c.Inventory[i].Quantity
			}
		}

		switch forgeron_choice {
		case 1:
			itemcraftable := min(countRavenFeather, countBoarLeather)
			if itemcraftable < 1 {
				fmt.Println("Vous n'avez pas les items requis pour craft le Chapeau de l'Aventurier")
			} else {
				fmt.Println("Vous pouvez craft", itemcraftable, "x Chapeau de l'Aventurier")
				fmt.Scan(&craft_confirmation)
			}
			switch craft_confirmation {
			case 1:
				fmt.Println("1 - Oui je suis sûr !")
				if inventory.CheckMaxInventory(c) {
					inventory.AddInventory(c, chapeauAventurier)
					fmt.Println("Vous venez de craft", itemcraftable, "x Chapeau de l'Aventurier")
					inventory.RemoveInventory(c, ravenFeather) // Enlever 1 plume de corbeau et 1 cuir de sanglier

				} else {
					fmt.Println("Vous n'avez pas de places dans votre inventaire, il est plein...")
				}

				return
			case 2:
				fmt.Println("1 - Non je ne suis pas sûr, retour !")

			}
		case 2:
			itemcraftable := min(countWolfFur/2, countTrollSkin)
			if itemcraftable < 1 {
				fmt.Println("Vous n'avez pas les items requis pour craft la Tunique de l'Aventurier")
			} else {
				fmt.Println("Vous pouvez craft", itemcraftable, "x Tunique de l'Aventurier")
			}
			switch craft_confirmation {
			case 1:
				fmt.Println("1 - Oui je suis sûr !")
				// Enlevez les items de l'inventaires et craft le ou les items

			case 2:
				fmt.Println("1 - Non je ne suis pas sûr, retour !")

			}
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
