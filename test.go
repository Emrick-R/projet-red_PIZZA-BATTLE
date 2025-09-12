package main

import "fmt"

// Structure Object (équivalent à inventory.go)
type Object struct {
	Name     string
	Quantity int
}

// Structure Character (équivalent à character.go)
type Character struct {
	Name      string
	Class     string
	Level     int
	MaxPv     int
	ActualPv  int
	Inventory []Object
}

// Fonction d'initialisation du personnage
func InitCharacter(name string, class string, level int, maxPv int, actualPv int, inv []Object) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxPv:     maxPv,
		ActualPv:  actualPv,
		Inventory: inv,
	}
}

// Méthodes du personnage
func (c *Character) DisplayInfo() {
	fmt.Printf("Nom: %s, Classe: %s, Niveau: %d, PV: %d/%d\n",
		c.Name, c.Class, c.Level, c.ActualPv, c.MaxPv)
}

func (c *Character) AccessInventory() {
	fmt.Println("Inventaire:")
	for _, item := range c.Inventory {
		fmt.Printf("- %s x%d\n", item.Name, item.Quantity)
	}
}

// Fonction TakePot (équivalent à items.go)
func TakePot(c *Character) {
	for i, obj := range c.Inventory {
		if obj.Name == "Potion" && obj.Quantity > 0 {
			c.ActualPv += 50
			if c.ActualPv > c.MaxPv {
				c.ActualPv = c.MaxPv
			}
			c.Inventory[i].Quantity--

			if c.Inventory[i].Quantity == 0 {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			fmt.Println("Potion utilisée ! +50 PV")
			return
		}
	}
	fmt.Println("Il n'y a pas de potions dans l'inventaire")
}

// Fonction principale
func main() {
	// Création de l'inventaire
	inv := []Object{
		{"Potion", 3},
	}

	// Création du perso
	c1 := InitCharacter("Harold", "Elfe", 1, 100, 50, inv)

	// Affichage
	fmt.Println("=== État initial ===")
	c1.DisplayInfo()
	c1.AccessInventory()

	fmt.Println("\n=== Utilisation d'une potion ===")
	TakePot(&c1)

	fmt.Println("\n=== État après potion ===")
	c1.DisplayInfo()
	c1.AccessInventory()

	choice := 0
	fmt.Print("======== Faites votre choix :======== \n")
	fmt.Print("1 - Commencer la partie \n")
	fmt.Print("2 - Quitter \n")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Commencer la partie")
	case 2:

	case 5:
		fmt.Print("Ariverderci !")
	}
}
