package main

import "fmt"

func f23() {
	var emprun bool
	var Jour_emprun uint8
	var etat bool

	var t_actuel uint8
	var t_max uint8
	var t_min uint8
	var t_ux_humidite uint8
	var vent uint8
	var pluie bool
	fmt.Print("Avez-vous emprunté un parapluie ? (true/false)")
	fmt.Print("\n")
	fmt.Scan(&emprun)
	if emprun == true {
		fmt.Print("Quel jour avez-vous emprunté le parapluie ? (1-31)")
		fmt.Print("\n")
		fmt.Scan(&Jour_emprun)
		if Jour_emprun < 1 || Jour_emprun > 31 {
			fmt.Print("Jour invalide")
			fmt.Print("\n")
		} else {
			fmt.Print("Quel est l'état du parapluie ? (true/false)")
			fmt.Print("\n")
			fmt.Scan(&etat)
			if etat == false {
				fmt.Print("Parapluie perdu, amende de 25 euros")
				fmt.Print("\n")
			} else {
				fmt.Print("Quel est la température actuelle ? (en °C)")
				fmt.Print("\n")
				fmt.Scan(&t_actuel)
				fmt.Print("Quel est la température maximale aujourd'hui ? (en °C)")
				fmt.Print("\n")
				fmt.Scan(&t_max)
				fmt.Print("Quel est la température minimale aujourd'hui ? (en °C)")
				fmt.Print("\n")
				fmt.Scan(&t_min)
				fmt.Print("Quel est le taux d'humidité actuel ? (en %)")
				fmt.Print("\n")
				fmt.Scan(&t_ux_humidite)
				fmt.Print("Quel est la vitesse du vent actuel ? (en km/h)")
				fmt.Print("\n")
				fmt.Scan(&vent)
				fmt.Print("Est-il en train de pleuvoir ? (true/false)")
				fmt.Print("\n")
				fmt.Scan(&pluie)
				if t_actuel < 0 || t_actuel > 40 || t_max < 0 || t_max > 40 || t_min < 0 || t_min > 40 || t_ux_humidite < 0 || t_ux_humidite > 100 || vent < 0 || vent > 100 {
					fmt.Print("Valeur invalide")
					fmt.Print("\n")
				} else {
					if pluie == true || t_ux_humidite > 80 || vent > 50 || t_actuel < 0 || t_max-t_min > 15 {
						fmt.Print("Vous avez bien fait de prendre un parapluie !")
						fmt.Print("\n")
					} else {
						fmt.Print("Vous n'aviez pas besoin de votre parapluie aujourd'hui.")
						fmt.Print("\n")
					}
				}
			}
		}
	} else {
		fmt.Print("Quel est la température actuelle ? (en °C)")
		fmt.Print("\n")
		fmt.Scan(&t_actuel)
		fmt.Print("Quel est la température maximale aujourd'hui ? (en °C)")
		fmt.Print("\n")
		fmt.Scan(&t_max)
		fmt.Print("Quel est la température minimale aujourd'hui ? (en °C)")
		fmt.Print("\n")
		fmt.Scan(&t_min)
		fmt.Print("Quel est le taux d'humidité actuel ? (en %)")
		fmt.Print("\n")
		fmt.Scan(&t_ux_humidite)
		fmt.Print("Quel est la vitesse du vent actuel ? (en km/h)")
		fmt.Print("\n")
		fmt.Scan(&vent)
		fmt.Print("Est-il en train de pleuvoir ? (true/false)")
		fmt.Print("\n")
		fmt.Scan(&pluie)
		if t_actuel < 0 || t_actuel > 40 || t_max < 0 || t_max > 40 || t_min < 0 || t_min > 40 || t_ux_humidite < 0 || t_ux_humidite > 100 || vent < 0 || vent > 100 {
			fmt.Print("Valeur invalide")
			fmt.Print("\n")
		} else {
			if pluie == true || t_ux_humidite > 80 || vent > 50 || t_actuel < 0 || t_max-t_min > 15 {
				fmt.Print("Vous auriez dû prendre un parapluie !")
				fmt.Print("\n")
			} else {
				fmt.Print("Vous n'aviez pas besoin de parapluie aujourd'hui.")
				fmt.Print("\n")
			}
		}
	}
}

func main1() {
	f23()
}
