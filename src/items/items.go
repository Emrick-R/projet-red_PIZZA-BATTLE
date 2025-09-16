package items

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
	"time"
)

// Fonction pour utiliser une potion de vie
func TakePot(c *structures.Character) {
	// Définir le nom de la potion de vie
	HpPot := structures.Object{Name: "Potion de Vie"}
	// Vérifier si le personnage a déjà les PV max
	if c.ActualHp == c.MaxHp {
		fmt.Printf("\nLes points de vie sont déjà au max\n\n")
		// Ne rien faire
		return
	}
	// Parcourir l'inventaire du personnage pour trouver la potion de vie
	for i := 0; i < len(c.Inventory); i++ {
		// Si la potion de vie est trouvée
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
			fmt.Printf("\nPotion consommée ! +50 PV\n")
			fmt.Printf("PV actuels: %d\n", c.ActualHp)
			// Sortir de la fonction
			return
		}
	}
	// Si la potion de vie n'est pas trouvée dans l'inventaire, afficher un message d'erreur
	fmt.Println("Il n'y a pas de potions de Vie dans l'inventaire")
}

// Fonction pour utiliser une potion de mana
// A FINIR
func TakeManaPot(c *structures.Character) {
	ManaPot := structures.Object{Name: "Potion de Mana"}
	if c.ActualMana == c.ManaMax {
		fmt.Printf("\nLes points de vie sont déjà au max\n\n")
		return
	}
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == ManaPot.Name {
			c.ActualHp += 20
			c.Inventory[i].Quantity--
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			if c.ActualMana >= c.ManaMax {
				c.ActualMana = c.ManaMax
			}
			fmt.Printf("\nPotion consommée !\n")
			fmt.Printf("Mana actuel: %d\n", c.ActualMana)

			return
		}
	}
	fmt.Println("Il n'y a pas de potions de Mana dans l'inventaire")
}

// Fonction pour utiliser une potion de poison
// A FINIR (effet sur plusieurs tours)
func ThrowPoisonPot(c *structures.Character, e *structures.Enemy) {
	// Définir le nom de la potion de poison
	PoisonPot := structures.Object{Name: "Potion de Poison"}
	// Parcourir l'inventaire du personnage pour trouver la potion de poison, si trouvée l'utiliser, la retirer
	// et infliger des dégâts sur plusieurs tours à l'ennemi
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == PoisonPot.Name {
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				e.ActualHp -= 10
				c.Inventory[i].Quantity--
			}
			// Effet de la potion de poison sur 3 tours (10 dégâts par tour)
			//A FINIR (ajouter un effet de poison sur plusieurs tours)
			fmt.Printf("\nPotion envoyée !\n")
			for i := 0; i < 3; i++ {
				e.ActualHp -= 10
				fmt.Printf("L'ennemi a perdu : %d hp\n", e.ActualHp)
				time.Sleep(1 * time.Second)
			}

		}
	}
	fmt.Println("Il n'y a pas de potions de Poison dans l'inventaire")
}

// Fonction pour ajouter de l'argent au personnage en fonction de la difficulté de l'ennemi vaincu
func AddMoney(c *structures.Character, e *structures.Enemy) int {
	switch e.Difficulty {
	case "Facile":
		c.Money += 5
		return 5
	case "Normal":
		c.Money += 10
		return 10
	case "Boss":
		c.Money += 20
		return 20
	}
	return 0
}
