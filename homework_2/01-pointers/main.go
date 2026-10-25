package main

import "fmt"

func Increment(value *int) {
	if value == nil {
		return
	}
	(*value)++
}

func main() {
	number := 10
	Increment(&number)
	fmt.Println(number) // 11

	Increment(&number)
	fmt.Println(number) // 12

	Increment(nil) // программа продолжает работать
}

// Объясните: что означают &number и *value? Почему для изменения исходного числа здесь нужен указатель?
// &number — получение адреса переменной в памяти
// *value — разыменовывание указателя, т.е. обращение к значению переменной по адресу
