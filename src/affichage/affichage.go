package affichage

import "fmt"

func separator() {
	fmt.Println("======================================")
}

func AffichageMenuDemarrage() {
	separator()
	fmt.Println("🍕 BIENVENUE DANS PIZZA BATTLE 🍕")
	separator()
	fmt.Println("1 - Commencer une nouvelle partie")
	fmt.Println("2 - Quitter")
	separator()
}

func AffichageMenuPrincipal() {
	separator()
	fmt.Println("👤 Menu Principal :")
	separator()
	fmt.Println("1 - Afficher le personnage")
	fmt.Println("2 - Afficher l'inventaire")
	fmt.Println("3 - Test de combat (1v1)")
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

func AffichageMenuCombatPotion() {
	separator()
	fmt.Println("🧪 Utiliser une potion :")
	separator()
	fmt.Println("1 - Potion de soin")
	fmt.Println("2 - Potion de poison")
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
