package character

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func DisplayCInfo(c *structures.Character) {
	fmt.Printf("\nNom: %s\nClasse: %s\nNiveau: %d\nPV: %d/%d\nArgent: %d\n",
		c.Name, c.Class, c.Level, c.ActualHp, c.MaxHp, c.Money)
	for _, skill := range c.SkillList {
		fmt.Printf("Liste des skills :\n%s, %d de dégats\n\n", skill.Name, skill.Damage)
	}

}

func DisplayEInfo(e *structures.Enemy) {
	fmt.Printf("\nNom: %s\nPV: %d/%d\n",
		e.Name, e.ActualHp, e.MaxHp)
}

func AccessInventory(c *structures.Character) {
	fmt.Println("\nInventaire :")
	for _, item := range c.Inventory {
		fmt.Printf("%s x%d\n", item.Name, item.Quantity)
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
