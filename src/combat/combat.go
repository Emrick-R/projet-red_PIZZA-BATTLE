package combat

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/score"
	"projet-red_PIZZA-BATTLE/skills"
	"projet-red_PIZZA-BATTLE/structures"
	"time"
)

func DisplayCombatInventory(c *structures.Character) {
	fmt.Println("Voici ton inventaire :")
	for i := range c.Inventory {
		fmt.Printf("%d - %s (x%d)\n", i+1, c.Inventory[i].Name, c.Inventory[i].Quantity)
	}
}

func CharacterIsDead(c *structures.Character) {
	if c.MaxHp <= 10 {
		fmt.Println("Tu es mort pour de bon !")
		fmt.Println("Impossibilité de renaître...")
		fmt.Println("========Fin de partie========")
		score.ShowScore(c)
		time.Sleep(15 * time.Second)
		os.Exit(0)
	}
	if c.ActualHp <= 0 {
		fmt.Println("Tu es mort !")
		c.MaxHp /= 2
		c.ActualHp = c.MaxHp
		fmt.Println("Tu viens de renaître avec 50% de HP en moins.")
	}
}

func EnemyIsDead(e *structures.Enemy) bool {
	if e.ActualHp <= 0 {
		fmt.Printf("Tu as vaincus %s !\n", e.Name)
		return true
	}
	return false
}
func EnemyPatern(c *structures.Character, e *structures.Enemy, t int) {
	/*Tour = T
	T1 = Enemy attaque 100% des dégats
	T2 = 100% notre perso
	T3 = 200% enemy attaque
	T4 = 100% notre perso
	T3*/
	if t%3 == 0 {
		fmt.Print(e.Name, " attaque ", c.Name, " et lui inflige ", e.Damage*2, " de dégâts")
		c.ActualHp = c.ActualHp - (e.Damage * 2)
		fmt.Printf("%s : %d/%d hp\n", c.Name, c.ActualHp, c.MaxHp)
	} else {
		// Autre tours
		c.ActualHp = c.ActualHp - e.Damage
		fmt.Print(e.Name, " attaque ", c.Name, " et lui inflige ", e.Damage, " de dégâts")
		fmt.Printf("%s : %d/%d hp\n", c.Name, c.ActualHp, c.MaxHp)
	}
}

func Combat(c *structures.Character, e *structures.Enemy) {
	for {
		var combat_choice int
		fmt.Print("Début du combat vous vous battez contre un", e.Name, "! Prudence !")
		fmt.Printf("1 - Attaquer ")
		fmt.Printf("2 - Inventaire ")
		fmt.Printf("3 - Fuir devant la puissance de l'ennemi")
		fmt.Scan(&combat_choice)
		switch combat_choice {
		case 1:
			skills.SkillChoice(c)
		case 2:
			affichage.AffichageMenuInventaire()
		case 3:
			fmt.Printf("Vous venez de fuir !")
		}

	}
}

func TurnCombat1v1(c *structures.Character, e *structures.Enemy) {
	/* 1v1 =
	Enemy commence tour 1 et chaque tour impair (T1,T3,T5,T7 etc.)
	Joueur commence tour 2 et chaque Tour pair (T2,T4,T6,T8,etc.)*/
	Turn := 1
	for {
		if Turn%2 == 0 { //Le tour du joueur, donc Tour 2
			fmt.Println("\nTour :", Turn)
			fmt.Printf("A ton tour %s!\n\n", c.Name)
			Combat(c, e)
			if EnemyIsDead(e) {
				score.Addscore(c, e)
			}
			Turn++
		} else { //Le tour de l'IA, donc Tour 2
			fmt.Println("\nTour :", Turn)
			fmt.Printf("C'est au tour de %s \n\n", e.Name)
			EnemyPatern(c, e, Turn)
			CharacterIsDead(c)
			Turn++
		}
	}

}
