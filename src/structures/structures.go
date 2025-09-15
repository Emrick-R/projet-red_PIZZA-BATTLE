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
}
type Enemy struct {
	Name     string
	MaxHp    int
	ActualHp int
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

func InitCharacter(level int, inv []Object, maxInv int, money int, skill []Skill, Armor Equipment) *Character {
	return &Character{
		Level:     level,
		Inventory: inv,
		MaxInv:    maxInv,
		Money:     money,
		SkillList: skill,
		Armor:     Armor,
	}
}
func InitEnemy(name string, maxhp int, actualhp int) *Enemy {
	return &Enemy{
		Name:     name,
		MaxHp:    maxhp,
		ActualHp: actualhp,
	}
}

func InitSkill(name string, damage int) *Skill {
	return &Skill{
		Name:   name,
		Damage: damage,
	}

}
