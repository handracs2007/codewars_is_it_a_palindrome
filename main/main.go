package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)

	for i := 0 ; i < len(str) / 2; i ++{
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsPalindrome("nope"))
	fmt.Println(IsPalindrome("Yay"))
}
