package main

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
import "fmt"

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	for idx := range len(haystack) {
		if haystack[idx] == needle[0] && idx+len(needle) <= len(haystack) && haystack[idx:idx+len(needle)] == needle {
			return idx
		}
	}
	return -1
}

func main() {
	haystack := "abc"
	needle := "c"
	fmt.Println(strStr(haystack, needle))
}
