package character

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func DisplayInfo(c *structures.Character) {
	fmt.Printf("\nNom: %s\nClasse: %s\nNiveau: %d\nPV: %d/%d\n\n",
		c.Name, c.Class, c.Level, c.ActualHp, c.MaxHp)
}

func AccessInventory(c *structures.Character) {
	fmt.Println("\nInventaire :")
	for _, item := range c.Inventory {
		fmt.Printf(" %s x%d\n", item.Name, item.Quantity)
	}
	fmt.Println()
}

func IsDead(c *structures.Character) {
	if c.ActualHp <= 0 {
		fmt.Println("Vous êtes mort !")
		c.MaxHp /= 2
		c.ActualHp = c.MaxHp
		fmt.Println("Vous venez de renaître avec 50% de HP en moins.")
	}
}
