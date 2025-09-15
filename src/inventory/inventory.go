package inventory

import (
	"fmt"
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

func AddEquipment(c *structures.Character, newObj structures.Object) {
	switch newObj.Name {
	case "Chapeau de l'aventurier":
		if c.Armor.Head.Name != "Chapeau de l'aventurier" {
			c.Armor.Head = &newObj
			fmt.Printf("\nTu équipes : Chapeau de l'aventurier\n\n")
			fmt.Printf("\nTu gagnes + 10 PV Max\n\n")
			fmt.Printf("%d/%d PV -> %d/%d PV\n\n", c.ActualHp, c.MaxHp, c.ActualHp, c.MaxHp+10)
			c.MaxHp += 10
			RemoveInventory(c, newObj)
		} else {
			fmt.Printf("\nTu as déjà équipé : %s\n\n", newObj.Name)
		}
	case "Tunique de l'aventurier":
		if c.Armor.Chest.Name != "Tunique de l'aventurier" {
			c.Armor.Chest = &newObj
			fmt.Printf("\nTu équipes : Tunique de l'aventurier\n\n")
			fmt.Printf("\nTu gagnes + 25 PV Max\n\n")
			fmt.Printf("%d/%d PV -> %d/%d PV\n\n", c.ActualHp, c.MaxHp, c.ActualHp, c.MaxHp+25)
			c.MaxHp += 25
			RemoveInventory(c, newObj)
		} else {
			fmt.Printf("\nTu as déjà équipé : %s\n\n", newObj.Name)
		}
	case "Bottes de l'aventurier":
		if c.Armor.Legs.Name != "Bottes de l'aventurier" {
			c.Armor.Legs = &newObj
			fmt.Printf("\nTu équipes : Bottes de l'aventurier\n\n")
			fmt.Printf("\nTu gagnes + 15 PV Max\n\n")
			fmt.Printf("%d/%d PV -> %d/%d PV\n\n", c.ActualHp, c.MaxHp, c.ActualHp, c.MaxHp+15)
			c.MaxHp += 15
			RemoveInventory(c, newObj)
		} else {
			fmt.Printf("\nTu as déjà équipé : %s\n\n", newObj.Name)
		}
	}
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
