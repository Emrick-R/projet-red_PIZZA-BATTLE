package character

type character struct {
	name      string
	class     string
	level     int
	max_pv    int
	actual_pv int
	inventory []objects
}

func InitCharacter() {
	personnage_c1 := character{"Harold", "class", 1, 100, []inventory{"3 potions"}}
}

func AccessInventory() {
	fmt.printlf(character.inventory)
}

func DisplayInfo() {
	fmt.printlf(character)
}

func characterCreation() {
}
