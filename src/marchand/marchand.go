package marchand

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func Marchand(c *structures.Character) {
	var marchand_choice int
	for {
		fmt.Println("======== Marchand : ========")
		fmt.Println("Bonjour je suis le marchand, quel est votre choix ?")
		fmt.Println("Choix 1 : Potion de vie (Gratuit)")
		fmt.Println("Choix 2 : Objet 2")
		fmt.Println("Choix 2 : Objet 3")
		fmt.Println("RETOUR")
		fmt.Scan(&marchand_choice)
		switch marchand_choice {
		case 1:
			c.AddInventory("Potion de vie")
			fmt.Println("Super ! Tu as obtenu 1 potion")
		}
	}
}
