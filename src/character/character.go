package character

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/affichage"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/structures"
)

// DisplayCInfo affiche les informations du personnage (nom, classe, PV, niveau, expérience, argent, initiative, score)
func DisplayCInfo(c *structures.Character) {
	fmt.Printf("\nNom: %s\nClasse: %s, PV: %d/%d\n Niveau: %d, Expérience: %d/%d\nArgent: %d\nInitiative: %d\nScore: %d\n",
		c.Name, c.Class, c.ActualHp, c.MaxHp, c.Level, c.ActualExp, c.MaxExp, c.Money, c.Initiative, c.Score)
}

// DisplayEInfo affiche les informations de l'ennemi (nom, PV, initiative)
func DisplayEInfo(e *structures.Enemy) {
	fmt.Printf("\nNom: %s\nPV: %d/%d\nInitiative: %d", e.Name, e.ActualHp, e.MaxHp, e.Initiative)
}

// AccessInventory affiche le contenu de l'inventaire du personnage (nom et quantité des objets)
func AccessInventory(c *structures.Character) {
	fmt.Println("\nInventaire :")
	for i := range c.Inventory {
		fmt.Printf("- %s (x%d)\n", c.Inventory[i].Name, c.Inventory[i].Quantity)
	}
}

// AccessSkills affiche la liste des compétences du personnage (nom et dégâts)
func AccessSkills(c *structures.Character) {
	fmt.Println("\nCompétences :")
	for s := range c.SkillList {
		fmt.Printf("%s: %d points de dégâts\n", c.SkillList[s].Name, c.SkillList[s].Damage)
	}
	fmt.Println()
}

// AccessEquipement affiche l'équipement porté par le personnage (nom de l'armure de tête, torse et jambes)
func AccessEquipement(c *structures.Character) {
	fmt.Println("\nEquipement :")
	H := c.Armor.Head.Name
	C := c.Armor.Chest.Name
	L := c.Armor.Legs.Name
	if c.Armor.Head.Name == "" {
		H = "Rien"
	}
	fmt.Println("Armure de tête : ", H)

	if c.Armor.Chest.Name == "" {
		C = "Rien"
	}
	fmt.Println("Armure de torse : ", C)

	if c.Armor.Legs.Name == "" {
		L = "Rien"
	}
	fmt.Println("Armure de jambes : ", L)

}

// InventoryChoice affiche l'inventaire, les equipements equipés, la liste des compétences et
// permet au joueur de choisir une action (utiliser une potion, équiper un équipement, ou revenir en arrière)
func InventoryChoice(c *structures.Character) {
	for {
		AccessInventory(c)
		AccessEquipement(c)
		AccessSkills(c)
		affichage.AffichageMenuInventaire()
		menuChoice := 0
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case 1:
			// Utiliser une potion
			items.TakePot(c)
		case 2:
			// Equiper un équipement
			EquipEquipment(c)
		case 3:
		// Retour
		default:
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}
		if menuChoice == 3 {
			menuChoice = 0
			return
		}
	}
}

// EquipEquipment permet au joueur d'équiper un équipement (chapeau, tunique, bottes)
// s'il le possède dans son inventaire
func EquipEquipment(c *structures.Character) {
	//Itinialisation des équipements disponibles
	var newEquipChoice int
	chapAvent := structures.Object{Name: "Chapeau de l'aventurier", Quantity: 1}
	tunAvent := structures.Object{Name: "Tunique de l'aventurier", Quantity: 1}
	botAvent := structures.Object{Name: "Bottes de l'aventurier", Quantity: 1}
	for {
		// Affichage du menu d'équipement
		fmt.Println("\nQuel équipement veux-tu porter ?")
		fmt.Println("1 - Chapeau de l'aventurier")
		fmt.Println("2 - Tunique de l'aventurier")
		fmt.Println("3 - Bottes de l'aventurier")
		fmt.Println("4 - RETOUR")
		fmt.Scan(&newEquipChoice)
		// Vérification si le joueur possède l'équipement dans son inventaire
		// Iniatialisation des booléens pour savoir si le joueur possède l'équipement
		hadChap := false
		hadTun := false
		hadBot := false
		// Parcours de l'inventaire pour vérifier la présence des équipements
		// et mise à jour des booléens en conséquence
		// Si le joueur possède l'équipement, il peut l'équiper
		for _, i := range c.Inventory {
			if i.Name == chapAvent.Name {
				hadChap = true
			}
			if i.Name == tunAvent.Name {
				hadTun = true
			}
			if i.Name == botAvent.Name {
				hadBot = true
			}
		}

		// Équipement de l'objet choisi si le joueur le possède
		// Sinon, affichage d'un message d'erreur
		// Retour au menu précédent si le choix est 4
		switch newEquipChoice {
		case 1:
			// Équiper le chapeau de l'aventurier
			if hadChap {
				// Ajout de l'équipement à l'inventaire du personnage
				inventory.AddEquipment(c, chapAvent)
			} else {
				// Message d'erreur si le joueur ne possède pas l'équipement
				fmt.Printf("\nTu ne possèdes pas : %s\n\n", chapAvent.Name)
			}
		case 2:
			// Équiper la tunique de l'aventurier
			if hadTun {
				inventory.AddEquipment(c, tunAvent)
			} else {
				fmt.Printf("\nTu ne possèdes pas : %s\n\n", tunAvent.Name)
			}
		case 3:
			// Équiper les bottes de l'aventurier
			if hadBot {
				inventory.AddEquipment(c, botAvent)
			} else {
				fmt.Printf("\nTu ne possèdes pas : %s\n\n", botAvent.Name)
			}
		case 4:

		default:
			fmt.Printf("\n❌ Il ne se passe rien... Choix invalide.\n")
		}
		if newEquipChoice == 4 {
			newEquipChoice = 0
			break
		}
	}
}

// CharacterCreation permet de créer un personnage en choisissant un pseudo et une classe
func CharacterCreation(c *structures.Character) {
	// Initialisation des variables
	var username string
	var valid bool

	// Boucle pour demander un pseudo valide
	for {
		fmt.Println("\nQuel est votre pseudo ?")
		fmt.Scan(&username)
		valid = true
		result := []rune(username)

		// Vérification que le pseudo n'est pas vide et ne contient que des lettres
		if len(result) == 0 {
			fmt.Println("❌ Le pseudo ne peut pas être vide.")
			valid = false
			continue
		}

		// Vérification que chaque caractère est une lettre
		for _, r := range result {
			if r < 65 || (r > 90 && r < 97) || r > 122 {
				fmt.Println("❌ Votre pseudo n'est pas correct, il ne contient que des lettres.")
				valid = false
				break
			}
		}

		// Si le pseudo est valide, on le formate (première lettre majuscule, le reste en minuscule)
		if valid {
			if result[0] >= 97 && result[0] <= 122 {
				result[0] = result[0] - ('a' - 'A')
			}
			for i := 1; i < len(result); i++ {
				if result[i] >= 65 && result[i] <= 90 {
					result[i] = result[i] + ('a' - 'A')
				}
			}

			// Assignation du pseudo au personnage
			username = string(result)
			c.Name = username
			// Sortie de la boucle
			break
		}
	}

	// Choix de la classe du personnage
	fmt.Println("\nPersonnage créé avec le nom :", c.Name)
	// Initialisation de la variable
	var class_choice int
	// Boucle pour demander un choix de classe valide
	for {
		fmt.Println("\nSuper", c.Name, ", quelle classe veux-tu choisir ?")
		fmt.Println("1 - Elfe : 80 PV Max")
		fmt.Println("2 - Nain : 120 PV Max")
		fmt.Println("3 - Humain : 100 PV Max")
		fmt.Scan(&class_choice)

		// Vérification que le choix est valide
		if class_choice >= 1 && class_choice <= 3 {
			// Choix valide, on sort de la boucle
			break
		}
		fmt.Println("❌ Choix invalide, essaye encore.")
	}

	// Assignation des caractéristiques en fonction de la classe choisie
	switch class_choice {
	case 1:
		// Classe Elfe (80 PV Max, 40 PV Actuels (on commence avec la moitier des PV max), 120 Mana Max)
		fmt.Println("\nTu as choisi la classe Elfe : agile, précis et en communion avec la nature.")
		c.MaxHp = 80
		c.ActualHp = 40
		c.ManaMax = 120
		c.ActualMana = 120
		c.Class = "Elfe"
	case 2:
		// Classe Nain (120 PV Max, 60 PV Actuels, 80 Mana Max)
		fmt.Println("\nTu as choisi la classe Nain : robuste, courageux et maître de la forge.")
		c.MaxHp = 120
		c.ActualHp = 60
		c.ManaMax = 80
		c.ActualMana = 80
		c.Class = "Nain"
	case 3:
		// Classe Humain (100 PV Max, 50 PV Actuels, 80 Mana Max)
		fmt.Println("\nTu as choisi la classe Humain : polyvalent, ingénieux et déterminé.")
		c.MaxHp = 100
		c.ActualHp = 50
		c.ManaMax = 80
		c.ActualMana = 80
		c.Class = "Humain"
	}
}

// Fonction pour ajouter de l'experience au personnage en fonction de la difficulté de l'ennemi vaincu
func AddExp(c *structures.Character, e *structures.Enemy) {
	c.ActualExp += e.GiveExp
}
