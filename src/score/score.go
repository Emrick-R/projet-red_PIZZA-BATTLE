package score

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func AddScore(c *structures.Character, e *structures.Enemy) int {
	switch e.Difficulty {
	case "Facile":
		c.Score += 5
		return 5
	case "Normal":
		c.Score += 10
		return 10
	case "Boss":
		c.Score += 20
		return 20
	}
	return 0
}

func ShowScore(c *structures.Character) {
	fmt.Printf("\n\nScore de %s : %d\n\n", c.Name, c.Score)
}
