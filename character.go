package character

type character struct {
	name      string
	class     string
	level     int
	max_pv    int
	actual_pv int
	inventory []objects
}

func AccessInventory() {
	fmt.printlf(character.inventory)
}
