// Enter a string and the program will reverse it and print it out.

package main

import (
  "fmt"
  "bufio"
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

  text := make([]rune, len(input))
  n := 0

  for _, r := range(input) {
    text[n] = rune(r)
    n++
  }

  text = text[:n-1]
  for i := 0; i < n/2; i++ {
    text[i], text[n-i-2] = text[n-i-2], text[i]
  }

  fmt.Println(string(text))
}
