package items

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func TakePot(c *structures.Character) {
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == "Potion de vie" {
			c.ActualHp += 50
			c.Inventory[i].Quantity--
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			fmt.Printf("Potion consommée !\n")
			fmt.Printf("PV actuels: %d\n", c.ActualHp)

			return
		}
	}
	fmt.Println("Il n'y a pas de potions dans l'inventaire")
}
