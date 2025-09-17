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

// SkillChoice permet au joueur de choisir une compétence à utiliser et retourne la compétence choisie
func SkillChoice(c *structures.Character) structures.Skill {
	var skill_choice int
	// Affiche la liste des compétences disponibles
	for {
		fmt.Println("\nQuelle compétence veux-tu utiliser ?")
		for i := range c.SkillList {
			fmt.Printf("%d - %s\n", i+1, c.SkillList[i].Name)
		}
		fmt.Print("👉 Ton choix : ")
		fmt.Scan(&skill_choice)
		// Vérifie que le choix est valide dans la liste des compétences
		if skill_choice >= 1 && skill_choice <= len(c.SkillList) {
			// Retourne la compétence choisie (indexée à partir de 0)
			return c.SkillList[skill_choice-1]
		}
		// Si le choix n'est pas valide, affiche un message d'erreur et redemande
		fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
	}
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
