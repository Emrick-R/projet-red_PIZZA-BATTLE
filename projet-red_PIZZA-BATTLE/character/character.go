package character

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/inventory"
)

type Character struct {
	name      string
	class     string
	level     int
	max_pv    int
	actual_pv int
	inventory []inventory.Objects
}

func InitCharacter(name string, class string, level int, max_pv int, actual_pv int, name_object string, quantity int) {
	personnage_c1 := Character{name, class, level, max_pv, []inventory.Object{name_object, quantity}}
	return personnage_c1
}

func AccessInventory() {
	fmt.Println(personnage_c1.inventory)
}

func DisplayInfo() {
	fmt.Println(personnage_c1)
}

func CharacterCreation() {
}
