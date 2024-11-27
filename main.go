package main

import "fmt"

func deleteDuplicates(nums []int) {
	if len(nums) == 0 {
		return
	}
	// Указатель на уникальный элемент.
	uniqueIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex-1] {
			nums[uniqueIndex] = nums[i]
			uniqueIndex++
		}
	}
	fmt.Printf("Количество уникальных чисел: %v\n", uniqueIndex)
}

func main() {
	fmt.Println("Version 4")
	array := []int{0, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}

	// Обработка массива.
	deleteDuplicates(array)

	// Вывод результата.
	fmt.Println("Array without duplicates:", array)
}
