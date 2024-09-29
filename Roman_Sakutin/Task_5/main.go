package main

import "fmt"

func main() {

	var (
		chest1 = "кристаллы"
		chest2 = "золото"
		chest3 = "алмазы"
	)

	fmt.Printf("На нашем складе есть сундуки с %s, %s и %s.", chest1, chest2, chest3)
	fmt.Println()
	fmt.Println("Что это такое, кто все перепутал? Быстро исправьте все!")

	chest1, chest2, chest3 = chest2, chest3, chest1

	fmt.Println("Спустя какое-то время...")
	fmt.Printf("Так в сундуках теперь находятся: %s, %s, %s.", chest1, chest2, chest3)
	fmt.Println()
	fmt.Print("Наконец-то все на своих местах. Чтобы такого больше не повторялось!")
}
