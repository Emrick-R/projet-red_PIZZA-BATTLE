package combat

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/structures"
)

func EnemyPatern(c *structures.Character, e *structures.Enemy, t int) {
	/*Tour = T
	T1 = Enemy attaque 100% des dégats
	T2 = 100% notre perso
	T3 = 200% enemy attaque
	T4 = 100% notre perso

	T3*/
	if t%3 == 0 {
		fmt.Print(e.Name, "attaque", c.Name, "et lui inflige", e.Damage*2, "de dégâts")
		c.ActualHp = c.ActualHp - (e.Damage * 2)
	} else {
		// Autre tours
		c.ActualHp = c.ActualHp - e.Damage
		fmt.Print(e.Name, "attaque", c.Name, "et lui inflige", e.Damage, "de dégâts")
	}
}

func TurnCombat1v1(c *structures.Character, e *structures.Enemy) {
	/* 1v1 =
	Enemy commence tour 1 et chaque tour impair (T1,T3,T5,T7 etc.)
	Joueur commence tour 2 et chaque Tour pair (T2,T4,T6,T8,etc.)*/
	Turn := 1
	if Turn%2 == 0 { //Le tour du joueur, donc Tour 2
		fmt.Printf("A ton tour %s!\n", c.Name)
		//Appel de la fonction tour du Joueur
		character.CharacterIsDead(c)
		Turn++
	} else { //Le tour de l'IA, donc Tour 2
		fmt.Printf("C'est au tour de %s \n", e.Name)
		EnemyPatern(c, e, Turn)
		Turn++
	}
}
