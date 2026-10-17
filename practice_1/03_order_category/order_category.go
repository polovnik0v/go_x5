package main

import "fmt"

func OrderCategory(amount int) {
	if amount < 0 {
		fmt.Println("Ошибка: сумма заказа не может быть отрицательной")
		return
	}

	var category string
	if amount < 10000 {
		category = "small"
	} else if amount < 50000 {
		category = "regular"
	} else {
		category = "large"
	}

	fmt.Println(fmt.Sprintf("Сумма заказа: %d коп.", amount))
	fmt.Println(fmt.Sprintf("Категория: %s \n", category))
}

func main() {
	OrderCategory(0)
	OrderCategory(9999)
	OrderCategory(10000)
	OrderCategory(49999)
	OrderCategory(50000)
	OrderCategory(-1)
}
