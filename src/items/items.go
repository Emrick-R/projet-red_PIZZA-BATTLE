package items

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
	"time"
)

// Fonction pour utiliser une potion de vie /
func TakePot(c *structures.Character) {
	// Définir le nom de la potion de vie /
	HpPot := structures.Object{Name: "Tiramisu"}
	// Vérifier si le personnage a déjà les PV max
	if c.ActualHp == c.MaxHp {
		fmt.Printf("\n❌ Les points de vie sont déjà au max\n\n")

		time.Sleep(2 * time.Second)

		// Ne rien faire
		return
	}
	// Parcourir l'inventaire du personnage pour trouver la potion de vie/
	for i := 0; i < len(c.Inventory); i++ {
		// Si la potion de vie/ est trouvée
		if c.Inventory[i].Name == HpPot.Name {
			// Augmenter les PV actuels du personnage de 50
			c.ActualHp += 50
			// Réduire la quantité de la potion de vie dans l'inventaire
			c.Inventory[i].Quantity--
			// Si la quantité de la potion de vie est égale à 0, la retirer de l'inventaire
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			// Si les PV actuels dépassent les PV max, les ramener aux PV max
			if c.ActualHp >= c.MaxHp {
				c.ActualHp = c.MaxHp
			}
			// Afficher un message indiquant que la potion a été consommée et les PV actuels
			fmt.Printf("\n🍰 Tiramisu consommé ! +50 PV\n")
			fmt.Printf("❤️  %s : %d/%dPV\n\n", c.Name, c.ActualHp, c.MaxHp)

			time.Sleep(2 * time.Second)

			// Sortir de la fonction
			return
		}
	}
	// Si la potion de vie n'est pas trouvée dans l'inventaire, afficher un message d'erreur
	fmt.Println("❌ Il n'y a pas de Tiramisu dans l'inventaire")

	time.Sleep(2 * time.Second)

}

// Fonction pour utiliser une potion de mana
func TakeManaPot(c *structures.Character) {
	ManaPot := structures.Object{Name: "Bocal de Sauce Tomate"}
	if c.ActualMana == c.MaxMana {
		fmt.Println("❌ La Sauce Tomate est déjà pleine")

		time.Sleep(2 * time.Second)

		return
	}
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == ManaPot.Name {
			c.ActualMana += 25
			c.Inventory[i].Quantity--
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			if c.ActualMana >= c.MaxMana {
				c.ActualMana = c.MaxMana
			}
			fmt.Printf("\n🫙 Bocal consommée !\n")
			fmt.Printf("🍅 %s : %d/%d Sauce Tomate\n", c.Name, c.ActualMana, c.MaxMana)

			time.Sleep(2 * time.Second)

			return
		}
	}
	fmt.Println("❌Il n'y a pas de Bocal de Sauce Tomate dans l'inventaire")

	time.Sleep(2 * time.Second)

}

// Fonction pour utiliser une potion de poison.
// A FINIR (effet sur plusieurs tours)
func ThrowPoisonPot(c *structures.Character, e *structures.Enemy) {
	// Définir le nom de la potion de poison
	PoisonPot := structures.Object{Name: "Tabasco"}
	// Parcourir l'inventaire du personnage pour trouver la potion de poison, si trouvée l'utiliser, la retirer
	// et infliger des dégâts sur plusieurs tours à l'ennemi
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == PoisonPot.Name {
			c.Inventory[i].Quantity--
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			// Effet de la potion de poison sur 3 tours (10 dégâts par tour)
			//A FINIR (ajouter un effet de poison sur plusieurs tours)
			fmt.Printf("\nTabasco envoyée !\n")

			time.Sleep(2 * time.Second)

			// Appliquer l’effet poison
			e.Poisoned = true
			e.PoisonTurns = 3   // dure 3 tours
			e.PoisonDamage = 10 // 10 dégâts par tour

			fmt.Println("\n💀 L'ennemi est empoisonné pour 3 tours !")

			time.Sleep(2 * time.Second)

			return
		}
	}
	fmt.Println("❌ Il n'y a pas de Tabasco dans l'inventaire.")

	time.Sleep(2 * time.Second)

}

func CheckPoisonStatus(e *structures.Enemy) {
	if e.Poisoned && e.PoisonTurns > 0 {
		e.ActualHp -= e.PoisonDamage
		e.PoisonTurns--
		fmt.Printf("☠️  %s subit %d dégâts de poison (%d tours restants)\n\n", e.Name, e.PoisonDamage, e.PoisonTurns)
		fmt.Printf("❤️  %s : %d/%d PV\n\n", e.Name, e.ActualHp, e.MaxHp)

		time.Sleep(2 * time.Second)

		if e.PoisonTurns == 0 {
			e.Poisoned = false
			fmt.Printf("✅ Le poison s'est dissipé.\n\n")

			time.Sleep(2 * time.Second)

		}
	}

}
