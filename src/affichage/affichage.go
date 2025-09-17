package affichage

import "fmt"

// Separator affiche une série de =
func Separator() {
	fmt.Println("==================================================")
}

// AffichageMenuDemarrage affiche le menu de démarrage
func AffichageMenuDemarrage() {
	Separator()
	fmt.Println("🍕 BIENVENUE DANS PIZZA BATTLE 🍕")
	Separator()
	fmt.Println("1 - Commencer une nouvelle partie")
	fmt.Println("2 - Quitter")
	Separator()
}

// AffichageMenuPrincipal affiche le menu principal
func AffichageMenuPrincipal() {
	Separator()
	fmt.Println("👤 Menu Principal :")
	Separator()
	fmt.Println("1 - Afficher le personnage")
	fmt.Println("2 - 🎒 Inventaire")
	fmt.Println("3 - Combat en 1 contre 1")
	fmt.Println("4 - 🛒 Marchand")
	fmt.Println("5 - Forgeron")
	fmt.Println("6 - ⬅️ RETOUR")
	Separator()
}

//AffichageMenuInventaire affiche le menu de l'inventaire
func AffichageMenuInventaire() {
	Separator()
	fmt.Println("🎒 Inventaire :")
	Separator()
	fmt.Println("1 - 🧪 Utiliser une potion")
	fmt.Println("2 - 🛡️ Equiper un équipement")
	fmt.Println("3 - ⬅️ RETOUR")
	Separator()
}

//AffichageMenuCombatPotion affiche le menu des potions uniquement lors des combats
func AffichageMenuCombatPotion() {
	Separator()
	fmt.Println("🧪 Utiliser une potion :")
	Separator()
	fmt.Println("1 - Potion de soin")
	fmt.Println("2 - Potion de poison")
	fmt.Println("3 - ⬅️ RETOUR")
	Separator()
}

//CharacterTurn affiche le menu lors du tour du joueur dans un combat
func CharacterTurn() {
	Separator()
	fmt.Println("⚔️  Combat :")
	Separator()
	fmt.Println("1 - 👊 Attaquer")
	fmt.Println("2 - 🎒 Inventaire")
	fmt.Println("3 - ⬅️ RETOUR")
	Separator()
}
