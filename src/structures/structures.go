package structures

type Object struct {
	Name     string
	Quantity int
}

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHp     int
	ActualHp  int
	Inventory []Object
	MaxInv    int
	Money     int
	SkillList []Skill
	Armor     Equipment
	Score     int
}
type Enemy struct {
	Name       string
	MaxHp      int
	ActualHp   int
	Damage     int
	Difficulty string //3 grades : Facile= score 5, Normal= score 10, Boss= score 20
	PowerCount int
}

type Skill struct {
	Name   string
	Damage int
}

type Equipment struct {
	Head  *Object //Pointeur unique vide de structure Objects
	Chest *Object //Un équipement pour le torse
	Legs  *Object //Un équipement pour les pieds

}

func InitCharacter(level int, inv []Object, maxInv int, money int, skill []Skill) *Character {
	return &Character{
		Level:     level,
		Inventory: inv,
		MaxInv:    maxInv,
		Money:     money,
		SkillList: skill,
		Armor: Equipment{
			Head:  &Object{Name: ""},
			Chest: &Object{Name: ""},
			Legs:  &Object{Name: ""},
		},
		Score: int(0),
	}
}

func InitEnemy(name string, maxhp int, actualhp int, damage int, grade string) *Enemy {
	return &Enemy{
		Name:       name,
		MaxHp:      maxhp,
		ActualHp:   actualhp,
		Damage:     damage,
		Difficulty: grade,
		PowerCount: 0,
	}
}

func InitSkill(name string, damage int) *Skill {
	return &Skill{
		Name:   name,
		Damage: damage,
	}

}
