package marchand

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/structures"
)

func Marchand(c *structures.Character) {
	var marchand_choice int
	PotionDeVie := structures.Object{"Potion de Vie", 1}
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
			inventory.AddInventory(c, PotionDeVie)
			fmt.Println("Super ! Tu as obtenu 1 potion")
			character.AccessInventory(c)
		}
	}
}
