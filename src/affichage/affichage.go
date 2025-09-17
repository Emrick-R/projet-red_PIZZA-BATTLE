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
	fmt.Println("1 - ▶️  Commencer une nouvelle partie")
	fmt.Println("2 - 🪺  Options (Easter Egg)")
	fmt.Println("3 - 👋 Quitter")
	Separator()
}

// AffichageMenuPrincipal affiche le menu principal
func AffichageMenuPrincipal() {
	Separator()
	fmt.Println("🏠 Menu Principal :")
	Separator()
	fmt.Println("1 - 👤 Afficher le personnage")
	fmt.Println("2 - 🎒 Inventaire")
	fmt.Println("3 - ⚔️  Combat en 1 contre 1")
	fmt.Println("4 - 🛒 Marchand")
	fmt.Println("5 - ⚒️  Atelier de l'Oncle Pepito")
	fmt.Println("6 - ⬅️  RETOUR")
	Separator()
}

//AffichageMenuInventaire affiche le menu de l'inventaire
func AffichageMenuInventaire() {
	Separator()
	fmt.Println("🎒 Inventaire :")
	Separator()
	fmt.Println("1 - 🧪 Utiliser un consommable")
	fmt.Println("2 - 🛡️  Equiper un équipement")
	fmt.Println("3 - ⬅️  RETOUR")
	Separator()
}

//AffichageMenuCombatPotion affiche le menu des potions uniquement lors des combats
func AffichageMenuCombatPotion() {
	Separator()
	fmt.Println("🧪 Utiliser un consommable :")
	Separator()
	fmt.Println("1 - ❤️ Tiramisu")
	fmt.Println("2 - 💀 Tabasco")
	fmt.Println("3 - 🔵 Bocal de Sauce Tomate")
	fmt.Println("4 - ⬅️  RETOUR")
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

func EasterEgg() {
	// Effacer l'écran.
	fmt.Print("\033[H\033[2J")

	fmt.Println("===============================================")
	fmt.Println("        *** EASTER EGG DÉCOUVERT! ***         ")
	fmt.Println("===============================================")

	fmt.Println("\n🍕 Pizza Battle 🍕")
	fmt.Printf("Développé par les légendaires goateurs de pizza:\n\n")
	fmt.Println("Emrick Rivet & Harold François")

	fmt.Println("\n-----------------------------------------------")
	fmt.Println("RÉFÉRENCES CACHÉES DANS LE JEU:")
	fmt.Println("-----------------------------------------------")

	fmt.Println("\nÉtape 2 - Références musicales:")
	fmt.Println("ABBA")

	fmt.Println("\nÉtape 3 - Références cinématographiques:")
	fmt.Println("Steven Spielberg")

	fmt.Println("\n===============================================")
	fmt.Println("Appuyez sur 0 pour revenir au menu principal")
	fmt.Println("===============================================")
}
