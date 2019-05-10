/*
Сортировка вставками, имея по-прежнему O(n2), работает несколько иначе.
Она всегда поддерживает в сортированном виде подсписок на нижних индексах списка.
Каждый новый элемент “вставляется” в упорядоченный на прошлой итерации подсписок так, чтобы тот остался сортированным и стал на один элемент больше.
*/

package main

import "fmt"

func main() {
	baseData := []int{11, 4, 7, 3, 8, 5, 0, 1, 3, 5, 43, 12}
	fmt.Println("Array before: ", baseData)

	newData, status := insertionSort(baseData)

	if status == "ok" {
		fmt.Println("Array after: ", newData)
	} else {
		fmt.Println("Empty array/slice!")
	}
}

func insertionSort(data []int) ([]int, string) {
	status := "error"
	length := len(data)

	if length > 1 {
		for i := 1; i < length; i++ {
			for j := 0; j < i; j++ {
				if data[i] < data[j] {
					data[i], data[j] = data[j], data[i]
				}
			}

			fmt.Println(data)
		}

		status = "ok"
	}

	return data, status
}
