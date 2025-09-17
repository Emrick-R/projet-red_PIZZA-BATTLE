package combat

import (
	"fmt"
	"math/rand"
	"os"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/score"
	"projet-red_PIZZA-BATTLE/skills"
	"projet-red_PIZZA-BATTLE/structures"
	"time"
)

// Fonction valeur absolue
func abs(x int) int {
	// Si x est négatif
	if x < 0 {
		// On retourne son opposé
		return -x
	}
	// Sinon on retourne x
	return x
}

// Fonction pour lancer un dé à 100 faces
func rollDice() int {
	// Génère et retourne un nombre aléatoire entre 1 et 100
	return rand.Intn(100) + 1
}

// Fonction pour déterminer l'initiative en utilisant le mini-jeu "Épreuve de la Mamma"
// Si le joueur gagne, il commence, sinon l'ennemi commence
func InitiativeMamma(c *structures.Character, e *structures.Enemy) bool {

	// Variable du choix du joueur
	var choix int
	affichage.Separator()
	fmt.Println("🟩⬜🟥 Épreuve de la Mamma 🟩⬜🟥")
	affichage.Separator()
	fmt.Println("Choisissez un nombre (1 à 100). Celui le plus proche du score de la Mamma commence !")
	// input joueur sécurisé
	for {
		fmt.Print("👉 Entres ton nombre : ")
		_, err := fmt.Scan(&choix)
		// vérification de l'input
		if err == nil && choix >= 1 && choix <= 100 {
			// input valide, on sort de la boucle
			break
		}
		fmt.Println("❌ Valeur invalide ! Tapes un nombre entre 1 et 100.")
	}

	// premier lancer
	mamma := rollDice()
	ennemi := rollDice()

	// affichage clair
	fmt.Printf("🎲 Ton nombre : %d | 🎲 Mamma : %d | 🎲 Ennemi : %d\n", choix, mamma, ennemi)

	// en cas d'égalité on relance les dés
	for choix == ennemi {
		fmt.Println("Égalité — relance du nombre !")
		mamma = rollDice()
		ennemi = rollDice()
		fmt.Printf("🎲 Chiffre : %d | 🎲 Mamma : %d | 🎲 Ennemi : %d\n", choix, mamma, ennemi)
	}

	// distances absolues (Rappel: la distance est petite == Gagnant)
	distJoueur := abs(choix + c.Initiative - mamma)
	distEnnemi := abs(ennemi + e.Initiative - mamma)

	// affichage du résultat

	if distJoueur < distEnnemi {
		// Joueur gagne
		fmt.Printf("✅ Tu est le plus proche du chiffre de la Mamma avec une distance de %d (Initiative de %d contre %d), vous commencez !\n", distJoueur, c.Initiative, e.Initiative)
		return true
	} else {
		// Ennemi gagne
		fmt.Printf("❌ L'ennemi est plus proche du chiffre de la Mamma avec une distance de %d (Initiative de %d contre %d), il commence !\n", distEnnemi, e.Initiative, c.Initiative)
		return false
	}
}

// Vérification de la mort du personnage avec résurrection (MaxHp/2), si mort définitive (MaxHp <= 10) fin de partie
func CharacterIsDead(c *structures.Character) {
	//Vérification si impossibilité de renaître (MaxHp <= 10)
	if c.MaxHp <= 10 {
		fmt.Println("\n💀 Tu es mort pour de bon !")
		fmt.Println("Impossibilité de renaître...")
		fmt.Println("======== Fin de partie ========")
		//Affichage du score final
		score.ShowScore(c)
		//Pause de 7 secondes avant fermeture du programme
		time.Sleep(7 * time.Second)
		os.Exit(0)
	}

	//Vérification de la mort du personnage puis résurrection avec moitié des PV max
	if c.ActualHp <= 0 {
		fmt.Println("\n💀 Tu es mort !\n")
		//Résurrection avec moitié des PV max
		c.MaxHp /= 2
		c.ActualHp = c.MaxHp
		fmt.Println("✨ Résurrection avec 50% de HP en moins.")
		fmt.Printf("❤️ PV actuels: %d/%d\n\n", c.ActualHp, c.MaxHp)
	}
}

// Vérification de la mort de l'ennemi, si mort renvoie true
func EnemyIsDead(e *structures.Enemy) bool {
	//Si les PV de l'ennemi sont inférieurs ou égaux à 0
	if e.ActualHp <= 0 {
		fmt.Printf("⚔️  Tu as vaincu %s !\n", e.Name)
		//Ennemi mort donc true
		return true
	}
	//Ennemi toujours en vie donc false
	return false
}

// Comportement de l'ennemi lors de son tour
func EnemyPattern(c *structures.Character, e *structures.Enemy, t int) {
	//Tout les 3 tours l'ennemi fait une attaque spéciale (double dégâts)
	if e.PowerCount == 3 {
		//Tour Spécial
		//Remise à 0 du compteur
		e.PowerCount = 0
		//Dégâts doublés sur ce tour
		p := e.Damage * 2
		fmt.Printf("🔥 %s utilise une ATTAQUE SPÉCIALE sur %s et inflige %d dégâts !\n", e.Name, c.Name, p)
		//Application des dégâts
		c.ActualHp = c.ActualHp - p
		//Affichage des PV restants
		fmt.Printf("❤️ %s : %d/%d HP\n", c.Name, c.ActualHp, c.MaxHp)
	} else {
		//Autre tours
		//Attaque normale
		c.ActualHp = c.ActualHp - e.Damage
		fmt.Printf("👊 %s attaque %s et inflige %d dégâts\n", e.Name, c.Name, e.Damage)
		fmt.Printf("❤️ %s : %d/%d HP\n", c.Name, c.ActualHp, c.MaxHp)
		//Incrémentation du compteur de l'attaque spéciale
		e.PowerCount++
	}
}

// Tour du personnage
func CharacterTurn(c *structures.Character, e *structures.Enemy) {
	for {
		//Boucle infinie jusqu'à un retour (fin du tour)
		var combat_choice int
		//Affichage du menu combat
		affichage.Separator()
		fmt.Println("⚔️  COMBAT :")
		affichage.Separator()
		fmt.Println("1 - 🗡️ Attaquer")
		fmt.Println("2 - 🎒 Inventaire")
		fmt.Println("3 - 💀 Suicide")

		fmt.Print("👉 Ton choix : ")
		fmt.Scan(&combat_choice)
		switch combat_choice {
		case 1:
			// Attaque
			chosenSkill := skills.SkillChoice(c)
			if skills.CheckMana(c, chosenSkill) {
				c.ActualMana -= chosenSkill.ManaCost
				if c.ActualMana < 0 {
					c.ActualMana = 0
				}
				fmt.Printf("🔵 Mana restant : %d/%d\n", c.ActualMana, c.ManaMax)
				skills.UseSkill(c, e, chosenSkill)
				fmt.Printf("\n💥 %s inflige %d points de dégâts à %s !\n", c.Name, chosenSkill.Damage, e.Name)
				fmt.Printf("❤️ %s : %d/%d HP\n", e.Name, e.ActualHp, e.MaxHp)
				return
			} else {
				fmt.Println("❌ Pas assez de mana pour utiliser cette compétence !")
				return
			}
		case 2:
			for {
				// Affichage de l'inventaire

				// Affichage de l'inventaire via une fonction
				character.AccessInventory(c)
				// Affichage de l'équipement
				character.AccessEquipement(c)
				// Affichage des compétences
				character.AccessSkills(c)
				// Affichage des choix
				affichage.AffichageMenuInventaire()
				menuChoice := 0
				fmt.Print("👉 Ton choix : ")
				fmt.Scan(&menuChoice)
				switch menuChoice {
				case 1:
					for {
						// Affichage des choix de potions
						affichage.AffichageMenuCombatPotion()
						menuChoice := 0
						fmt.Print("👉 Ton choix : ")
						fmt.Scan(&menuChoice)
						switch menuChoice {
						case 1:
							HpPot := structures.Object{Name: "Potion de Vie", Quantity: 1}
							for i := 0; i < len(c.Inventory); i++ {
								if c.Inventory[i].Name == HpPot.Name {
									//Utiliser une potion de vie
									items.TakePot(c)
									// Fin du tour du joueur
									return
								} else {
									fmt.Println("❌ Il n'y a pas de vie de poison dans l'inventaire.")
								}
							}
						case 2:
							PoisonPot := structures.Object{Name: "Potion de Poison"}
							for i := 0; i < len(c.Inventory); i++ {
								if c.Inventory[i].Name == PoisonPot.Name {
									//Utiliser une potion de poison
									items.ThrowPoisonPot(c, e)
									// Fin du tour du joueur
									return
								} else {
									fmt.Println("❌ Il n'y a pas de potion de poison dans l'inventaire.")
								}
							}

						case 3:
						//Retour
						default:
							// Choix autre que 1, 2 ou 3
							fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
						}
						//Reset de la variable menuChoice
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
					// Choix autre que 1, 2 ou 3
					fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
				}
				//Reset de la variable menuChoice
				if menuChoice == 3 {
					menuChoice = 0
					return
				}
			}
		case 3:
			c.ActualHp = 0
			return
		default:
			// Choix autre que 1 ou 2
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}

	}
}

// Fonction principale du combat 1v1 entre le personnage et l'ennemi
func TurnCombat1v1(c *structures.Character, e *structures.Enemy) {
	// Initialisation du tour
	Turn := 1
	TrueTurn := 1
	//Initialisation de l'initiative
	if InitiativeMamma(c, e) {
		Turn = 2
	} else {
		Turn = 1
	}
	//Boucle de combat
	for {
		//Le tour du joueur (Turn == pair)
		if Turn%2 == 0 {
			affichage.Separator()
			fmt.Printf("🎯 Tour %d — À ton tour %s !\n", TrueTurn, c.Name)
			affichage.Separator()
			//Déroulement du tour du joueur
			CharacterTurn(c, e)
			//Vérification de la mort
			if EnemyIsDead(e) {
				//L'ennemi est mort
				break
			}
			CharacterIsDead(c)
			Turn++
			TrueTurn++

		} else {
			//Le tour de l'IA (Turn == impair)
			affichage.Separator()
			fmt.Printf("🎯 Tour %d — C'est au tour de %s !\n", TrueTurn, e.Name)
			affichage.Separator()
			//Verification de l'effet de poison
			items.CheckPoisonStatus(e)
			if EnemyIsDead(e) {
				//L'ennemi est mort
				break
			}
			//Déroulement du tour de l'IA
			EnemyPattern(c, e, Turn)
			//Vérification de la mort
			CharacterIsDead(c)
			Turn++
			TrueTurn++
		}
	}
	//Fin du combat (ennemi mort)
	fmt.Printf("🏆 Bravo ! Tu as terrassé %s !\n", e.Name)
	e.ActualHp = e.MaxHp
	//Récompenses du combat (Argent + Score)
	score.AddScore(c, e)
	inventory.AddMoney(c, e)
	character.AddExp(c, e)
	fmt.Printf("\n💰 +%d argent", e.GiveMoney)
	fmt.Printf("\n📚 +%d expérience", e.GiveExp)
	fmt.Printf("\n⭐ +%d points de score\n", e.GiveScore)
	//Affichage de l'argent, de l'Exp et du score
	affichage.Separator()
	character.NextLevel(c)
	fmt.Printf("💵 Argent : %d | 📖 Exp : %d | 🏅 Score : %d\n", c.Money, c.ActualExp, c.Score)
	affichage.Separator()
	// Retour au menu principal
}
