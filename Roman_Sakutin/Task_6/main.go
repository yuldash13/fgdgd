package main

import "fmt"

func main() {

	var gold, quantity int

	var (
		gems = "кристаллы"
	)

	var (
		gemsCost = 3
	)

	fmt.Println("Я: Ох, наконец-то я дошел до магазина! Нести мешок золота было очень сложно. Сколько у меня там его?")
	fmt.Scan(&gold)
	fmt.Printf("Я: И зачем мне дали %d золота? Ах да, я должен купить %s.", gold, gems)
	fmt.Println()
	fmt.Println("Торговец: Здравствуйте, чем могу помочь?")
	fmt.Printf("Я: Здравствуйте, %s в наличии? Сколько они стоят?", gems)
	fmt.Println()
	fmt.Printf("Торговец: Да, есть. Один стоит %d золота.", gemsCost)
	fmt.Println()
	fmt.Println("Торговец: Сколько вам нужно?")
	fmt.Scan(&quantity)

	finalCost := quantity * gemsCost
	remainingGold := gold - finalCost

	if remainingGold >= 0 {
		fmt.Printf("Торговец: C вас %d золота.", finalCost)
	} else if remainingGold < 0 {
		fmt.Println("Торговец: У вас недостаточно золота. Возьмите меньше.")
		fmt.Scan(&quantity)
		fmt.Printf("Торговец: C вас %d золота.", finalCost)
	}

	fmt.Println()
	fmt.Println("Я: Вот держите.")
	fmt.Printf("(Получено %d кристаллов, осталось %d золота)", quantity, remainingGold)
}
