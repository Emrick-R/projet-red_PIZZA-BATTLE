package structures

type Object struct {
	Name     string
	Quantity int
}

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
	ExpActual  int
	ExpMax     int
	ActualMana int
	ManaMax    int
}
type Enemy struct {
	Name       string
	MaxHp      int
	ActualHp   int
	Damage     int
	Difficulty string //3 types : Facile= score 5, Normal= score 10, Boss= score 20
	Initiative int
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

func InitCharacter(level int, inv []Object, maxInv int, money int, skill []Skill, initiative int, ExpActual, ExpMax int) *Character {
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
		Score:      int(0),
		Initiative: initiative,
		ExpActual:  0,
		ExpMax:     100,
	}
}

func InitEnemy(name string, maxhp int, actualhp int, damage int, grade string, initiative int) *Enemy {
	return &Enemy{
		Name:       name,
		MaxHp:      maxhp,
		ActualHp:   actualhp,
		Damage:     damage,
		Difficulty: grade,
		Initiative: initiative,
		PowerCount: 0,
	}
}

func InitSkill(name string, damage int) *Skill {
	return &Skill{
		Name:   name,
		Damage: damage,
	}

}
