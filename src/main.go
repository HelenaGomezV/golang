package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	text2 := strings.ToLower(text)
	textReverse := ""
	for i := len(text2) - 1; i >= 0; i-- {
		textReverse += string(text2[i])
	}
	if text2 == textReverse {
		fmt.Println("is Palindrome")
	} else {
		fmt.Println("no es")
	}

}

func main() {
	isPalindrome("Ama")

}
