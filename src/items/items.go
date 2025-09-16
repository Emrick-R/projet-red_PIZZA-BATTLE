package items

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
	"time"
)

func TakePot(c *structures.Character) {
	HpPot := structures.Object{Name: "Potion de Vie"}
	if c.ActualHp == c.MaxHp {
		fmt.Printf("\nLes points de vie sont déjà au max\n\n")
		return
	}
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == HpPot.Name {
			c.ActualHp += 50
			c.Inventory[i].Quantity--
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			if c.ActualHp >= c.MaxHp {
				c.ActualHp = c.MaxHp
			}
			fmt.Printf("\nPotion consommée ! +50 PV\n")
			fmt.Printf("PV actuels: %d\n", c.ActualHp)

			return
		}
	}
	fmt.Println("Il n'y a pas de potions de Vie dans l'inventaire")
}
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

func ThrowPoisonPot(c *structures.Character, e *structures.Enemy) {
	PoisonPot := structures.Object{Name: "Potion de Poison"}
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == PoisonPot.Name {
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
				e.ActualHp -= 10
				c.Inventory[i].Quantity--
			}
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
