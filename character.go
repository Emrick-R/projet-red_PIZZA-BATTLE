package character

import "inventory"

type character struct {
	name      string
	class     string
	level     int
	max_pv    int
	actual_pv int
	inventory []inventory.objects
}

func InitCharacter() {
	personnage_c1 := character{"Harold", "Elfe", 1, 100, []inventory.Objects{"potions", 3}}
}

func AccessInventory() {
	fmt.printlf(character.inventory)
}

func DisplayInfo() {
	fmt.printlf(character)
}

func characterCreation() {
}
