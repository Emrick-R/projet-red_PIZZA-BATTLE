package inventory

import "projet-red_PIZZA-BATTLE/structures"

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
				// retirer complètement si quantité <= obj.Quantity
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			return
		}
	}
}
