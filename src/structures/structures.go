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
}

func InitCharacter(name string, class string, level int, maxhp int, actualhp int, inv []Object) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHp:     maxhp,
		ActualHp:  actualhp,
		Inventory: inv,
	}
}
