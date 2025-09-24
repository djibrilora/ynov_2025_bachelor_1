package main

import "fmt"

var t_matin int
var t_aprem int
var t_max int
var t_min int

func main2() {
	fmt.Print("Quel temperature fait-il ce matin")
	fmt.Print("\n")
	fmt.Scan(&t_matin)
	fmt.Print("Quel temperature fait-il cette après-midi")
	fmt.Print("\n")
	fmt.Scan(&t_aprem)
}
