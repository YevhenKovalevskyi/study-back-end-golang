/*
Пузырьковая сортировка состоит в повторяющихся проходах по сортируемому массиву.
За каждый проход элементы последовательно сравниваются попарно и, если порядок в паре неверный, выполняется обмен элементов.
Проходы по массиву повторяются до тех пор, пока на очередном проходе не окажется, что обмены больше не нужны, что означает — массив отсортирован.
При проходе алгоритма, элемент, стоящий не на своём месте, «всплывает» до нужной позиции как пузырёк в воде, отсюда и название алгоритма.
Для понимания и реализации этот алгоритм — простейший, но эффективен он лишь для небольших массивов. Сложность алгоритма: O(n²).
*/

package main

import "fmt"

func main() {
	baseData := []int{1, 4, 7, 3, 8, 9, 0, 1, 3, 5, 43, 12}
	fmt.Println("Array before: ", baseData)

	newData, status := bubbleSort(baseData)

	if status == "ok" {
		fmt.Println("Array after: ", newData)
	} else {
		fmt.Println("Empty array/slice!")
	}
}

func bubbleSort(data []int) ([]int, string) {
	status := "error"
	length := len(data)

	if length > 1 {
		for i := 0; i < length; i++ {
			for j := i + 1; j < length; j++ {
				if data[i] > data[j] {
					data[i], data[j] = data[j], data[i]
				}
			}

			fmt.Println(data)
		}

		status = "ok"
	}

	return data, status
}
