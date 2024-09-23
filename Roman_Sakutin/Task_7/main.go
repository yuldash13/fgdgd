package main

import "fmt"

func main() {

	var people int
	minutes := 10

	var hour string

	var (
		hour1  = "час"
		hour2  = "часа"
		hour3  = "часов"
		minute = "минут"
	)

	fmt.Println("Сколько бабулек в очереди?")
	fmt.Scan(&people)

	waiting2 := people * minutes
	waiting1 := waiting2 / 60
	waiting := waiting2 % 60

	if waiting1 == 1 {
		hour = hour1
	} else if waiting1 > 1 && waiting1 < 5 {
		hour = hour2
	} else if waiting1 >= 5 || waiting1 == 0 {
		hour = hour3
	}

	fmt.Printf("Вам придется ждать: %d %s, %d %s.", waiting1, hour, waiting, minute)
}
