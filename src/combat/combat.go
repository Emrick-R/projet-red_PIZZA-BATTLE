package combat

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

func EnemyPatern(c *structures.Character, e *structures.enemy) {
	// T1
	c.ActualHp = c.ActualHp - e.Damage
	fmt.Print(e.Name, "attaque", c.Name, "et lui inflige", e.Damage, "de dégâts")

	// T3
	//if nombre de tours % 3 == 0 alors
	fmt.Print(e.Name, "attaque", c.Name, "et lui inflige", e.Damage*2, "de dégâts")
	c.ActualHp = c.ActualHp - (e.Damage * 2)

}

/*Tour = T
T1 = Enemy attaque 100% des dégats
T2 = 100% notre perso
T3 = 200% enemy attaque
T4 = 100% notre perso
*/

func TurnCombat1v1() {
	/* 1v1 =
	Enemy commence tour 1 et chaque tour+2 (T1,T3,T5,T7 etc.)
	Joueur commence tour 2 et chaque Tour+2 (T2,T4,T6,T8,etc.)*/
	Turn := 1
	if Turn%2 == 0 { //Le tour du joueur, donc Tour 2
		//Appel de la fonction tour du Joueur
		Turn++
	} else { //Le tour de l'IA, donc Tour 2
		//Appel de la fonction du tour de l'IA
		Turn++
	}
}
