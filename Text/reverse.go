// Enter a string and the program will reverse it and print it out.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the text to be reversed: ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	text := []rune(input)
	n := len(text)

	for i := 0; i < n/2; i++ {
		text[i], text[n-i-2] = text[n-i-2], text[i]
	}

	fmt.Println(string(text))
}
