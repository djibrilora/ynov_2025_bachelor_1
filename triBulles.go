package main

import (
	"fmt"
	"time"
)

func triBulles(liste []int) []int {
	for i := 0; i < len(liste)-1; i++ {
		for j := 0; j < len(liste)-i-1; j++ {
			if liste[j] > liste[j+1] {
				liste[j], liste[j+1] = liste[j+1], liste[j]
			}
		}
	}
	return liste
}

func main() {
	nombres := []int{33, 10, 55, 71, 29, 5, 18}
	fmt.Println("Avant tri :", nombres)
	start := time.Now()
	trie := triBulles(nombres)
	duration := time.Since(start)
	fmt.Println("Après tri :", trie)
	fmt.Printf("Temps d'exécution de triBulles : %v\n", duration)
}
