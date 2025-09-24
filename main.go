package main

import "fmt"

func main4() {
	var age uint8
	var prenom string
	fmt.Print("Quel est votre âge ?")
	fmt.Print("\n")
	fmt.Scan(&age)
	fmt.Print("Quel est votre prénom ?")
	fmt.Print("\n")
	fmt.Scan(&prenom)
	if age >= 18 {
		fmt.Print("Vous êtes majeur, voici votre document !")
		fmt.Print("\n")
	} else {
		fmt.Print("Vous êtes mineur, rentrez chez vous !")
		fmt.Print("\n")
	}
}
