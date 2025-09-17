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
	// Variables locales
	var choix, distJoueur, distEnnemi, mamma, ennemi, counter int

	// Affichage du mini-jeu
	affichage.Separator()
	fmt.Println("🟩⬜🟥 Épreuve de la Mamma 🟩⬜🟥")
	affichage.Separator()
	fmt.Println("Choisissez un nombre (1 à 100). Celui le plus proche du score de la Mamma commence !")

	// Input joueur sécurisé
	for {
		fmt.Print("👉 Entres ton nombre : ")
		_, err := fmt.Scan(&choix)
		if err == nil && choix >= 1 && choix <= 100 {
			break
		}
		fmt.Print("\033[H\033[2J")
		fmt.Println("❌ Valeur invalide ! Tapes un nombre entre 1 et 100.")

		time.Sleep(2 * time.Second)

	}

	// Premier lancer
	mamma = rollDice()
	ennemi = rollDice()
	fmt.Printf("🎲 Ton nombre : %d | 🎲 Mamma : %d | 🎲 Ennemi : %d\n", choix, mamma, ennemi)

	// Gestion des égalités initiales
	for choix == ennemi {
		fmt.Println("Égalité — relance du nombre !")
		mamma = rollDice()
		ennemi = rollDice()
		fmt.Printf("🎲 Ton nombre : %d | 🎲 Mamma : %d | 🎲 Ennemi : %d\n", choix, mamma, ennemi)
	}

	// Calcul des distances
	distJoueur = abs(choix-mamma) - c.Initiative
	distEnnemi = abs(ennemi-mamma) - e.Initiative

	if distJoueur < 0 {
		distJoueur = 0
	}
	if distEnnemi < 0 {
		distEnnemi = 0
	}
	fmt.Printf("Distance après initiative %s : %d | Distance après initiative %s : %d\n", c.Name, distJoueur, e.Name, distEnnemi)

	// Gestion des égalités après initiative
	for distJoueur == distEnnemi {
		counter++
		if counter == 3 {
			// Après 3 égalités, le joueur gagne par défaut
			fmt.Println("⚠️  Trop d'égalités, tu gagnes par défaut !")

			time.Sleep(2 * time.Second)

			return true
		}

		fmt.Println("Égalité avec initiative — relance du nombre !")
		mamma = rollDice()
		ennemi = rollDice()

		distJoueur = abs(choix-mamma) - c.Initiative
		distEnnemi = abs(ennemi-mamma) - e.Initiative

		if distJoueur < 0 {
			distJoueur = 0
		}
		if distEnnemi < 0 {
			distEnnemi = 0
		}
		fmt.Printf("Distance après initiative %s : %d | Distance après initiative %s : %d\n", c.Name, distJoueur, e.Name, distEnnemi)
	}

	// Retour du résultat : true si joueur commence
	if distJoueur < distEnnemi {
		// Joueur gagne
		fmt.Printf("✅ Tu est le plus proche du chiffre de la Mamma avec une distance de %d contre %d (Initiative de %d contre %d), vous commencez !\n", distJoueur, distEnnemi, c.Initiative, e.Initiative)

		time.Sleep(2 * time.Second)

		return true
	} else {
		// Ennemi gagne
		fmt.Printf("❌ L'ennemi est plus proche du chiffre de la Mamma avec une distance de %d contre %d (Initiative de %d contre %d), il commence !\n", distEnnemi, distJoueur, e.Initiative, c.Initiative)

		time.Sleep(2 * time.Second)

		return false
	}
}

// Vérification de la mort du personnage avec résurrection (MaxHp/2), si mort définitive (MaxHp <= 10) fin de partie
func CharacterIsDead(c *structures.Character) {
	//Vérification si impossibilité de renaître (MaxHp <= 10)
	if c.MaxHp <= 10 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("\n💀 Tu es mort pour de bon !")
		fmt.Println("🪦 Impossibilité de renaître...")
		affichage.Separator()
		fmt.Println("🎮 Fin de partie")
		affichage.Separator()
		//Affichage du score final
		score.ShowScore(c)
		affichage.Separator()
		//Pause de 7 secondes avant fermeture du programme
		time.Sleep(7 * time.Second)
		os.Exit(0)
	}

	//Vérification de la mort du personnage puis résurrection avec moitié des PV max
	if c.ActualHp <= 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("\n💀 Tu es mort !\n\n")
		//Résurrection avec moitié des PV max
		c.MaxHp /= 2
		c.ActualHp = c.MaxHp
		fmt.Println("✨ Résurrection avec 50% de HP en moins.")
		fmt.Printf("❤️  PV actuels: %d/%d\n\n", c.ActualHp, c.MaxHp)

		time.Sleep(2 * time.Second)

	}
}

// Vérification de la mort de l'ennemi, si mort renvoie true
func EnemyIsDead(e *structures.Enemy) bool {
	//Si les PV de l'ennemi sont inférieurs ou égaux à 0
	if e.ActualHp <= 0 {
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

		time.Sleep(2 * time.Second)

	} else {
		//Autre tours
		//Attaque normale
		c.ActualHp = c.ActualHp - e.Damage
		fmt.Printf("👊 %s attaque %s et inflige %d dégâts\n", e.Name, c.Name, e.Damage)
		fmt.Printf("❤️  %s : %d/%d HP\n", c.Name, c.ActualHp, c.MaxHp)
		//Incrémentation du compteur de l'attaque spéciale
		e.PowerCount++

		time.Sleep(2 * time.Second)

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
		fmt.Println("1 - 🗡️  Attaquer")
		fmt.Println("2 - 🎒 Inventaire")
		fmt.Println("3 - 👹 Information ennemi")
		fmt.Println("4 - 💀 Suicide")
		affichage.Separator()

		fmt.Print("👉 Ton choix : ")
		fmt.Scan(&combat_choice)
		switch combat_choice {
		case 1:
			// Attaque
			for {
				// Choix de la compétence
				var chosenSkill structures.Skill
				var skill_choice int
				var index int
				affichage.Separator()
				fmt.Println("👊  Quelle compétence veux-tu utiliser ?")
				affichage.Separator()
				fmt.Printf("🍅 %s : %d/%d Sauce Tomate\n\n", c.Name, c.ActualMana, c.MaxMana)
				// Affiche la liste des compétences disponibles
				for i := range c.SkillList {
					fmt.Printf("%d - %s: %d Dégats, -%d Sauce Tomate\n", i+1, c.SkillList[i].Name, c.SkillList[i].Damage, c.SkillList[i].ManaCost)
					index = len(c.SkillList) + 1
				}
				fmt.Printf("%d - ⬅️  RETOUR\n", index)
				fmt.Print("👉 Ton choix : ")
				fmt.Scan(&skill_choice)
				// Vérifie que le choix est valide dans la liste des compétences
				if skill_choice >= 1 && skill_choice <= len(c.SkillList) {
					// Retourne la compétence choisie (indexée à partir de 0)
					chosenSkill = c.SkillList[skill_choice-1]

					// Vérification du mana
					if skills.CheckMana(c, chosenSkill) {
						c.ActualMana -= chosenSkill.ManaCost
						if c.ActualMana < 0 {
							c.ActualMana = 0
						}
						skills.UseSkill(c, e, chosenSkill)
						fmt.Printf("\n🍅 Sauce Tomate restant : %d/%d\n", c.ActualMana, c.MaxMana)
						fmt.Printf("\n💥 %s utilise %s et inflige %d points de dégâts à %s !\n", c.Name, chosenSkill.Name, chosenSkill.Damage, e.Name)
						fmt.Printf("❤️ %s : %d/%d HP\n", e.Name, e.ActualHp, e.MaxHp)

						time.Sleep(2 * time.Second)

						return

					} else {
						fmt.Println("❌ Pas assez de Sauce Tomate pour utiliser cette compétence !")

						time.Sleep(2 * time.Second)

					}
				} else if skill_choice == index {
					// Retour

					// Reset de la variable chosenSkill
					skill_choice = 0
					// Reset de la variable skill_choice
					index = 0
					// Sortie de la boucle
					break
				} else {
					// Si le choix n'est pas valide, affiche un message d'erreur et redemande
					fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")

					time.Sleep(2 * time.Second)

				}
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
							// Utiliser une potion de vie
							if c.ActualHp == c.MaxHp {
								fmt.Printf("\n❌ Les points de vie sont déjà au max\n\n")

								time.Sleep(2 * time.Second)

							} else {
								HpPot := structures.Object{Name: "Tiramisu"}
								for i := 0; i < len(c.Inventory); i++ {
									if c.Inventory[i].Name == HpPot.Name {
										//Vérification si le personnage a déjà les PV max
										if c.ActualHp == c.MaxHp {
											//Message d'erreur si les PV sont déjà au max
											fmt.Printf("\n❌ Les points de vie sont déjà au max\n\n")

											time.Sleep(2 * time.Second)

										} else {
											//Utiliser une potion de vie
											items.TakePot(c)

											time.Sleep(2 * time.Second)

											// Fin du tour du joueur
											return
										}
									}
								}
								fmt.Println("❌ Il n'y a pas de Tiramisu dans l'inventaire.")

								time.Sleep(2 * time.Second)

							}
						case 2:
							// Utiliser une potion de poison
							PoisonPot := structures.Object{Name: "Tabasco"}
							for i := 0; i < len(c.Inventory); i++ {
								if c.Inventory[i].Name == PoisonPot.Name {
									//Utiliser une potion de poison
									items.ThrowPoisonPot(c, e)

									time.Sleep(2 * time.Second)

									// Fin du tour du joueur
									return
								}
							}
							fmt.Println("❌ Il n'y a pas de Tabasco dans l'inventaire.")

							time.Sleep(2 * time.Second)

						case 3:
							// Utiliser une potion de mana.
							if c.ActualMana == c.MaxMana {
								fmt.Printf("\n❌ La Sauce Tomate est déjà pleine\n\n")

								time.Sleep(2 * time.Second)

							} else {

								ManaPot := structures.Object{Name: "Bocal de Sauce Tomate"}
								for i := 0; i < len(c.Inventory); i++ {
									if c.Inventory[i].Name == ManaPot.Name {
										if c.ActualMana == c.MaxMana {
											fmt.Printf("\n❌ La Sauce Tomate est déjà pleine\n\n")

											time.Sleep(2 * time.Second)

										} else {
											items.TakeManaPot(c)

											time.Sleep(2 * time.Second)

											return
										}
									}
								}
								fmt.Println("❌ Il n'y a pas de Bocal de Sauce Tomate dans l'inventaire.")

								time.Sleep(2 * time.Second)

							}
						case 4:
						//Retour
						default:
							// Choix autre que 1, 2, 3 ou 4
							fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")

							time.Sleep(2 * time.Second)

						}
						//Reset de la variable menuChoice
						if menuChoice == 4 {
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

					time.Sleep(2 * time.Second)

				}
				//Reset de la variable menuChoice
				if menuChoice == 3 {
					menuChoice = 0
					break
				}
			}
		case 3:
			// Afficher les informations de l'ennemi
			character.DisplayEInfo(e)
		case 4:
			c.ActualHp = 0
			return
		default:
			// Choix autre que 1 ou 2
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")

			time.Sleep(2 * time.Second)

		}

	}
}

// SetDifficulty ajuste les statistiques de l'ennemi en fonction de la difficulté choisie par le joueur
func SetDifficulty(c *structures.Character, e *structures.Enemy) {
	e.MaxHp = c.Difficulty * e.MaxHp
	e.Damage = c.Difficulty * e.Damage
	e.ActualHp = e.MaxHp
}

// Fonction principale du combat 1v1 entre le personnage et l'ennemi
func TurnCombat1v1(c *structures.Character) {

	// Effacer l'écran
	fmt.Print("\033[H\033[2J")

	e := &structures.Enemy{}
	// Initialisation de l'ennemi en fonction de la progression du joueur
	if c.Progress == 1 {
		fmt.Println("Tu affronte le Petit Giovanni !, Prudence...")

		time.Sleep(2 * time.Second)

		e = structures.InitEnemy("Petit Giovanni", "Facile")
	} else if c.Progress > 1 && c.Progress < 5 {
		r := rand.Intn(2)
		if r == 0 {
			fmt.Println("Tu affronte El Don Pastabox 3000 !, Attention !")

			time.Sleep(2 * time.Second)

			e = structures.InitEnemy("El Don Pastabox 3000", "Normale")
		} else {
			fmt.Println("Tu affronte le Petit Giovanni !, Prudence...")

			time.Sleep(2 * time.Second)

			e = structures.InitEnemy("Petit Giovanni", "Facile")
		}
	} else if c.Progress == 6 {
		fmt.Println("Tu affronte le Ultra Mega Hyper Giovanni EX Turbo GX !, les enjeux sont à leur comble !")

		time.Sleep(2 * time.Second)

		e = structures.InitEnemy("Ultra Mega Hyper Giovanni EX Turbo GX", "Boss")
	}

	// Initialisation de la difficulté de l'ennemi en fonction de la progression du joueur
	SetDifficulty(c, e)

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
			fmt.Print("\033[H\033[2J")
			//Affichage du tour
			affichage.Separator()
			fmt.Printf("🎯 Tour %d — À ton tour %s !\n", TrueTurn, c.Name)
			affichage.Separator()

			time.Sleep(2 * time.Second)

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
			fmt.Print("\033[H\033[2J")
			//Le tour de l'IA (Turn == impair)
			affichage.Separator()
			fmt.Printf("🎯 Tour %d — C'est au tour de %s !\n", TrueTurn, e.Name)
			affichage.Separator()

			time.Sleep(2 * time.Second)

			//Verification de l'effet de poison
			items.CheckPoisonStatus(e)
			if EnemyIsDead(e) {
				//L'ennemi est mort
				break
			}

			time.Sleep(2 * time.Second)

			//Déroulement du tour de l'IA
			EnemyPattern(c, e, Turn)

			time.Sleep(2 * time.Second)

			//Vérification de la mort
			CharacterIsDead(c)
			Turn++
			TrueTurn++
		}
	}
	//Le combat est terminé

	//Le joueur a gagné
	// Fin de partie si le boss final est vaincu
	//
	if c.Progress == 6 {

		// Effacer l'écran
		fmt.Print("\033[H\033[2J")
		//Reset de la progression
		c.Progress = 1
		endChoice := 0
		score.ShowScore(c)
		affichage.GameEnd()
		for {
			fmt.Print("👉 Ton choix : ")
			fmt.Scan(&endChoice)
			switch endChoice {
			case 1:
				c.Difficulty++
				c.VictoryLap++

			case 2:
				// Quitter le jeu
				os.Exit(0)
			default:
				fmt.Print("\033[H\033[2J")
				// Choix autre que 1 ou 2
				fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
			}
			if endChoice == 1 {
				endChoice = 0
				break
			}
		}
	}
	// Effacer l'écran
	fmt.Print("\033[H\033[2J")

	//Fin du combat (ennemi mort)
	affichage.Separator()
	fmt.Printf("🏆 Bravo ! Tu as terrassé %s !\n", e.Name)
	affichage.Separator()

	time.Sleep(2 * time.Second)

	//Récompenses du combat (Argent + Score)
	score.AddScore(c, e)
	inventory.AddMoney(c, e)
	character.AddExp(c, e)
	fmt.Printf("💵 +%d argent", e.GiveMoney)
	fmt.Printf("\n📚 +%d expérience", e.GiveExp)
	fmt.Printf("\n⭐ +%d points de score\n", e.GiveScore)

	time.Sleep(2 * time.Second)

	//Affichage de l'argent, de l'Exp et du score
	affichage.Separator()
	character.NextLevel(c)
	fmt.Printf("💵 Argent : %d | 📖 Exp : %d | 🏅 Score : %d\n", c.Money, c.ActualExp, c.Score)
	affichage.Separator()

	time.Sleep(2 * time.Second)

	// Retour au menu principal
}
