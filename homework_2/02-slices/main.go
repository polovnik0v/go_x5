package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	if len(numbers) != 0 {
		for i := range numbers {
			sum += numbers[i]
		}
	}
	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(Sum(numbers))                // 30
	fmt.Println(Sum([]int{}))                // 0
	fmt.Println(Sum(nil))                    // 0
	fmt.Println(Sum([]int{-2, -1, 0, 1, 2})) // 0

	numbers = append(numbers, 12)
	fmt.Println(len(numbers), Sum(numbers)) // 6 42

	shared := numbers
	shared[0] = 100
	fmt.Println(numbers[0], shared[0]) // 100 100
}

// Объясните: что копируется при присваивании среза и при передаче его в Sum?
// Копируются указатель на срез, длина и капаситет

// Почему изменение через shared видно через numbers? Зачем сохранять результат append?
// Потому что shared смотрит на ту же область памяти что и numbers

// Зачем сохранять результат append?
// append возвращает новый заголовок, поэтому результат нужно присваивать
