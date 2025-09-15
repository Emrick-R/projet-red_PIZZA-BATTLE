package main

import "fmt"

func CharacterCreation() {
	var username string
	var valid bool

	for {
		fmt.Println("Quel est votre pseudo ?")
		fmt.Scan(&username)

		valid = true
		result := []rune(username)

		if len(result) == 0 {
			fmt.Println("Le pseudo ne peut pas être vide.")
			valid = false
			continue
		}

		for _, r := range result {
			if r < 65 || (r > 90 && r < 97) || r > 122 {
				fmt.Println("Votre pseudo n'est pas correct, il ne contient pas que des lettres.")
				valid = false
				break
			}
		}

		if valid {
			// Première lettre en majuscule
			if result[0] >= 97 && result[0] <= 122 {
				result[0] = result[0] - ('a' - 'A')
			}
			// Reste des lettres en minuscules
			for i := 1; i < len(result); i++ {
				if result[i] >= 65 && result[i] <= 90 {
					result[i] = result[i] + ('a' - 'A')
				}
			}

			username = string(result)
			c.Name = username
			break
		}
	}

	fmt.Println("Personnage créé avec le nom :", c.Name)
}
