package forgeron

import (
	"fmt"
	structures "projet-red_PIZZA-BATTLE/Structures"
)

func Forgeron(c *structures.Character) {
	var forgeron_choice int
	for {
		fmt.Println("======== Forgeron : ========")
		fmt.Println("Bonjour je suis le Forgeron, quel est votre choix ?")
		fmt.Println("1 - Craft : Chapeau de l’aventurier (Nécessite 1 Plume de Corbeau et 1 Cuir de Sanglier)")
		fmt.Println("2 - Craft : Tunique de l’aventurier (Nécessite 2 Fourrures de Loup et 1 Peau de Troll)")
		fmt.Println("3 - Craft : Bottes de l’aventurier (Nécessite 1 Fourrure de Loup et 1 Cuir de Sanglier)")
		fmt.Println("4 - RETOUR")
		fmt.Scan(&forgeron_choice)
		switch forgeron_choice {
		// Rajouter les équipements dans structures puis mettre les conditions et la logique pour les craft
		case 1:
			fmt.Println("Super ! Tu as craft Chapeau de l'aventurier.")
			fmt.Println("Tu as maintenant")

		case 2:
			fmt.Println("Super ! Tu as craft Tunique de l’aventurier")
			fmt.Println("Tu as maintenant")
		case 3:
			fmt.Println("Super ! Tu as craft Bottes de l’aventurier")
			fmt.Println("Tu as maintenant")
		case 4:
			return
		}
	}
}
