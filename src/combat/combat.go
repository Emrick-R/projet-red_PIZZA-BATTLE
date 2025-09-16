package combat

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/score"
	"projet-red_PIZZA-BATTLE/skills"
	"projet-red_PIZZA-BATTLE/structures"
	"time"
)

func DisplayCombatInventory(c *structures.Character, e *structures.Enemy) {
	for {
		character.AccessInventory(c)
		character.AccessEquipement(c)
		character.AccessSkills(c)
		affichage.AffichageMenuInventaire()
		menuChoice := 0
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			for {
				affichage.AffichageMenuCombatPotion()
				menuChoice := 0
				fmt.Scan(&menuChoice)
				switch menuChoice {
				case 1:
					// Utiliser une potion
					items.TakePot(c)
					return
				case 2:
					//Utiliser une potion de poison
					items.ThrowPoisonPot(c, e)
					return
				case 3:
				//Retour
				default:
					fmt.Printf("\nIl ne se passe rien... Choix invalide.\n")
				}
				if menuChoice == 3 {
					menuChoice = 0
					break
				}
			}
		case 2:
			// Equiper un équipement
			character.EquipEquipment(c)
		case 3:
		// Retour
		default:
			fmt.Printf("\nIl ne se passe rien... Choix invalide.\n")
		}
		if menuChoice == 3 {
			menuChoice = 0
			return
		}
	}
}

func CharacterIsDead(c *structures.Character) {
	if c.MaxHp <= 10 {
		fmt.Println("\nTu es mort pour de bon !")
		fmt.Println("Impossibilité de renaître...")
		fmt.Println("========Fin de partie========")
		score.ShowScore(c)
		time.Sleep(7 * time.Second)
		os.Exit(0)
	}
	if c.ActualHp <= 0 {
		fmt.Println("\nTu es mort !")
		c.MaxHp /= 2
		c.ActualHp = c.MaxHp
		fmt.Println("Tu viens de renaître avec 50% de HP en moins.")
		fmt.Printf("PV actuels: %d/%d\n", c.ActualHp, c.MaxHp)
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
	if e.PowerCount == 3 {
		//Tour Spécial
		e.PowerCount = 0
		p := e.Damage * 2
		fmt.Print(e.Name, " attaque ", c.Name, " et lui inflige ", p, " de dégâts\n")
		c.ActualHp = c.ActualHp - p
		fmt.Printf("%s : %d/%d hp\n", c.Name, c.ActualHp, c.MaxHp)
	} else {
		// Autre tours
		c.ActualHp = c.ActualHp - e.Damage
		fmt.Print(e.Name, " attaque ", c.Name, " et lui inflige ", e.Damage, " de dégâts\n")
		fmt.Printf("%s : %d/%d hp\n", c.Name, c.ActualHp, c.MaxHp)
		e.PowerCount++
	}
}

func CharacterTurn(c *structures.Character, e *structures.Enemy) {
	for {
		var combat_choice int
		fmt.Println("======== Combat : ========")
		fmt.Println("Début du combat vous vous battez contre un", e.Name, "! Prudence !")
		fmt.Println("1 - Attaquer ")
		fmt.Println("2 - Inventaire ")
		fmt.Scan(&combat_choice)
		switch combat_choice {
		case 1:
			chosenSkill := skills.SkillChoice(c)
			skills.UseSkill(c, e, chosenSkill)
			fmt.Printf("\n%s a infligé %d points de dégâts à %s\n", c.Name, chosenSkill.Damage, e.Name)
			fmt.Printf("%s : %d/%d hp\n", e.Name, e.ActualHp, e.MaxHp)
			return
		case 2:
			for {
				character.AccessInventory(c)
				character.AccessEquipement(c)
				character.AccessSkills(c)
				affichage.AffichageMenuInventaire()
				menuChoice := 0
				fmt.Scan(&menuChoice)
				switch menuChoice {
				case 1:
					for {
						affichage.AffichageMenuCombatPotion()
						menuChoice := 0
						fmt.Scan(&menuChoice)
						switch menuChoice {
						case 1:
							// Utiliser une potion
							items.TakePot(c)
							return
						case 2:
							//Utiliser une potion de poison
							items.ThrowPoisonPot(c, e)
							return
						case 3:
						//Retour
						default:
							fmt.Printf("\nIl ne se passe rien... Choix invalide.\n")
						}
						if menuChoice == 3 {
							menuChoice = 0
							break
						}
					}
				case 2:
					// Equiper un équipement
					character.EquipEquipment(c)
				case 3:
				// Retour
				default:
					fmt.Printf("\nIl ne se passe rien... Choix invalide.\n")
				}
				if menuChoice == 3 {
					menuChoice = 0
					return
				}
			}
		default:
			fmt.Printf("\nIl ne se passe rien... Choix invalide.\n")
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
			CharacterTurn(c, e)
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
