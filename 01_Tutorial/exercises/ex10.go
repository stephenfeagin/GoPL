// Find a web site that produces a large amount of data. Investigate caching by
// running `fetchall` twice in succession to see whether the reported times change
// much. Do you get the same content each time? Modify `fetchall` to print its
// output to a file so it can be examined.
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
	start := time.Now()

	// A channel is a communication mechanism that allows one goroutine to pass
	// values of a specified type to another goroutine.
	ch := make(chan string)
	for index, url := range os.Args[1:] {
		// A goroutine is a concurrent function execution.
		go fetch(index, url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(index int, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Send err to channel ch
		return
	}

	destination, err := os.Create(strconv.Itoa(index) + ".txt")
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer destination.Close()

	nbytes, err := io.Copy(destination, resp.Body)
	resp.Body.Close() // Don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
