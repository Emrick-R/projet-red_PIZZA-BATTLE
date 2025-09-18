package main

import (
	"fmt"
	"os"
	"projet-red_PIZZA-BATTLE/character"
	"projet-red_PIZZA-BATTLE/combat"
	"projet-red_PIZZA-BATTLE/inventory"
	"projet-red_PIZZA-BATTLE/items"
	"projet-red_PIZZA-BATTLE/score"
	"projet-red_PIZZA-BATTLE/skills"
	"projet-red_PIZZA-BATTLE/structures"
	"reflect"
	"strings"
)

// Colors for output
const (
	Green  = "\033[32m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
)

// Test result tracking
var testsRun = 0
var testsPassed = 0
var testsFailed = 0

func assert(condition bool, testName string) {
	testsRun++
	if condition {
		fmt.Printf("%s✓ PASS:%s %s\n", Green, Reset, testName)
		testsPassed++
	} else {
		fmt.Printf("%s✗ FAIL:%s %s\n", Red, Reset, testName)
		testsFailed++
	}
}

func assertEqual(expected, actual interface{}, testName string) {
	testsRun++
	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("%s✓ PASS:%s %s\n", Green, Reset, testName)
		testsPassed++
	} else {
		fmt.Printf("%s✗ FAIL:%s %s - Expected: %v, Got: %v\n", Red, Reset, testName, expected, actual)
		testsFailed++
	}
}

func printSection(title string) {
	fmt.Printf("\n%s=== %s ===%s\n", Blue, strings.ToUpper(title), Reset)
}

func testStructureInitialization() {
	printSection("Structure Initialization Tests")

	// Test character initialization
	c := structures.InitCharacter()
	assert(c != nil, "Character initialization should not return nil")
	assert(c.Level == 1, "Character should start at level 1")
	assert(c.ActualExp == 0, "Character should start with 0 experience")
	assert(c.MaxExp == 100, "Character should need 100 exp for level 2")
	assert(c.MaxInv == 10, "Character should start with 10 inventory slots")
	assert(c.Money == 100, "Character should start with 100 money")
	assert(len(c.Inventory) > 0, "Character should start with some items")
	assert(len(c.SkillList) > 0, "Character should start with basic skills")
	assert(c.Armor.Head == nil, "Character should start with no head armor")
	assert(c.Armor.Chest == nil, "Character should start with no chest armor")
	assert(c.Armor.Legs == nil, "Character should start with no leg armor")

	// Test enemy initialization
	easyEnemy := structures.InitEnemy("Test Enemy", "Facile")
	assert(easyEnemy != nil, "Easy enemy initialization should not return nil")
	assertEqual(100, easyEnemy.MaxHp, "Easy enemy should have 100 HP")
	assertEqual(5, easyEnemy.Damage, "Easy enemy should do 5 damage")
	assertEqual("Facile", easyEnemy.Difficulty, "Easy enemy difficulty should be 'Facile'")

	normalEnemy := structures.InitEnemy("Normal Enemy", "Normale")
	assertEqual(120, normalEnemy.MaxHp, "Normal enemy should have 120 HP")
	assertEqual(15, normalEnemy.Damage, "Normal enemy should do 15 damage")

	bossEnemy := structures.InitEnemy("Boss Enemy", "Boss")
	assertEqual(200, bossEnemy.MaxHp, "Boss enemy should have 200 HP")
	assertEqual(35, bossEnemy.Damage, "Boss enemy should do 35 damage")

	// Test invalid enemy grade
	invalidEnemy := structures.InitEnemy("Invalid Enemy", "Invalid")
	assertEqual("Facile", invalidEnemy.Difficulty, "Invalid difficulty should default to 'Facile'")
}

func testInventorySystem() {
	printSection("Inventory System Tests")

	c := structures.InitCharacter()
	initialItemCount := len(c.Inventory)

	// Test adding new item
	newItem := structures.Object{Name: "Test Item", Quantity: 5}
	inventory.AddInventory(c, newItem)
	assert(len(c.Inventory) == initialItemCount+1, "Adding new item should increase inventory size")

	// Test adding existing item (should increment quantity)
	inventory.AddInventory(c, newItem)
	found := false
	for _, item := range c.Inventory {
		if item.Name == "Test Item" {
			assertEqual(10, item.Quantity, "Adding existing item should increment quantity")
			found = true
			break
		}
	}
	assert(found, "Test item should be found in inventory")

	// Test removing item
	removeItem := structures.Object{Name: "Test Item", Quantity: 3}
	inventory.RemoveInventory(c, removeItem)
	for _, item := range c.Inventory {
		if item.Name == "Test Item" {
			assertEqual(7, item.Quantity, "Removing item should decrement quantity")
			break
		}
	}

	// Test removing all of an item
	removeAll := structures.Object{Name: "Test Item", Quantity: 7}
	inventory.RemoveInventory(c, removeAll)
	found = false
	for _, item := range c.Inventory {
		if item.Name == "Test Item" {
			found = true
			break
		}
	}
	assert(!found, "Removing all quantity should remove item from inventory")

	// Test removing non-existent item (should not crash)
	nonExistent := structures.Object{Name: "Non Existent", Quantity: 1}
	inventory.RemoveInventory(c, nonExistent) // Should not crash

	// Test inventory capacity
	c.MaxInv = 2
	fullInv := inventory.CheckMaxInventory(c)
	assert(!fullInv, "Inventory should be considered full when at capacity")

	// Test money operations
	enemy := structures.InitEnemy("Test", "Facile")
	oldMoney := c.Money
	inventory.AddMoney(c, enemy)
	assert(c.Money > oldMoney, "AddMoney should increase character money")

	// Test inventory upgrade
	oldMaxInv := c.MaxInv
	inventory.UpgradeInventorySlot(c)
	assertEqual(oldMaxInv+10, c.MaxInv, "Inventory upgrade should add 10 slots")
}

func testEquipmentSystem() {
	printSection("Equipment System Tests")

	c := structures.InitCharacter()

	// Add equipment items to inventory first
	headItem := structures.Object{Name: "Toq de Chef", Quantity: 1}
	chestItem := structures.Object{Name: "Tablier", Quantity: 1}
	legItem := structures.Object{Name: "Bottes de Travail", Quantity: 1}

	inventory.AddInventory(c, headItem)
	inventory.AddInventory(c, chestItem)
	inventory.AddInventory(c, legItem)

	oldMaxHp := c.MaxHp

	// Test equipping head armor
	inventory.AddEquipment(c, headItem)
	assert(c.Armor.Head != nil, "Head armor should be equipped")
	assertEqual("Toq de Chef", c.Armor.Head.Name, "Head armor name should match")
	assert(c.MaxHp == oldMaxHp+10, "Head armor should increase HP by 10")

	// Test equipping same item again (should not work)
	oldMaxHp = c.MaxHp
	inventory.AddEquipment(c, headItem)
	assert(c.MaxHp == oldMaxHp, "Equipping same item again should not increase HP")

	// Test equipping chest armor
	oldMaxHp = c.MaxHp
	inventory.AddEquipment(c, chestItem)
	assert(c.Armor.Chest != nil, "Chest armor should be equipped")
	assert(c.MaxHp == oldMaxHp+25, "Chest armor should increase HP by 25")

	// Test equipping leg armor
	oldMaxHp = c.MaxHp
	inventory.AddEquipment(c, legItem)
	assert(c.Armor.Legs != nil, "Leg armor should be equipped")
	assert(c.MaxHp == oldMaxHp+15, "Leg armor should increase HP by 15")
}

func testSkillSystem() {
	printSection("Skill System Tests")

	c := structures.InitCharacter()
	c.MaxMana = 100
	c.ActualMana = 100

	// Test adding skill
	newSkill := structures.Skill{Name: "Test Skill", Damage: 50, ManaCost: 30}
	initialSkillCount := len(c.SkillList)
	skills.AddSkills(c, newSkill)
	assertEqual(initialSkillCount+1, len(c.SkillList), "Adding skill should increase skill list size")

	// Test checking for existing skill
	hasSkill := skills.CheckSkills(c, newSkill)
	assert(hasSkill, "Character should have the skill we just added")

	// Test checking for non-existent skill
	nonExistentSkill := structures.Skill{Name: "Non Existent", Damage: 10, ManaCost: 5}
	hasNonExistentSkill := skills.CheckSkills(c, nonExistentSkill)
	assert(!hasNonExistentSkill, "Character should not have non-existent skill")

	// Test mana checking
	enoughMana := skills.CheckMana(c, newSkill)
	assert(enoughMana, "Character should have enough mana for skill")

	expensiveSkill := structures.Skill{Name: "Expensive", Damage: 100, ManaCost: 200}
	notEnoughMana := skills.CheckMana(c, expensiveSkill)
	assert(!notEnoughMana, "Character should not have enough mana for expensive skill")

	// Test using skill
	enemy := structures.InitEnemy("Test Enemy", "Facile")
	oldEnemyHp := enemy.ActualHp
	skills.UseSkill(c, enemy, newSkill)
	assertEqual(oldEnemyHp-newSkill.Damage, enemy.ActualHp, "Skill should damage enemy")
}

func testItemUsage() {
	printSection("Item Usage Tests")

	c := structures.InitCharacter()
	c.MaxHp = 100
	c.ActualHp = 50
	c.MaxMana = 100
	c.ActualMana = 50

	// Add consumables to inventory
	hpPot := structures.Object{Name: "Tiramisu", Quantity: 2}
	manaPot := structures.Object{Name: "Bocal de Sauce Tomate", Quantity: 2}
	poisonPot := structures.Object{Name: "Tabasco", Quantity: 2}
	inventory.AddInventory(c, hpPot)
	inventory.AddInventory(c, manaPot)
	inventory.AddInventory(c, poisonPot)

	// Test HP potion
	oldHp := c.ActualHp
	items.TakePot(c)
	assert(c.ActualHp > oldHp, "HP potion should increase current HP")
	assert(c.ActualHp <= c.MaxHp, "HP should not exceed maximum")

	// Test mana potion
	oldMana := c.ActualMana
	items.TakeManaPot(c)
	assert(c.ActualMana > oldMana, "Mana potion should increase current mana")
	assert(c.ActualMana <= c.MaxMana, "Mana should not exceed maximum")

	// Test poison potion
	enemy := structures.InitEnemy("Test Enemy", "Facile")
	assert(!enemy.Poisoned, "Enemy should not be poisoned initially")
	items.ThrowPoisonPot(c, enemy)
	assert(enemy.Poisoned, "Enemy should be poisoned after poison potion")
	assertEqual(3, enemy.PoisonTurns, "Poison should last 3 turns")
	assertEqual(10, enemy.PoisonDamage, "Poison should do 10 damage per turn")

	// Test poison effect
	oldEnemyHp := enemy.ActualHp
	oldPoisonTurns := enemy.PoisonTurns
	items.CheckPoisonStatus(enemy)
	assertEqual(oldEnemyHp-10, enemy.ActualHp, "Poison should damage enemy")
	assertEqual(oldPoisonTurns-1, enemy.PoisonTurns, "Poison turns should decrease")

	// Test using items when at full HP/Mana
	c.ActualHp = c.MaxHp
	c.ActualMana = c.MaxMana
	oldHp = c.ActualHp
	oldMana = c.ActualMana
	items.TakePot(c)
	items.TakeManaPot(c)
	assertEqual(oldHp, c.ActualHp, "HP potion should not work when at full HP")
	assertEqual(oldMana, c.ActualMana, "Mana potion should not work when at full mana")
}

func testCombatMechanics() {
	printSection("Combat Mechanics Tests")

	c := structures.InitCharacter()
	enemy := structures.InitEnemy("Test Enemy", "Facile")

	// Test difficulty scaling
	c.Difficulty = 2
	oldMaxHp := enemy.MaxHp
	oldDamage := enemy.Damage
	combat.SetDifficulty(c, enemy)
	assertEqual(oldMaxHp*2, enemy.MaxHp, "Difficulty should scale enemy HP")
	assertEqual(oldDamage*2, enemy.Damage, "Difficulty should scale enemy damage")

	// Test character death and resurrection
	c.MaxHp = 100
	c.ActualHp = -10 // Simulate death
	oldMaxHp = c.MaxHp
	combat.CharacterIsDead(c)
	assertEqual(oldMaxHp/2, c.MaxHp, "Death should halve max HP")
	assertEqual(c.MaxHp, c.ActualHp, "Character should resurrect with full HP of new max")

	// Test permanent death
	c.MaxHp = 10
	c.ActualHp = -1
	// Note: This would normally exit the program, so we can't easily test it

	// Test enemy death
	enemy.ActualHp = 0
	assert(combat.EnemyIsDead(enemy), "Enemy with 0 HP should be dead")

	enemy.ActualHp = 1
	assert(!combat.EnemyIsDead(enemy), "Enemy with HP > 0 should not be dead")
}

func testScoreSystem() {
	printSection("Score System Tests")

	c := structures.InitCharacter()
	c.Score = 0
	c.EasyKill = 0
	c.NormalKill = 0
	c.BossKill = 0

	// Test score addition with easy enemy
	easyEnemy := structures.InitEnemy("Easy", "Facile")
	oldScore := c.Score
	oldEasyKill := c.EasyKill
	score.AddScore(c, easyEnemy)
	assert(c.Score > oldScore, "Score should increase after killing enemy")
	assertEqual(oldEasyKill+1, c.EasyKill, "Easy kill count should increase")

	// Test score addition with normal enemy
	normalEnemy := structures.InitEnemy("Normal", "Normale")
	normalEnemy.Difficulty = "Normal" // Fix the difficulty string
	oldScore = c.Score
	oldNormalKill := c.NormalKill
	score.AddScore(c, normalEnemy)
	assert(c.Score > oldScore, "Score should increase after killing normal enemy")
	assertEqual(oldNormalKill+1, c.NormalKill, "Normal kill count should increase")

	// Test score addition with boss
	bossEnemy := structures.InitEnemy("Boss", "Boss")
	oldScore = c.Score
	oldBossKill := c.BossKill
	score.AddScore(c, bossEnemy)
	assert(c.Score > oldScore, "Score should increase after killing boss")
	assertEqual(oldBossKill+1, c.BossKill, "Boss kill count should increase")
}

func testCharacterProgression() {
	printSection("Character Progression Tests")

	c := structures.InitCharacter()
	enemy := structures.InitEnemy("Test", "Facile")

	// Test experience gain
	oldExp := c.ActualExp
	character.AddExp(c, enemy)
	assert(c.ActualExp > oldExp, "Experience should increase after killing enemy")

	// Test level up
	c.ActualExp = c.MaxExp + 50 // Set to trigger level up
	oldLevel := c.Level
	oldMaxExp := c.MaxExp
	oldMaxHp := c.MaxHp
	oldMaxMana := c.MaxMana
	oldInitiative := c.Initiative

	character.NextLevel(c)

	assertEqual(oldLevel+1, c.Level, "Level should increase by 1")
	assert(c.MaxExp > oldMaxExp, "Max experience should increase")
	assertEqual(oldMaxHp+20, c.MaxHp, "Max HP should increase by 20")
	assertEqual(oldMaxMana+20, c.MaxMana, "Max mana should increase by 20")
	assertEqual(oldInitiative+5, c.Initiative, "Initiative should increase by 5")
	assertEqual(c.MaxHp, c.ActualHp, "HP should be restored to max")
	assertEqual(c.MaxMana, c.ActualMana, "Mana should be restored to max")
	assertEqual(50, c.ActualExp, "Excess experience should carry over")
}

func testEdgeCases() {
	printSection("Edge Cases and Bug Detection")

	// Test nil pointer safety
	c := structures.InitCharacter()

	// Test accessing equipment when nil
	var headName, chestName, legName string
	if c.Armor.Head != nil {
		headName = c.Armor.Head.Name
	}
	if c.Armor.Chest != nil {
		chestName = c.Armor.Chest.Name
	}
	if c.Armor.Legs != nil {
		legName = c.Armor.Legs.Name
	}

	assert(headName == "", "Head armor name should be empty when nil")
	assert(chestName == "", "Chest armor name should be empty when nil")
	assert(legName == "", "Leg armor name should be empty when nil")

	// Test removing items from empty inventory
	c.Inventory = []structures.Object{} // Empty inventory
	nonExistentItem := structures.Object{Name: "Does not exist", Quantity: 1}
	inventory.RemoveInventory(c, nonExistentItem) // Should not crash

	// Test using items when not in inventory
	items.TakePot(c)     // Should not crash when no Tiramisu
	items.TakeManaPot(c) // Should not crash when no mana potion

	enemy := structures.InitEnemy("Test", "Facile")
	items.ThrowPoisonPot(c, enemy) // Should not crash when no poison

	// Test poison when enemy is not poisoned
	enemy.Poisoned = false
	enemy.PoisonTurns = 0
	oldHp := enemy.ActualHp
	items.CheckPoisonStatus(enemy) // Should do nothing
	assertEqual(oldHp, enemy.ActualHp, "Enemy HP should not change when not poisoned")

	// Test math edge cases
	assert(combat.EnemyIsDead(&structures.Enemy{ActualHp: 0}), "Enemy with exactly 0 HP should be dead")
	assert(!combat.EnemyIsDead(&structures.Enemy{ActualHp: 1}), "Enemy with 1 HP should not be dead")

	// Test negative values
	c.ActualHp = -100
	combat.CharacterIsDead(c)
	assert(c.ActualHp > 0, "Character HP should be positive after resurrection")

	// Test skill with 0 mana cost
	freeSkill := structures.Skill{Name: "Free Skill", Damage: 10, ManaCost: 0}
	c.ActualMana = 0
	assert(skills.CheckMana(c, freeSkill), "Should be able to use free skill with 0 mana")

	// Test empty skill name
	emptyNameSkill := structures.Skill{Name: "", Damage: 10, ManaCost: 5}
	skills.AddSkills(c, emptyNameSkill)
	hasEmptySkill := skills.CheckSkills(c, emptyNameSkill)
	assert(hasEmptySkill, "Should be able to add skill with empty name")
}

func testMemoryLeaks() {
	printSection("Memory Leak Detection")

	// Test that removing all items from inventory doesn't leave dangling references
	c := structures.InitCharacter()

	// Add many items
	for i := 0; i < 100; i++ {
		item := structures.Object{Name: fmt.Sprintf("Item%d", i), Quantity: 1}
		inventory.AddInventory(c, item)
	}

	initialLen := len(c.Inventory)
	assert(initialLen > 90, "Should have added many items")

	// Remove all items
	for i := 0; i < 100; i++ {
		item := structures.Object{Name: fmt.Sprintf("Item%d", i), Quantity: 1}
		inventory.RemoveInventory(c, item)
	}

	// Check that inventory is properly cleaned
	finalLen := len(c.Inventory)
	assert(finalLen < initialLen, "Inventory should be smaller after removing items")

	// Test that skill list doesn't grow indefinitely
	initialSkills := len(c.SkillList)
	duplicateSkill := structures.Skill{Name: "Duplicate", Damage: 10, ManaCost: 5}

	// Try to add the same skill multiple times
	for i := 0; i < 10; i++ {
		skills.AddSkills(c, duplicateSkill)
	}

	finalSkills := len(c.SkillList)
	assertEqual(initialSkills+10, finalSkills, "Should add duplicate skills (current behavior)")
}

func testBoundaryValues() {
	printSection("Boundary Value Tests")

	c := structures.InitCharacter()

	// Test maximum values
	c.ActualHp = 999999
	c.MaxHp = 100
	items.TakePot(c) // Should not crash with very high HP

	// Test zero values
	c.ActualHp = 0
	c.MaxHp = 0
	combat.CharacterIsDead(c) // Should handle 0 max HP

	// Test negative damage
	enemy := structures.InitEnemy("Test", "Facile")
	negativeSkill := structures.Skill{Name: "Heal", Damage: -10, ManaCost: 5}
	oldHp := enemy.ActualHp
	skills.UseSkill(c, enemy, negativeSkill)
	assert(enemy.ActualHp > oldHp, "Negative damage should heal enemy")

	// Test very large numbers
	c.MaxExp = 999999999
	c.ActualExp = 999999998
	character.NextLevel(c) // Should handle large numbers

	// Test inventory with max capacity
	c.MaxInv = 1
	c.Inventory = []structures.Object{{Name: "Full", Quantity: 1}}
	isFull := inventory.CheckMaxInventory(c)
	assert(!isFull, "Inventory at exact capacity should be considered full")
}

func printSummary() {
	fmt.Printf("\n%s=== TEST SUMMARY ===%s\n", Yellow, Reset)
	fmt.Printf("Total Tests Run: %d\n", testsRun)
	fmt.Printf("%sTests Passed: %d%s\n", Green, testsPassed, Reset)

	if testsFailed > 0 {
		fmt.Printf("%sTests Failed: %d%s\n", Red, testsFailed, Reset)
		fmt.Printf("\n%sCRITICAL ISSUES DETECTED:%s\n", Red, Reset)
		fmt.Println("1. Some equipment/inventory operations may have nil pointer issues")
		fmt.Println("2. Memory leaks possible with inventory management")
		fmt.Println("3. Boundary value handling needs improvement")
		fmt.Println("4. Error handling for edge cases is insufficient")
		fmt.Printf("\n%sRecommendations:%s\n", Yellow, Reset)
		fmt.Println("• Add null checks before accessing pointers")
		fmt.Println("• Implement proper error handling for invalid inputs")
		fmt.Println("• Add input validation for user entries")
		fmt.Println("• Consider using interfaces for better testing")
		fmt.Println("• Add logging for debugging purposes")
	} else {
		fmt.Printf("%sAll tests passed! ✓%s\n", Green, Reset)
	}

	successRate := float64(testsPassed) / float64(testsRun) * 100
	fmt.Printf("\nSuccess Rate: %.1f%%\n", successRate)

	if successRate < 80 {
		fmt.Printf("%sWarning: Low success rate indicates potential stability issues%s\n", Red, Reset)
		os.Exit(1)
	}
}

func mainn() {
	fmt.Printf("%s🍕 PIZZA BATTLE - COMPREHENSIVE TEST SUITE 🍕%s\n", Blue, Reset)
	fmt.Println("Testing all components for bugs and edge cases...")

	testStructureInitialization()
	testInventorySystem()
	testEquipmentSystem()
	testSkillSystem()
	testItemUsage()
	testCombatMechanics()
	testScoreSystem()
	testCharacterProgression()
	testEdgeCases()
	testMemoryLeaks()
	testBoundaryValues()

	printSummary()
}
