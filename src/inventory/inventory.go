package inventory

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
)

type Object struct {
	Name     string
	Quantity int
}

func AddInventory(c *character.Character, obj Object) {
	/* Si un objet existe déjà dans l'inventaire incrémenter sa quantité sinon créer un nouvel objet et mettre sa quantité à 1 */
	for i, obj := range c.Inventory {
		if c.Inventory[i] == c.Inventory[i].Name {
			obj.Quantity = obj.Quantity + 1
		} else {
			c.Inventory = append(c.Inventory, obj)
		}
	}
}

func RemoveInventory(c *character.Character, obj Object) {
	/* Si un objet existe déjà dans l'inventaire vérifier si sa quantité est égal à 1 puis si c'est 1 l'enlever de l'inventaire sinon rien pcq pas présent dans l'inventaire*/
	for i, obj := range c.Inventory {
		if c.Inventory[i] == c.Inventory[i].Name {
			if obj.Quantity == 1 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			} else {
				fmt.Println("Vous ne pouvez pas perdre cet objet car il n'existe pas dans l'inventaire")
				return
			}
		}
	}
}
