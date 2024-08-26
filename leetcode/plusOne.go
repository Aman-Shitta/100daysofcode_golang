// https://leetcode.com/problems/plus-one/description/
package main

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 1; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	digits := []int{1, 2, 3}
	println(plusOne(digits))
}
