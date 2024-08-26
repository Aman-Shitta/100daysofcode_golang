package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	count := 0
	var candidate int

	for idx, num := range nums {
		fmt.Println(count)
		if idx == 0 {
			nums[count] = num
			count = 1
		} else if num != candidate {
			nums[count] = num
			count += 1
		}
		candidate = num
		nums[idx] = 0
	}
	return count
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums := []int{1, 1, 2}
	fmt.Println(nums)
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
