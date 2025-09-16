package affichage

import "fmt"

func separator() {
	fmt.Println("======================================")
}

func AffichageMenuPrincipal() {
	separator()
	fmt.Println("🍕 BIENVENUE DANS PIZZA BATTLE 🍕")
	separator()
	fmt.Println("1 - Commencer une nouvelle partie")
	fmt.Println("2 - Quitter")
	separator()
}

func AffichageMenuPersonnage() {
	separator()
	fmt.Println("👤 Menu Personnage :")
	separator()
	fmt.Println("1 - Afficher le personnage")
	fmt.Println("2 - Afficher l'inventaire")
	fmt.Println("3 - Test de combat : Utiliser une potion de poison")
	fmt.Println("4 - Marchand")
	fmt.Println("5 - Forgeron")
	fmt.Println("6 - RETOUR")
	separator()
}

func AffichageMenuInventaire() {
	separator()
	fmt.Println("🎒 Inventaire :")
	separator()
	fmt.Println("1 - Utiliser une potion")
	fmt.Println("2 - Equiper un équipement")
	fmt.Println("3 - RETOUR")
	separator()
}

func CharacterTurn() {
	separator()
	fmt.Println("⚔️  Combat :")
	separator()
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Inventaire")
	fmt.Println("3 - RETOUR")
	separator()
}

// Pour afficher les attaques, passe la liste des skills en paramètre
func Attaques(skills []string) {
	separator()
	fmt.Println("💥 Attaques :")
	for i, skill := range skills {
		fmt.Printf("%d - %s\n", i+1, skill)
	}
	fmt.Println("0 - RETOUR")
	separator()
}
