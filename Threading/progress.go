// Create a progress bar for applications that can keep track of a download in
// progress. The progress bar will be on a separate thread and will communicate
// with the main thread using delegates.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	download := NewDownload("https://download.mozilla.org/?product=firefox-28.0-SSL&os=osx&lang=en-US",
		"firefox.dmg")
	progressBar := ProgressBar{download}
	progressBar.Start()
}

type Download struct {
	File          *os.File
	Response      *http.Response
	ContentLength int
	Done          bool
}

func (download *Download) StartDownload() {
	_, err := io.Copy(download.File, download.Response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		download.Done = true
		return
	}
	download.Done = true
}

func NewDownload(url string, fileName string) *Download {
	// Allocate memory for a new Download object
	download := new(Download)

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	download.File = out

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	download.Response = response
	download.ContentLength, _ = strconv.Atoi(response.Header.Get("content-length"))
	download.Done = false

	return download
}

func (download *Download) BytesDownloaded() int {
	info, err := download.File.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return int(info.Size())
}

type ProgressBar struct {
	Download *Download
}

func (progressBar *ProgressBar) Start() {
	// Run the download in a goroutine
	go progressBar.Download.StartDownload()
	progressBar.Show()

	// Remember to close the file and connection!
	progressBar.Download.Response.Body.Close()
	progressBar.Download.File.Close()
}

func (progressBar *ProgressBar) Show() {
	var progress int
	totalBytes := int(progressBar.Download.ContentLength)
	lastTime := false

	// While the download is in progress, keep redrawing the progress bar
	for !progressBar.Download.Done || lastTime {
		// Start the progress bar - carriage return to overwrite previous iteration
		fmt.Print("\r[")
		bytesDone := progressBar.Download.BytesDownloaded()
		progress = 40 * bytesDone / totalBytes

		// Build the progress bar
		for i := 0; i < 40; i++ {
			if i < progress {
				fmt.Print("=")
			} else if i == progress {
				fmt.Print(">")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("] ")
		fmt.Printf("%d/%dkB", bytesDone/1000, totalBytes/1000)
		time.Sleep(100 * time.Millisecond)

		// After the download completes, we need one more loop iteration
		if progressBar.Download.Done && !lastTime {
			lastTime = true
		} else {
			lastTime = false
		}
	}
	fmt.Println()
}
