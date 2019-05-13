package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	s := "yes我爱慕课网!"
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", 
		utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()
}
