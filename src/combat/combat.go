package combat

func EnemyPatern() {
	/*Tour = T
	T1 = Enemy attaque 100% des dégats
	T2 = 100%
	T3 = 200%
	T4 = 100%
	*/
}

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
