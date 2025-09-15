package marchand

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/structures"
)

func Marchand(c *structures.Character) {
	var marchand_choice int
	HpPot := structures.Object{Name: "Potion de Vie", Quantity: 1}
	PoisonPot := structures.Object{Name: "Potion de Poison", Quantity: 1}
	FireballBook := structures.Object{Name: "Livre de Sort : Boule de feu", Quantity: 1} //J'ai mit Objet mais je pense que ca appartient aux skills au skill list faudra revoir ca
	WolfFur := structures.Object{Name: "Peau de Loup", Quantity: 1}
	TrollSkin := structures.Object{Name: "Peau de Troll", Quantity: 1}
	BoarLeather := structures.Object{Name: "Cuir de Sanglier", Quantity: 1}
	RavenFeather := structures.Object{Name: "Plume de Corbeau", Quantity: 1}

	for {
		fmt.Println("\n======== Marchand : ========")
		fmt.Printf("Bonjour je suis le marchand, quel est votre choix ?\n\n")
		fmt.Println("1 - Potion de vie - 3 Pièces d'or")
		fmt.Println("2 - Potion de poison - 6 Pièces d'or")
		fmt.Println("3 - Livre de Sort : Boule de feu - 25 pièces d'or")
		fmt.Println("4 - Fourrure de Loup  - 4 pièces d'or")
		fmt.Println("5 - Peau de Troll - 7 pièces d'or")
		fmt.Println("6 - Cuir de Sanglier - 3 pièces d'or")
		fmt.Println("7 - Plume de Corbeau - 1 pièce d'or")
		fmt.Println("8 - RETOUR")
		fmt.Scan(&marchand_choice)
		switch marchand_choice {
		case 1:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, HpPot)
				c.Money -= 3
				fmt.Println("Super ! Tu as acheté une Potion de vie. Tu perds 3 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			}
			fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
			character.AccessInventory(c)
		case 2:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, PoisonPot)
				c.Money -= 6
				fmt.Println("Super ! Tu as acheté une Potion de poison. Tu perds 6 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			}
			fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
			character.AccessInventory(c)
		case 3:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, FireballBook)
				c.Money -= 25
				fmt.Println("Super ! Tu as acheté un Livre de Sort : Boule de feu. Tu perds 25 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			}
			fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
			character.AccessInventory(c)
		case 4:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, WolfFur)
				c.Money -= 4
				fmt.Println("Super ! Tu as acheté une Fourrure de Loup. Tu perds 4 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			}
			fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
			character.AccessInventory(c)
		case 5:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, TrollSkin)
				c.Money -= 7
				fmt.Println("Super ! Tu as acheté une Peau de Troll. Tu perds 7 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			}
			fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
			character.AccessInventory(c)
		case 6:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, BoarLeather)
				c.Money -= 3
				fmt.Println("Super ! Tu as acheté un Cuir de Sanglier. Tu perds 3 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			}
			fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
			character.AccessInventory(c)
		case 7:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, RavenFeather)
				c.Money -= 1
				fmt.Println("Super ! Tu as acheté une Plume de Corbeau. Tu perds 1 Pièce d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			}
			fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
			character.AccessInventory(c)
		case 8:
			return

		}
	}
}
