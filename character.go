package character

import (
	"fmt"
	"inventory"
)

type Character struct {
	name      string
	class     string
	level     int
	max_pv    int
	actual_pv int
	inventory []inventory.objects
}

func InitCharacter() {
	personnage_c1 := Character{"Harold", "Elfe", 1, 100, []inventory.Objects{"potions", 3}}
}

func AccessInventory() {
	fmt.Printlf(personnage_c1.inventory)
}

func DisplayInfo() {
	fmt.Printlf(personnage_c1)
}

func characterCreation() {
}
