package items

import (
	"projet-red_PIZZA-BATTLE/character"
)

func TakePot(c1 character.Character) {
	for i := 0; i < len(c1.Inventory); i++ {
		if c1.Inventory[i].Name == "Potion" {
			c1.ActualPv = c1.ActualPv + 50
			c1.Inventory[i].Quantity = c1.Inventory[1].Quantity - 1
			if c1.Inventory[1].Quantity == 0 {
				c1.Inventory = append(c1.Inventory[:i], c1.Inventory[i+1:]...)
				return
			}
		}
	}

}

/* Fonction access inventory
check si il y a une potion de vie dedans
et si oui il y'en a la consommer et rajouter 50 pv aux pdv actuels du perso  si non, dire il n'y a pas de potions dans l'inventaire*/
