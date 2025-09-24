package main

import "fmt"

func main3() {
	cards := []int{15, 4, 12, 7, 9, 13, 8, 2}
	var card int
	fmt.Print("Quel est la carte rechercher ?")
	fmt.Print("\n")
	fmt.Scan(&card)
	for counter := 0; counter < len(cards); counter++ {
		if cards[counter] == card {
			fmt.Println("trouvé !")
		}
	}
}
