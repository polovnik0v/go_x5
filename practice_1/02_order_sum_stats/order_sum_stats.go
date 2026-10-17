package main

import (
	"fmt"
)

func OrderSumStats(orders []int) {
	if len(orders) == 0 {
		fmt.Println("Нет заказов для анализа")
		return
	}
	min, max := orders[0], orders[0]
	total := 0

	for _, v := range orders {
		total += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	avg := float64(total) / float64(len(orders))

	fmt.Println(fmt.Sprintf("Количество заказов: %d", len(orders)))
	fmt.Println(fmt.Sprintf("Минимальный заказ: %d коп.", min))
	fmt.Println(fmt.Sprintf("Максимальный заказ: %d коп.", max))
	fmt.Println(fmt.Sprintf("Общая сумма: %d коп.", total))
	fmt.Println(fmt.Sprintf("Средний заказ: %v коп.", avg))
}

func main() {
	OrderSumStats([]int{12000, 7500, 31000, 4500})
	fmt.Println("___________________________")
	OrderSumStats([]int{})
}
