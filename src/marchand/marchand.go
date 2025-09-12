package marchand

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/structures"
)

func Marchand(c *structures.Character) {
	var marchand_choice int
	HpPot := structures.Object{"Potion de Vie", 1}
	for {
		fmt.Println("======== Marchand : ========")
		fmt.Println("Bonjour je suis le marchand, quel est votre choix ?")
		fmt.Println("1 - Potion de vie (Gratuit)")
		fmt.Println("2 - /vide/")
		fmt.Println("3 - /vide/")
		fmt.Println("4 - RETOUR")
		fmt.Scan(&marchand_choice)
		switch marchand_choice {
		case 1:
			inventory.AddInventory(c, HpPot)
			fmt.Println("Super ! Tu as obtenu 1 potion")
			character.AccessInventory(c)
		case 2:
			fmt.Println("Désolé ! Pas d'objet pour le moment.")
		case 3:
			fmt.Println("Désolé ! Pas d'objet pour le moment.")
		case 4:
			return

		}
	}
}
