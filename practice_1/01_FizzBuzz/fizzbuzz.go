package main

import "fmt"

func FizzBuzz() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("0FizzBuzz")
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}

func main() {
	FizzBuzz()
}
