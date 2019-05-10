/*
Сортировка перемешиванием, или Шейкерная сортировка, или двунаправленная — разновидность пузырьковой сортировки. Анализируя метод пузырьковой сортировки, можно отметить два обстоятельства.
Во-первых, если при движении по части массива перестановки не происходят, то эта часть массива уже отсортирована и, следовательно, её можно исключить из рассмотрения.
Во-вторых, при движении от конца массива к началу минимальный элемент «всплывает» на первую позицию, а максимальный элемент сдвигается только на одну позицию вправо.
Эти две идеи приводят к следующим модификациям в методе пузырьковой сортировки.
Границы рабочей части массива (то есть части массива, где происходит движение) устанавливаются в месте последнего обмена на каждой итерации.
Массив просматривается поочередно справа налево и слева направо.
Лучший случай для этой сортировки — отсортированный массив (O(n)), худший — отсортированный в обратном порядке (O(n2)).
Наименьшее число сравнений в алгоритме Шейкер-сортировки C = N−1. Это соответствует единственному проходу по упорядоченному массиву (лучший случай)
*/

package main

import "fmt"

func main() {
	baseData := []int{1, 4, 7, 3, 8, 9, 0, 1, 3, 5, 43, 12}
	fmt.Println("Array before: ", baseData)

	newData, status := shakerSort(baseData)

	if status == "ok" {
		fmt.Println("Array after: ", newData)
	} else {
		fmt.Println("Empty array/slice!")
	}
}

func shakerSort(data []int) ([]int, string) {
	status := "error"
	length := len(data)

	if length > 1 {
		left := 0
		right := length - 1

		for left < right {
			for i := left; i < right; i++ {
				if data[i] > data[i+1] {
					data[i], data[i+1] = data[i+1], data[i]
				}
			}

			right--

			for i := right; i > left; i-- {
				if data[i] < data[i-1] {
					data[i], data[i-1] = data[i-1], data[i]
				}
			}

			left++

			fmt.Println(data)
		}

		status = "ok"
	}

	return data, status
}