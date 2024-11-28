package main

import (
	"fmt"
)

// QuickSort сортирует слайс "на месте".
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return // Уже отсортировано
	}

	// Выбор pivot (опорного элемента)
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Разделение массива
	left, right := 0, len(arr)-1

	// Перемещаем элементы относительно pivot
	for left <= right {
		// Сдвигаем указатели до тех пор, пока не найдём элемент для обмена
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// Рекурсивно сортируем две части
	QuickSort(arr[:right+1]) // Левая часть
	QuickSort(arr[left:])    // Правая часть
}

func main() {
	data := []int{6, 9, 4, 2, 5, 3, 7, 0, 1, 8}
	fmt.Println("До сортировки:", data)
	QuickSort(data)
	fmt.Println("После сортировки:", data)
}
