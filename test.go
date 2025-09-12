package main

import "fmt"

func main() {
	choice := 0
	fmt.Print("======== Faites votre choix :======== \n")
	fmt.Print("1 - Commencer la partie \n")
	fmt.Print("2 - Quitter \n")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Commencer la partie")
	case 2 : 


	case 5:
		fmt.Print("Ariverderci !")
	}
}
c