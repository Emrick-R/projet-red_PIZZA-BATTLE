package inventory

import (
	"fmt"
	"projet-red_PIZZA-BATTLE/structures"
)

// AddInventory ajoute un objet à l'inventaire du personnage, en augmentant la quantité si l'objet existe déjà
func AddInventory(c *structures.Character, newObj structures.Object) {
	for i := range c.Inventory {
		if c.Inventory[i].Name == newObj.Name {
			c.Inventory[i].Quantity += newObj.Quantity
			return
		}
	}
	c.Inventory = append(c.Inventory, newObj)
}

// RemoveInventory retire un objet de l'inventaire du personnage, en diminuant la quantité
// ou en supprimant l'objet si la quantité atteint zéro
func RemoveInventory(c *structures.Character, obj structures.Object) {
	for i := range c.Inventory {
		if c.Inventory[i].Name == obj.Name {
			if c.Inventory[i].Quantity > obj.Quantity {
				c.Inventory[i].Quantity -= obj.Quantity
			} else {
				c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			}
			return
		}
	}
}

// AddEquipment équipe un objet d'armure au personnage, en remplaçant l'armure existante si nécessaire.
func AddEquipment(c *structures.Character, newObj structures.Object) {
	// Équipe l'armure en fonction de son type (tête, torse, jambes)
	switch newObj.Name {
	case "Chapeau de l'aventurier":
		// Vérifie si l'armure de tête est déjà équipée
		if c.Armor.Head.Name != "Chapeau de l'aventurier" {
			// Si l'armure de tête n'est pas équipée, l'équipe
			c.Armor.Head = &newObj
			fmt.Printf("\n✅ Tu équipes : Chapeau de l'aventurier\n\n")
			fmt.Printf("\nTu gagnes + 10 PV Max\n\n")
			fmt.Printf("❤️ %d/%d PV -> %d/%d PV\n\n", c.ActualHp, c.MaxHp, c.ActualHp, c.MaxHp+10)
			// Augmente les PV max du personnage de 10
			c.MaxHp += 10
			// Retire l'objet de l'inventaire
			RemoveInventory(c, newObj)
		} else {
			// Si l'armure de tête est déjà équipée, affiche un message et ne fait rien
			fmt.Printf("\n❌ Tu as déjà équipé : %s\n\n", newObj.Name)
		}
	case "Tunique de l'aventurier":
		// Vérifie si l'armure de torse est déjà équipée, si ce n'est pas le cas, l'équipe
		// et augmente les PV max de 25
		if c.Armor.Chest.Name != "Tunique de l'aventurier" {
			c.Armor.Chest = &newObj
			fmt.Printf("\n✅ Tu équipes : Tunique de l'aventurier\n\n")
			fmt.Printf("\nTu gagnes + 25 PV Max\n\n")
			fmt.Printf("❤️ %d/%d PV -> %d/%d PV\n\n", c.ActualHp, c.MaxHp, c.ActualHp, c.MaxHp+25)
			c.MaxHp += 25
			RemoveInventory(c, newObj)
		} else {
			fmt.Printf("\n❌ Tu as déjà équipé : %s\n\n", newObj.Name)
		}
	case "Bottes de l'aventurier":
		// Vérifie si l'armure de jambes est déjà équipée, si ce n'est pas le cas, l'équipe
		// et augmente les PV max de 15
		if c.Armor.Legs.Name != "Bottes de l'aventurier" {
			c.Armor.Legs = &newObj
			fmt.Printf("\n✅ Tu équipes : Bottes de l'aventurier\n\n")
			fmt.Printf("\nTu gagnes + 15 PV Max\n\n")
			fmt.Printf("❤️ %d/%d PV -> %d/%d PV\n\n", c.ActualHp, c.MaxHp, c.ActualHp, c.MaxHp+15)
			c.MaxHp += 15
			RemoveInventory(c, newObj)
		} else {
			fmt.Printf("\n❌ Tu as déjà équipé : %s\n\n", newObj.Name)
		}
	}
}

// Fonction pour ajouter de l'argent au personnage en fonction de la difficulté de l'ennemi vaincu
func AddMoney(c *structures.Character, e *structures.Enemy) {
	c.Money += e.GiveMoney
}

// CheckMaxInventory vérifie si l'inventaire du personnage a de la place pour ajouter de nouveaux objets
// True = Il y a de la place, False = Inventaire plein
func CheckMaxInventory(c *structures.Character) bool {
	counter := 0
	for i := range c.Inventory {
		counter = counter + c.Inventory[i].Quantity
	}
	return counter < c.MaxInv
}

// UpgradeInventorySlot augmente la capacité maximale de l'inventaire du personnage de 10 emplacements
func UpgradeInventorySlot(c *structures.Character) {
	c.MaxInv = c.MaxInv + 10
}
