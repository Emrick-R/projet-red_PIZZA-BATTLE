package score

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

// AddScore ajoute des points au score du personnage en fonction de la difficulté de l'ennemi vaincu
func AddScore(c *structures.Character, e *structures.Enemy) {
	c.Score += e.GiveScore
}

// ShowScore affiche le score actuel du personnage
func ShowScore(c *structures.Character) {
	fmt.Printf("\nScore de %s : %d\n\n", c.Name, c.Score)
}
