package character

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/inventory"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHp     int
	ActualHp  int
	Inventory []inventory.Object
}

func InitCharacter(name string, class string, level int, maxhp int, actualhp int, inv []inventory.Object) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHp:     maxhp,
		ActualHp:  actualhp,
		Inventory: inv,
	}
}

func (c *Character) DisplayInfo() {
	fmt.Printf("\nNom: %s\nClasse: %s\nNiveau: %d\nPoint de vie: %d/%d\n\n", c.Name, c.Class, c.Level, c.ActualHp, c.MaxHp)
}

func (c Character) AccessInventory() {
	fmt.Println("\nInventaire:")
	for _, item := range c.Inventory {
		fmt.Printf("\n %s %d\n", item.Name, item.Quantity)
	}
	fmt.Print("\n")
}
