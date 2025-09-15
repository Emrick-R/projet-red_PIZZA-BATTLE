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

func AccessSkills(c *structures.Character) {
	fmt.Println("\nCompétences :")
	for _, Skill := range c.SkillList {
		fmt.Printf("%s: %d points de dégâts\n", Skill.Name, Skill.Damage)
	}
	fmt.Println()
}

func AccessEquipement(c *structures.Character) {
	fmt.Println("\nEquipement :")
	H := c.Armor.Head.Name
	C := c.Armor.Chest.Name
	L := c.Armor.Legs.Name
	if c.Armor.Head.Name == "" {
		H = "Rien"
	}
	fmt.Println("Armure de tête : ", H)

	if c.Armor.Chest.Name == "" {
		C = "Rien"
	}
	fmt.Println("Armure de torse : ", C)

	if c.Armor.Legs.Name == "" {
		L = "Rien"
	}
	fmt.Println("Armure de jambes : ", L)

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

func EquipEquipment(c *structures.Character, slot string, newEquip structures.Object) {
	chapAvent := structures.Object{Name: "Chapeau de l'aventurier", Quantity: 1}
	fmt.Println("Quel équipement veux-tu équiper ?")
	for _, i := range c.Inventory {
		if i.Name == chapAvent.Name {
		}
		fmt.Println("1 - Tête")
		fmt.Println("2 - Torse")
		fmt.Println("3 - Jambes")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			c.Armor.Head = &newEquip
		case 2:
			c.Armor.Chest = &newEquip
		case 3:
			c.Armor.Legs = &newEquip
		default:
			fmt.Println("Slot d'équipement invalide.")
		}
	}
}

func CharacterCreation(c *structures.Character) {
	var username string
	var valid bool

	for {
		fmt.Println("\nQuel est votre pseudo ?")
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
				fmt.Println("Votre pseudo n'est pas correct, il ne contient que des lettres.")
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

	fmt.Println("\nPersonnage créé avec le nom :", c.Name)
	var class_choice int
	for {
		fmt.Println("\nSuper", c.Name, ", quelle classe veux-tu choisir ?")
		fmt.Println("1 - Elfe : 80 PV Max")
		fmt.Println("2 - Nain : 120 PV Max")
		fmt.Println("3 - Humain : 100 PV Max")
		fmt.Scan(&class_choice)

		if class_choice >= 1 && class_choice <= 3 {
			break
		}
		fmt.Println("Choix invalide, essaye encore.")
	}

	switch class_choice {
	case 1:
		fmt.Println("\nTu as choisi la classe Elfe : agile, précis et en communion avec la nature.")
		c.MaxHp = 80
		c.ActualHp = 40
		c.Class = "Elfe"
	case 2:
		fmt.Println("\nTu as choisi la classe Nain : robuste, courageux et maître de la forge.")
		c.MaxHp = 120
		c.ActualHp = 60
		c.Class = "Nain"
	case 3:
		fmt.Println("\nTu as choisi la classe Humain : polyvalent, ingénieux et déterminé.")
		c.MaxHp = 100
		c.ActualHp = 50
		c.Class = "Humain"
	}
}
