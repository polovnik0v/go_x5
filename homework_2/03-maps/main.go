package main

import "fmt"

func StudentScores() map[string]int {
	scoresMap := make(map[string]int)

	scoresMap["Анна"] = 3
	scoresMap["Борис"] = 0
	scoresMap["Вера"] = 5
	scoresMap["Анна"] += 2

	delete(scoresMap, "Вера")

	return scoresMap
}

func main() {
	scores := StudentScores()
	fmt.Println(scores["Анна"])

	for _, student := range []string{"Борис", "Глеб"} {
		score, ok := scores[student]
		fmt.Println(student, score, ok)
	}

	fmt.Println(len(scores))
}

// Почему одних баллов 0 недостаточно, чтобы понять, есть ли студент в map?
// Чтение отсутствующего ключа возвращает zero value, который равен реальному значению 0,
// ok явно говорит, был ли ключ в map

// Что произойдёт при записи, если вместо make использовать только var scores map[string]int?
// var scores map[string]int создаёт nil-map. Чтение из неё безопасно и даёт zero value,
// но первая же запись вызовет panic
