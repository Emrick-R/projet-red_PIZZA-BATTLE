package skills

import (
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
