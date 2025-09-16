package marchand

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/skills"
	"projet-red_PIZZA-BATTLE/structures"
)

// Marchand permet au joueur d'acheter des objets et compétences s'il a assez d'argent
func Marchand(c *structures.Character) {
	//Initialisation des objets et compétences vendus par le marchand
	var marchand_choice int
	//Objets
	HpPot := structures.Object{Name: "Potion de Vie", Quantity: 1}
	ManaPot := structures.Object{Name: "Potion de Vie", Quantity: 1}
	PoisonPot := structures.Object{Name: "Potion de Poison", Quantity: 1}
	WolfFur := structures.Object{Name: "Peau de Loup", Quantity: 1}
	TrollSkin := structures.Object{Name: "Peau de Troll", Quantity: 1}
	BoarLeather := structures.Object{Name: "Cuir de Sanglier", Quantity: 1}
	RavenFeather := structures.Object{Name: "Plume de Corbeau", Quantity: 1}
	//Compétences
	FireBall := structures.Skill{Name: "Boule de feu", Damage: 20}
	DeathSpell := structures.Skill{Name: "Sort de la mort qui tue", Damage: 500}

	//Boucle principale du marchand
	for {
		//Menu du marchand
		fmt.Println("\n======== Marchand : ========")
		fmt.Printf("Bonjour je suis le marchand, quel est votre choix ?\n\n")
		fmt.Printf("1 - %s - 3 Pièces d'or\n", HpPot.Name)
		fmt.Printf("2 - %s - 6 Pièces d'or\n", PoisonPot.Name)
		fmt.Printf("3 - %s - 4 Pièces d'or\n", ManaPot.Name)
		fmt.Printf("4 - Livre de Sort : %s - 25 pièces d'or\n", FireBall.Name)
		fmt.Printf("5 - %s - 4 pièces d'or\n", WolfFur.Name)
		fmt.Printf("6 - %s - 7 pièces d'or\n", TrollSkin.Name)
		fmt.Printf("7 - %s - 3 pièces d'or\n", BoarLeather.Name)
		fmt.Printf("8 - %s - 1 pièce d'or\n", RavenFeather.Name)
		fmt.Printf("9 - Augmenter la taille de l'inventaire +10 slots - 30 pièce d'or\n")
		fmt.Println("10 - RETOUR")
		// Test
		fmt.Printf("11 - Livre de Sort : %s - 50 pièces d'or (test)\n", DeathSpell.Name)
		// Fin test
		fmt.Scan(&marchand_choice)
		switch marchand_choice {
		case 1:
			// Achat d'une potion de vie
			// Vérification de la place dans l'inventaire
			if inventory.CheckMaxInventory(c) {
				// Ajout de l'objet dans l'inventaire
				inventory.AddInventory(c, HpPot)
				// Déduction de l'argent
				c.Money -= 3
				// Message de confirmation
				fmt.Println("Super ! Tu as acheté une Potion de vie. Tu perds 3 Pièces d'or.")
				// Affichage de l'argent restant
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				// Affichage de l'inventaire
				character.AccessInventory(c)
				// Retour au menu du marchand
			} else {
				// Message d'erreur si l'inventaire est plein
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				// Affichage de l'inventaire
				character.AccessInventory(c)
				// Retour au menu du marchand
			}
		case 2:
			// Achat d'une potion de poison
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, PoisonPot)
				c.Money -= 6
				fmt.Println("Super ! Tu as acheté une Potion de poison. Tu perds 6 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 3:
			// Achat d'une potion de mana
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, ManaPot)
				c.Money -= 4
				fmt.Println("Super ! Tu as acheté une Potion de Mana. Tu perds 4 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 4:
			// Achat d'un livre de sort (compétence)
			// Vérification si le personnage possède déjà la compétence
			if !skills.CheckSkills(c, FireBall) {
				// Ajout de la compétence au personnage
				skills.AddSkills(c, FireBall)
				c.Money -= 25
				fmt.Printf("\nSuper ! Tu as acheté un Livre de Sort : %s. Tu perds 25 Pièces d'or.\n", FireBall.Name)
				// Affichage de la compétence apprise
				fmt.Printf("Tu connais maintenant la compétance %s : %d de dégats\n", FireBall.Name, FireBall.Damage)
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Print("\n❌ Tu possèdes déjà cette compétence\n\n")
				character.AccessInventory(c)
			}
		case 5:
			// Achat d'une fourrure de loup
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, WolfFur)
				c.Money -= 4
				fmt.Println("\nSuper ! Tu as acheté une Fourrure de Loup. Tu perds 4 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 6:
			// Achat d'une peau de troll
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, TrollSkin)
				c.Money -= 7
				fmt.Println("\nSuper ! Tu as acheté une Peau de Troll. Tu perds 7 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 7:
			// Achat d'un cuir de sanglier
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, BoarLeather)
				c.Money -= 3
				fmt.Println("\nSuper ! Tu as acheté un Cuir de Sanglier. Tu perds 3 Pièces d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}
		case 8:
			// Achat d'une plume de corbeau
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, RavenFeather)
				c.Money -= 1
				fmt.Println("\nSuper ! Tu as acheté une Plume de Corbeau. Tu perds 1 Pièce d'or.")
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				character.AccessInventory(c)
			}

		case 9:
			// Achat d'une augmentation de l'inventaire
			c.Money -= 30
			fmt.Println("\nSuper ! Ton inventaire s'est agrandi de 10 places. Tu perds 30 Pièce d'or.")
			fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
			// Augmentation de la taille de l'inventaire de 10 slots
			inventory.UpgradeInventorySlot(c)

		case 10:
			// Retour au menu précédent
		case 11:
			// Test achat d'un livre de sort (compétence)
			if !skills.CheckSkills(c, DeathSpell) {
				skills.AddSkills(c, DeathSpell)
				fmt.Printf("\nSuper ! Tu as acheté un Livre de Sort : %s. Tu test le sort.\n", DeathSpell.Name)
				fmt.Printf("Tu connais maintenant la compétance %s : %d de dégats\n", DeathSpell.Name, DeathSpell.Damage)
				fmt.Println("Tu as maintenant", c.Money, "Pièces d'or")
				character.AccessInventory(c)
			} else {
				fmt.Print("\n❌ Tu possèdes déjà cette compétence\n\n")
				character.AccessInventory(c)
			}
		default:
			// Choix autre que 1 à 10
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}
		// Retour au menu précédent si le choix est 10
		// Reset de la variable marchand_choice si le choix est 10
		// Permet de ne pas rester bloqué dans la boucle du marchand
		if marchand_choice == 10 {
			marchand_choice = 0
			break
		}
	}
}
