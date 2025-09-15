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
	/*Un équipement de tête
	Un équipement pour le torse
	Un équipement pour les pieds */

}

func InitCharacter(name string, class string, level int, maxhp int, actualhp int, inv []Object, maxInv int, money int, skill []Skill) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHp:     maxhp,
		ActualHp:  actualhp,
		Inventory: inv,
		MaxInv:    maxInv,
		Money:     money,
		SkillList: skill,
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
