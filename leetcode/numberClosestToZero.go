package main

// Given an integer array nums of size n,
// return the number with the value closest to 0 in nums.
// If there are multiple answers, return the number with the largest value.

import (
	"fmt"
	"math"
)

// func findClosestNumber(nums []int) int { // O(n²)

// 	closestNums := make(map[float64][]int)
// 	for _, num := range nums {
// 		numAbsDist := math.Abs(float64(num - 0))
// 		if closetNumList, ok := closestNums[numAbsDist]; ok {
// 			closestNums[numAbsDist] = append(closetNumList, num)
// 		} else {
// 			closestNums[numAbsDist] = []int{num}
// 		}
// 	}
// 	smallest_num := math.Inf(10)
// 	for key, valueList := range closestNums {

// 		if math.Abs(key) < math.Abs(smallest_num) {
// 			if len(valueList) > 1 {
// 				someNum := math.Inf(-1)
// 				for _, val := range valueList {
// 					if float64(val) > someNum {
// 						someNum = float64(val)
// 					}
// 				}
// 				smallest_num = someNum
// 			} else {
// 				smallest_num = float64(valueList[0])
// 			}
// 		}
// 	}
// 	return int(smallest_num)
// }

func findClosestNumber(nums []int) int { // O(n)
	smallestNum := nums[0]

	for _, num := range nums {

		if math.Abs(float64(num)) <= math.Abs(float64(smallestNum)) {
			if math.Abs(float64(num)) == math.Abs(float64(smallestNum)) && num != smallestNum {
				smallestNum = int(math.Abs((float64(num))))
			} else {
				smallestNum = num
			}
		}
	}
	return smallestNum
}

func main() {
	nums := []int{-4, -2, 1, 4, 8}
	fmt.Print(findClosestNumber(nums))
}
