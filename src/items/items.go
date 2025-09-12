package items

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func TakePot(c *structures.Character) {
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i].Name == "Potion" {
			c.ActualHp += 50
			c.Inventory[i].Quantity--
			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			fmt.Println("Potion consommée !")
			return
		}
	}
	fmt.Println("Il n'y a pas de potions dans l'inventaire")
}
