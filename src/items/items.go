package items

func TakePot(c1 *Character) {
	inv := c1.inventory
	for i := 0; i < len(c1.inventory); i++ {
		if Object.Name == "potion" {
			c1.ActualPv = c1.ActualPv + 50
			Object.Quantity[i] = Object.Quantity[i] - 1
			if Object.Quantity[i] == 0 {
				Object.Name[i] = ""
			}
		}

	}

}

/* Fonction access inventory
check si il y a une potion de vie dedans
et si oui il y'en a la consommer et rajouter 50 pv aux pdv actuels du perso  si non, dire il n'y a pas de potions dans l'inventaire*/
