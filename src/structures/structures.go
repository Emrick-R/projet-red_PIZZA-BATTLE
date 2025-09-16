package structures

// Object représente un objet dans l'inventaire ou l'équipement
type Object struct {
	Name     string
	Quantity int
}

// Character représente un personnage joueur avec ses attributs et son inventaire
type Character struct {
	Name       string
	Class      string
	Level      int
	MaxHp      int
	ActualHp   int
	Inventory  []Object
	MaxInv     int
	Money      int
	SkillList  []Skill
	Armor      Equipment
	Score      int
	Initiative int
	ActualExp  int
	MaxExp     int
	ActualMana int
	ManaMax    int
}

// Enemy représente un ennemi avec ses attributs
type Enemy struct {
	Name         string
	MaxHp        int
	ActualHp     int
	Damage       int
	Difficulty   string //3 types : Facile= score 5, Normale= score 10, Boss= score 20
	Initiative   int
	PowerCount   int
	GiveScore    int
	GiveMoney    int
	GiveExp      int
	Poisoned     bool
	PoisonTurns  int
	PoisonDamage int
}

// Skill représente une compétence que le personnage peut utiliser en combat
type Skill struct {
	Name     string
	Damage   int
	ManaCost int
}

// Equipment représente l'équipement porté par le personnage
type Equipment struct {
	Head  *Object //Pointeur unique vide de structure Objects
	Chest *Object //Un équipement pour le torse
	Legs  *Object //Un équipement pour les pieds

}

// InitCharacter initialise un personnage avec des valeurs par défaut sans nom ni classe (définis par CharacterCreation)
func InitCharacter() *Character {
	return &Character{
		Level:      1,
		ActualExp:  0,
		MaxExp:     100,
		Initiative: 100,
		// Initialisation avec 3 potions
		Inventory: []Object{
			{Name: "Potion de Vie", Quantity: 3},
		},
		ActualHp: 100,
		MaxInv:   10,
		Money:    100,
		// Initialisation de la compétance de base
		SkillList: []Skill{
			{Name: "Coup de poing", Damage: 10, ManaCost: 0},
		},
		// Initialisation de l'armure (rien d'éauipé)
		Armor: Equipment{
			Head:  &Object{Name: ""},
			Chest: &Object{Name: ""},
			Legs:  &Object{Name: ""},
		},
		//Score de fin de partie
		Score: 0,
	}
}

// InitEnemy initialise un Ennemi avec des valeurs donnée selon la difficulté donnée et les valeurs données
func InitEnemy(name string, grade string) *Enemy {
	switch grade {
	// si La difficulté donnée est "Facile"
	case "Facile":

		return &Enemy{
			Name:       name,
			MaxHp:      100,
			ActualHp:   100,
			Damage:     5,
			Difficulty: "Facile",
			Initiative: 100,
			PowerCount: 0,
			GiveScore:  5,
			GiveMoney:  5,
			GiveExp:    100,
		}
		// si La difficulté donnée est "Normale"
	case "Normale":

		return &Enemy{
			Name:       name,
			MaxHp:      120,
			ActualHp:   120,
			Damage:     15,
			Difficulty: "Normal",
			Initiative: 120,
			PowerCount: 0,
			GiveScore:  10,
			GiveMoney:  10,
			GiveExp:    150,
		}
		// si La difficulté donnée est "Boss"
	case "Boss":

		return &Enemy{
			Name:       name,
			MaxHp:      200,
			ActualHp:   200,
			Damage:     25,
			Difficulty: "Boss",
			Initiative: 150,
			PowerCount: 0,
			GiveScore:  5,
			GiveMoney:  5,
			GiveExp:    300,
		}
	}
	// si La difficulté donnée est rien, initialise quand même un ennemi "Facile"
	return &Enemy{
		Name:       name,
		MaxHp:      100,
		ActualHp:   100,
		Damage:     5,
		Difficulty: "Facile",
		Initiative: 100,
		PowerCount: 0,
		GiveScore:  5,
		GiveMoney:  5,
		GiveExp:    100,
	}
}

//Initialise une compétence utilisable dans la liste des compétences
func InitSkill(name string, damage int) *Skill {
	return &Skill{
		Name:   name,
		Damage: damage,
	}

}
