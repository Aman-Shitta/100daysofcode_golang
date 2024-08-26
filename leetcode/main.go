
package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/length-of-last-word/submissions/1365914323/
func lengthOfLastWord(s string) int {
    
    stringList := strings.Split(s, " ")
    for i := len(stringList); i > 0; i-- {
        if len(stringList[i-1]) > 0 {
            return len(stringList[i-1])
        }
    }
    return 0
}


// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
    indices := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if idx, ok := indices[complement]; ok {
            return []int{idx, i}
        }
        indices[num] = i
    }
    return nil
}


func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))

	nums := []int{2,7,11,15}
	target :=  9
	fmt.Println(twoSum(nums, target))
}
