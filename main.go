package main

import "fmt"

func deleteRepit(nums []int) {
	k := 1
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println(k)
}

func main() {
	fmt.Println("Version 4")
	array := []int{0, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	deleteRepit(array)
	fmt.Println(array)
}
