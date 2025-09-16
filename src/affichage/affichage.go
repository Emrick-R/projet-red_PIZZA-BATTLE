package affichage

import "fmt"

func AffichageMenuPrincipal() {
	fmt.Println("======== Faites votre choix : ========")
	fmt.Println("1 - Commencer une nouvelle partie")
	fmt.Println("2 - Quitter")
}

func AffichageMenuPersonnage() {
	fmt.Println("\n======== Menu Personnage : ========")
	fmt.Println("1 - Afficher le personnage")
	fmt.Println("2 - Afficher l'inventaire")
	fmt.Println("3 - Test de combat : Utiliser une potion de poison")
	fmt.Println("4 - Marchand")
	fmt.Println("5 - Forgeron")
	fmt.Println("6 - RETOUR")

}

func AffichageMenuInventaire() {
	fmt.Println("======== Menu Inventaire : ========")
	fmt.Println("1 - Utiliser une potion")
	fmt.Println("2 - Equiper un équipement")
	fmt.Println("3 - RETOUR")
}
