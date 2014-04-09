// Checks if the string entered by the user is a palindrome. That is that it
// reads the same forwards as backwards like “racecar”
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the text to determine if it is a palindrome: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if isPalindrome(input) {
		fmt.Println(input[:len(input)-1], "is a palindrome!")
	} else {
		fmt.Println(input[:len(input)-1], "is NOT a palindrome!")
	}
}

func isPalindrome(s string) bool {
	text := []rune(s)
	n := len(text)

	for i := 0; i < n/2; i++ {
		if text[i] != text[n-i-2] {
			return false
		}
	}
	return true
}
