package score

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/structures"
)

// AddScore ajoute des points au score du personnage en fonction de la difficulté de l'ennemi vaincu.
func AddScore(c *structures.Character, e *structures.Enemy) {
	c.Score += e.GiveScore
	switch e.Difficulty {
	case "Facile":
		c.EasyKill++
	case "Normal":
		c.NormalKill++
	case "Boss":
		c.BossKill++
	}
}

// ShowScore affiche le score actuel du personnage
func ShowScore(c *structures.Character) {
	affichage.Separator()
	fmt.Println("🏆 Tableau des Scores 🏆")
	affichage.Separator()
	fmt.Printf("\nScore de %s : %d\n", c.Name, c.Score)
	fmt.Printf("Tours de Victoire : %d\n", c.VictoryLap)
	fmt.Printf("Ennemis Faciles vaincus : %d\n", c.EasyKill)
	fmt.Printf("Ennemis Normaux vaincus : %d\n", c.NormalKill)
	fmt.Printf("Boss vaincus : %d\n\n", c.BossKill)
}
