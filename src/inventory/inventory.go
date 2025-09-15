package inventory

import (
	"projet-red_PIZZA-BATTLE/structures"
)

func AddInventory(c *structures.Character, newObj structures.Object) {
	for i := range c.Inventory {
		if c.Inventory[i].Name == newObj.Name {
			c.Inventory[i].Quantity += newObj.Quantity
			return
		}
	}
	c.Inventory = append(c.Inventory, newObj)
}

func RemoveInventory(c *structures.Character, obj structures.Object) {
	for i := range c.Inventory {
		if c.Inventory[i].Name == obj.Name {
			if c.Inventory[i].Quantity > obj.Quantity {
				c.Inventory[i].Quantity -= obj.Quantity
			} else {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			return
		}
	}
}

func CheckMaxInventory(c *structures.Character) bool { // True = Il y a de la place
	counter := 0
	for i := range c.Inventory {
		counter = counter + c.Inventory[i].Quantity
	}
	return counter < c.MaxInv
}

func UpgradeInventorySlot(c *structures.Character) {
	c.MaxInv = c.MaxInv + 10
}
