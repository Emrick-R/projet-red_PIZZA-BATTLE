package skills

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

// AddSkills ajoute une compétence à la liste des compétences du personnage
func AddSkills(c *structures.Character, newSkill structures.Skill) {
	c.SkillList = append(c.SkillList, newSkill)
}

// CheckSkills vérifie si le personnage a déjà la compétence
func CheckSkills(c *structures.Character, newSkill structures.Skill) bool { //True = Il a déjà la comp.
	for i := range c.SkillList {
		if c.SkillList[i].Name == newSkill.Name {
			return true
		}
	}
	return false
}

// UseSkill applique les effets de la compétence sur les pv de l'ennemi
func UseSkill(c *structures.Character, e *structures.Enemy, skill structures.Skill) {
	e.ActualHp -= skill.Damage
}

// CheckMana vérifie si le personnage a assez de mana pour utiliser une compétence
func CheckMana(c *structures.Character, skill structures.Skill) {
	if c.ActualMana < skill.ManaCost {
		fmt.Println("Vous n'avez pas assez de mana pour lancer ce sort")
	}

}
