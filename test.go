package main

import (
	"fmt"
	"math/rand"
)

var choix int

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func rollDice() int {
	return rand.Intn(100) + 1
}

func mainn() {

	fmt.Println("🟩⬜🟥 Épreuve de la Mamma : choisissez un nombre, celui le plus proche du score de la Mamma commence.")

	// input joueur sécurisé
	for {
		fmt.Print("Entrez votre nombre (1-100) : ")
		_, err := fmt.Scan(&choix)
		if err == nil && choix >= 1 && choix <= 100 {
			break
		}
		fmt.Println("Valeur invalide ! Tapez un nombre entre 1 et 100.")
	}

	// premier lancer
	mamma := rollDice()
	ennemi := rollDice()

	// affichage clair
	fmt.Printf("Ton Chiffre : %d | Chiffre de la Mamma : %d | Chiffre de l'ennemi : %d\n", choix, mamma, ennemi)

	// en cas d'égalité
	for choix == ennemi {
		fmt.Println("Égalité — relance des dés !")
		mamma = rollDice()
		ennemi = rollDice()
		fmt.Printf("Chiffre : %d | Chiffre de la Mamma : %d | Chiffre de l'ennemi : %d\n", choix, mamma, ennemi)
	}

	// distances absolues
	distJoueur := abs(choix - mamma)
	distEnnemi := abs(ennemi - mamma)

	// initiative
	if distJoueur < distEnnemi {
		fmt.Printf("✅ Vous êtes le plus proche du chiffre de la Mamma avec une distance de %d, vous commencez !\n", distJoueur)
		initiative = initiative + distJoueur
		fmt.Printf("Vous avez %d\n", initiative)
	} else {
		fmt.Printf("❌ L'ennemi le plus proche du chiffre de la Mamma avec une distance de %d, il commence !\n", distEnnemi)
		initiative = initiative - distEnnemi
		fmt.Printf("Vous avez %d\n", initiative)
	}
}
