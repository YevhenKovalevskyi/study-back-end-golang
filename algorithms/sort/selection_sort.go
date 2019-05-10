/*
Сортировка выбором улучшает пузырьковую сортировку, совершая всего один обмен за каждый проход по списку.
Чтобы сделать это, она ищет наибольший элемент и помещает его на соответствующую позицию.
Как и для пузырьковой сортировки, после первого прохода самый большой элемент находится на правильном месте.
После второго - на своё место становится следующий наибольший элемент.
Процесс продолжается, требуя n−1 проход для сортировки n элементов, поскольку последний из них автоматически оказывается на своём месте.
*/

package main

import "fmt"

func main() {
	baseData := []int{1, 4, 7, 3, 8, 9, 0, 1, 3, 5, 43, 12}
	fmt.Println("Array before: ", baseData)

	newData, status := selectionSort(baseData)

	if status == "ok" {
		fmt.Println("Array after: ", newData)
	} else {
		fmt.Println("Empty array/slice!")
	}
}

func selectionSort(data []int) ([]int, string) {
	status := "error"
	length := len(data)

	if length > 1 {
		for i := 0; i < length-1; i++ {
			min := i

			for j := i + 1; j < length; j++ {
				if data[j] < data[min] {
					min = j
				}
			}

			data[i], data[min] = data[min], data[i]

			fmt.Println(data)
		}

		status = "ok"
	}

	return data, status
}
