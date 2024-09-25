package main

import "fmt"

func main() {

	var (
		chest1 = "кристаллы"
		chest2 = "золото"
		chest3 = "алмазы"
	)

	var chest4 string

	fmt.Printf("На нашем складе есть сундуки с %s, %s, %s.", chest1, chest2, chest3)
	fmt.Println()
	fmt.Println("Что это такое, кто все перепутал? Быстро исправьте все!")

	chest4 = chest1
	chest1 = chest3
	chest3 = chest2
	chest2 = chest4

	fmt.Println("Спустя какое-то время...")
	fmt.Printf("Так в сундуках теперь находятся: %s, %s, %s.", chest1, chest2, chest3)
	fmt.Println()
	fmt.Print("Наконец-то все на своих местах. Чтобы такого больше не повторялось!")
}
