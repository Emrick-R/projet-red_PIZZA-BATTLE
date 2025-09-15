package marchand

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/skills"
	"projet-red_PIZZA-BATTLE/structures"
)

func Marchand(c *structures.Character) {
	var marchand_choice int
	HpPot := structures.Object{Name: "Potion de Vie", Quantity: 1}
	PoisonPot := structures.Object{Name: "Potion de Poison", Quantity: 1}
	WolfFur := structures.Object{Name: "Peau de Loup", Quantity: 1}
	TrollSkin := structures.Object{Name: "Peau de Troll", Quantity: 1}
	BoarLeather := structures.Object{Name: "Cuir de Sanglier", Quantity: 1}
	RavenFeather := structures.Object{Name: "Plume de Corbeau", Quantity: 1}

	FireBall := structures.Skill{Name: "Boule de feu", Damage: 20}

	for {
		fmt.Println("======== Marchand : ========")
		fmt.Printf("Bonjour je suis le marchand, quel est votre choix ?\n\n")
		fmt.Printf("1 - %s - 3 Pièces d'or\n", HpPot.Name)
		fmt.Printf("2 - %s - 6 Pièces d'or\n", PoisonPot.Name)
		fmt.Printf("3 - Livre de Sort : %s - 25 pièces d'or\n", FireBall.Name)
		fmt.Printf("4 - %s - 4 pièces d'or\n", WolfFur.Name)
		fmt.Printf("5 - %s - 7 pièces d'or\n", TrollSkin.Name)
		fmt.Printf("6 - %s - 3 pièces d'or\n", BoarLeather.Name)
		fmt.Printf("7 - %s - 1 pièce d'or\n", RavenFeather.Name)
		fmt.Printf("8 - Augmenter la taille de l'inventaire +10 slots - 30 pièce d'or\n")
		fmt.Println("9 - RETOUR")
		fmt.Scan(&marchand_choice)
		switch marchand_choice {
		case 1:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, HpPot)
				c.Money -= 3
				fmt.Println("Super ! Tu as acheté une Potion de vie. Tu perds 3 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 2:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, PoisonPot)
				c.Money -= 6
				fmt.Println("Super ! Tu as acheté une Potion de poison. Tu perds 6 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 3:
			if !skills.CheckSkills(c, FireBall) {
				skills.AddSkills(c, FireBall)
				c.Money -= 25
				fmt.Printf("Super ! Tu as acheté un Livre de Sort : %s. Tu perds 25 Pièces d'or.\n", FireBall.Name)
				fmt.Printf("Tu connais maintenant la compétance %s : %d de dégats\n", FireBall.Name, FireBall.Damage)
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Print("\nTu possèdes déjà cette compétence\n\n")
				character.AccessInventory(c)
			}
		case 4:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, WolfFur)
				c.Money -= 4
				fmt.Println("Super ! Tu as acheté une Fourrure de Loup. Tu perds 4 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 5:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, TrollSkin)
				c.Money -= 7
				fmt.Println("Super ! Tu as acheté une Peau de Troll. Tu perds 7 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 6:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, BoarLeather)
				c.Money -= 3
				fmt.Println("Super ! Tu as acheté un Cuir de Sanglier. Tu perds 3 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 7:
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, RavenFeather)
				c.Money -= 1
				fmt.Println("Super ! Tu as acheté une Plume de Corbeau. Tu perds 1 Pièce d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\nIl n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 8:
			return

		}
	}
}
