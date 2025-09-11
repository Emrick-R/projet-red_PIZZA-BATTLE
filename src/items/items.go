package items

func TakePot(c *Character) {
	inv := c.inventory
	for i := 0; i < len(c.inventory); i++ {
		if Object.Name == "potion" {
			c.ActualPv = c.ActualPv + 50
			Object.Quantity[i] = Object.Quantity[i] - 1
			if Object.Quantity[i] == 0 {
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
				return
			}
		}
	}

}

/* Fonction access inventory
check si il y a une potion de vie dedans
et si oui il y'en a la consommer et rajouter 50 pv aux pdv actuels du perso  si non, dire il n'y a pas de potions dans l'inventaire*/
