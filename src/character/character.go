package character

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func DisplayCInfo(c *structures.Character) {
	fmt.Printf("\nNom: %s\nClasse: %s\nNiveau: %d\nPV: %d/%d\nArgent: %d\n",
		c.Name, c.Class, c.Level, c.ActualHp, c.MaxHp, c.Money)
	for _, skill := range c.SkillList {
		fmt.Printf("Liste des skills :\n%s, %d de dégats\n\n", skill.Name, skill.Damage)
	}

}

func DisplayEInfo(e *structures.Enemy) {
	fmt.Printf("\nNom: %s\nPV: %d/%d\n",
		e.Name, e.ActualHp, e.MaxHp)
}

func AccessInventory(c *structures.Character) {
	fmt.Println("\nInventaire :")
	for _, item := range c.Inventory {
		fmt.Printf("%s x%d\n", item.Name, item.Quantity)
	}
	fmt.Println()
}

func IsDead(c *structures.Character) {
	if c.ActualHp <= 0 {
		fmt.Println("Vous êtes mort !")
		c.MaxHp /= 2
		c.ActualHp = c.MaxHp
		fmt.Println("Vous venez de renaître avec 50% de HP en moins.")
	}
}
func IsAlpha(s string) bool {
	for _, c := range s {
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}

func CharacterCreation(c *structures.Character) {
	var username string
	var valid bool

	for {
		fmt.Println("Quel est votre pseudo ?")
		fmt.Scan(&username)

		valid = true
		result := []rune(username)

		if len(result) == 0 {
			fmt.Println("Le pseudo ne peut pas être vide.")
			valid = false
			continue
		}

		for _, r := range result {
			if r < 65 || (r > 90 && r < 97) || r > 122 {
				fmt.Println("Votre pseudo n'est pas correct, il ne contient pas que des lettres.")
				valid = false
				break
			}
		}

		if valid {
			if result[0] >= 97 && result[0] <= 122 {
				result[0] = result[0] - ('a' - 'A')
			}
			for i := 1; i < len(result); i++ {
				if result[i] >= 65 && result[i] <= 90 {
					result[i] = result[i] + ('a' - 'A')
				}
			}

			username = string(result)
			c.Name = username
			break
		}
	}

	fmt.Println("Personnage créé avec le nom :", c.Name)
}
