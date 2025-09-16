package score

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func Addscore(c *structures.Character, e *structures.Enemy) {
	switch e.Difficulty {
	case "Facile":
		c.Score += 5
	case "Normal":
		c.Score += 10
	case "Boss":
		c.Score += 20
	}
}

func ShowScore(c *structures.Character) {
	fmt.Printf("Score de %s : %d\n", c.Name, c.Score)
}
