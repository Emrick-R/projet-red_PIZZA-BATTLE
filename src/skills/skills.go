package skills

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func AddSkills(c *structures.Character, newSkill structures.Skill) {
	c.SkillList = append(c.SkillList, newSkill)
}

func CheckSkills(c *structures.Character, newSkill structures.Skill) bool { //True = Il a déjà la comp.
	for i := range c.SkillList {
		if c.SkillList[i].Name == newSkill.Name {
			return true
		}
	}
	return false
}

func SkillChoice(c *structures.Character) structures.Skill {
	var skill_choice int
	for {
		fmt.Println("\nQuelle compétence veux-tu utiliser ?")
		for i := range c.SkillList {
			fmt.Printf("%d - %s\n", i+1, c.SkillList[i].Name)
		}
		fmt.Scan(&skill_choice)
		if skill_choice >= 1 && skill_choice <= len(c.SkillList) {
			break
		}
		fmt.Printf("\nIl ne se passe rien... Choix invalide.\n")
	}
	return c.SkillList[skill_choice-1]
}

func UseSkill(c *structures.Character, e *structures.Enemy, skill structures.Skill) {
	e.ActualHp -= skill.Damage
}
