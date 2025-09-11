package character

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/inventory"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxPv     int
	ActualPv  int
	Inventory []inventory.Object
}

func InitCharacter(name string, class string, level int, maxPv int, actualPv int, inv []inventory.Object) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxPv:     maxPv,
		ActualPv:  actualPv,
		Inventory: inv,
	}
}

func (c Character) DisplayInfo() {
	fmt.Println(c)
}

func (c Character) AccessInventory() {
	fmt.Println(c.Inventory)
}
