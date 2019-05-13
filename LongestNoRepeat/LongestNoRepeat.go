package main

import (
	"fmt"
)

func longest(s string) int {
	m := make(map[rune]int)
	start, maxLen := -1, 0
	for idx, letter := range []rune(s) {
		if val, ok := m[letter]; ok && start < val {
			start = val
		}
		m[letter] = idx
		if idx-start > maxLen {
			maxLen = idx - start
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longest("abcabcabb"))
	fmt.Println(longest("bbbbb"))
	fmt.Println(longest("pwwkew"))
	fmt.Println(longest("abcdef"))
	fmt.Println(longest("fjakhglkafgj"))
	fmt.Println(longest("房宝成房宝成"))
	fmt.Println(longest(""))

}
