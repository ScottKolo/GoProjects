// Create a progress bar for applications that can keep track of a download in
// progress. The progress bar will be on a separate thread and will communicate
// with the main thread using delegates.

package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  )

func main() {
  downloadFromUrl("http://www.google.com/robots.txt", "robots.txt")
}

// Adapted from GitHub user thbar's golang-playground
// https://github.com/thbar/golang-playground/blob/master/download-files.go
func downloadFromUrl(url string, fileName string) {
  fmt.Println("Starting Download")

  out, err := os.Create(fileName)
  defer out.Close()
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  resp, err := http.Get(url)
  defer resp.Body.Close()
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  contentLength := resp.ContentLength
  fmt.Println("Starting Download of length", contentLength, "bytes")
  n, err := io.Copy(out, resp.Body)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  fmt.Println("File created of length", n, "bytes")
}
