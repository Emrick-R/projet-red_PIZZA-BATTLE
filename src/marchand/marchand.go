package marchand

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/affichage"
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
	HpPot := structures.Object{Name: "Tiramisu", Quantity: 1}
	ManaPot := structures.Object{Name: "Bocal de Sauce Tomate", Quantity: 1}
	PoisonPot := structures.Object{Name: "Tabasco", Quantity: 1}
	WolfFur := structures.Object{Name: "Laine", Quantity: 1}
	TrollSkin := structures.Object{Name: "Cuir", Quantity: 1}
	BoarLeather := structures.Object{Name: "Ficelle", Quantity: 1}
	RavenFeather := structures.Object{Name: "Coton", Quantity: 1}

	//Compétences
	FireBall := structures.Skill{Name: "Lancer de Piments", Damage: 20, ManaCost: 20}
	CurryExplosion := structures.Skill{Name: "Explosion de Curry", Damage: 60, ManaCost: 60}
	Ultimate := structures.Skill{Name: "Ma che vuoi ? 🤌 ", Damage: 500, ManaCost: 500}
	DeathSpell := structures.Skill{Name: "Covid-19", Damage: 500}

	// Effacer l'écran
	fmt.Print("\033[H\033[2J")

	//Boucle principale du marchand
	for {
		// Affichage de l'inventaire
		character.AccessInventory(c)
		//Menu du marchand
		affichage.Separator()
		fmt.Println("🛒 Bienvenue au SuperMarché !")
		affichage.Separator()
		fmt.Println("Que veux-tu acheter ?")
		fmt.Printf("1 - 🍰 %s (+50PV) - 3💵\n", HpPot.Name)
		fmt.Printf("2 - 🌶️  %s (10 dégâts sur 3 Tours) - 6💵\n", PoisonPot.Name)
		fmt.Printf("3 - 🫙  %s (+25 Sauces Tomate) - 4💵\n", ManaPot.Name)
		fmt.Printf("4 - 📖 Livre de Cuisine : %s - 20💵\n", FireBall.Name)
		fmt.Printf("5 - 📖 Livre de Cuisine : %s - 60💵\n", CurryExplosion.Name)
		fmt.Printf("6 - 📖 Livre de Cuisine : %s - 250💵\n", Ultimate.Name)
		fmt.Printf("7 - 🧶 %s - 4💵\n", WolfFur.Name)
		fmt.Printf("8 - 🟫 %s - 7💵\n", TrollSkin.Name)
		fmt.Printf("9 - 🪢  %s - 3💵\n", BoarLeather.Name)
		fmt.Printf("10 - ☁️  %s - 1💵\n", RavenFeather.Name)
		fmt.Printf("11 - 📦 Agrandir inventaire (+10 emplacements) - 30💵\n")
		fmt.Println("12 - ⬅️  RETOUR")
		// Test
		fmt.Printf("\n13 - 📖 Livre de Cuisine : %s - 0💵 (test)\n", DeathSpell.Name)
		// Fin test

		affichage.Separator()
		fmt.Print("👉 Ton choix : ")
		fmt.Scan(&marchand_choice)
		switch marchand_choice {
		case 1:
			// Achat d'un Tiramisu (potion de vie)
			// Vérification de la place dans l'inventaire
			if inventory.CheckMaxInventory(c) {
				// Ajout de l'objet dans l'inventaire
				inventory.AddInventory(c, HpPot)
				// Déduction de l'argent
				c.Money -= 3
				// Message de confirmation
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Println("✅ Tu as acheté un Tiramisu (-3💵)")
				// Affichage de l'argent restant
				fmt.Println("Tu as maintenant", c.Money, "💵")
				// Retour au menu du marchand
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				// Message d'erreur si l'inventaire est plein
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
				// Retour au menu du marchand
			}
		case 2:
			// Achat d'une potion de poison
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, PoisonPot)
				c.Money -= 6
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Println("✅ Tu as acheté un Tabasco (-6💵)")
				fmt.Println("Tu as maintenant", c.Money, "💵")

			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")

			}
		case 3:
			// Achat d'une potion de mana
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, ManaPot)
				c.Money -= 4
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Println("✅ Tu as acheté un Bocal de Sauce Tomate (-4💵)")
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
			}
		case 4:
			// Achat d'un livre de sort (compétence)
			// Vérification si le personnage possède déjà la compétence
			if !skills.CheckSkills(c, FireBall) {
				// Ajout de la compétence au personnage
				skills.AddSkills(c, FireBall)
				c.Money -= 25
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("✅ Tu as acheté un Livre de Cuisine : %s (-25💵)\n", FireBall.Name)
				// Affichage de la compétence apprise
				fmt.Printf("🔥 Nouvelle compétence : %s (%d dégâts)\n", FireBall.Name, FireBall.Damage)
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Print("\n❌ Tu possèdes déjà cette compétence\n\n")
			}
		case 5:
			if !skills.CheckSkills(c, CurryExplosion) {
				skills.AddSkills(c, CurryExplosion)
				c.Money -= 60
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("✅ Tu as acheté un Livre de Cuisine : %s (-60💵)\n", CurryExplosion.Name)
				fmt.Printf("💥 Nouvelle compétence : %s (%d dégâts)\n", CurryExplosion.Name, CurryExplosion.Damage)
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Print("\n❌ Tu possèdes déjà cette compétence\n\n")
			}
		case 6:
			if !skills.CheckSkills(c, Ultimate) {
				skills.AddSkills(c, Ultimate)
				c.Money -= 250
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("✅ Tu as acheté un Livre de Cuisine : %s (-250💵)\n", Ultimate.Name)
				fmt.Printf("🤌 Nouvelle compétence : %s (%d dégâts)\n", Ultimate.Name, Ultimate.Damage)
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Print("\n❌ Tu possèdes déjà cette compétence\n\n")
			}
		case 7:
			// Achat d'une Laine
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, WolfFur)
				c.Money -= 4
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Println("✅ Tu as acheté une Laine (-4💵)")
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
			}
		case 8:
			// Achat d'une Cuir
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, TrollSkin)
				c.Money -= 7
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Println("✅ Tu as acheté une Cuir (-7💵)")
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
			}
		case 9:
			// Achat d'un Ficelle
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, BoarLeather)
				c.Money -= 3
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Println("✅ Tu as acheté un Ficelle (-3💵)")
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
			}
		case 10:
			// Achat d'une plume de corbeau
			if inventory.CheckMaxInventory(c) {
				inventory.AddInventory(c, RavenFeather)
				c.Money -= 1
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Println("✅ Tu as acheté une Coton (-1💵)")
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("\n❌ Il n'y a pas assez de place dans l'inventaire\n\n")
			}

		case 11:
			// Achat d'une augmentation de l'inventaire
			c.Money -= 30
			// Effacer l'écran
			fmt.Print("\033[H\033[2J")
			fmt.Println("✅ Ton inventaire s'est agrandi de 10 places (-30💵)")
			fmt.Println("Tu as maintenant", c.Money, "💵")
			// Augmentation de la taille de l'inventaire de 10 slots
			inventory.UpgradeInventorySlot(c)

		case 12:
			// Retour au menu précédent
		case 13:
			// Test achat d'un livre de sort (compétence)
			if !skills.CheckSkills(c, DeathSpell) {
				skills.AddSkills(c, DeathSpell)
				// Pas de coût en argent (test)
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Printf("✅ Tu as acheté un Livre de Cuisine : %s\n", DeathSpell.Name)
				fmt.Printf("💀 Nouvelle compétence : %s (%d dégâts)\n", DeathSpell.Name, DeathSpell.Damage)
				fmt.Println("Tu as maintenant", c.Money, "💵")
			} else {
				// Effacer l'écran
				fmt.Print("\033[H\033[2J")
				fmt.Print("\n❌ Tu possèdes déjà cette compétence\n\n")
			}
		default:
			// Effacer l'écran
			fmt.Print("\033[H\033[2J")
			// Choix autre que 1 à 13
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}
		// Retour au menu précédent si le choix est 10
		// Reset de la variable marchand_choice si le choix est 10
		// Permet de ne pas rester bloqué dans la boucle du marchand
		if marchand_choice == 12 {
			marchand_choice = 0

			// Effacer l'écran
			fmt.Print("\033[H\033[2J")

			break
		}
	}
}
