package main

import "fmt"

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
