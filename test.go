package main

import (
	"fmt"
	"time"
)

func main() {
	// Nombre de répétitions
	max := 20

	for i := 0; i < max; i++ {
		fmt.Println("brr")
		time.Sleep(200 * time.Millisecond) // 0,2 seconde entre chaque "brr"
	}

	fmt.Println("Terminé !")
}
