// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// If to concurrent requests try to access the same variable at the same time,
// the variable may not be updated properly. This is called a race condition.
// The Mutex blocks access to the variable so that race conditions are avoided.
var mu sync.Mutex
var count int

// The server has two handlers: handler and counter.
// The URL determines which one is called. The server runs the handler for
// all incoming requests through separate goroutines, so multiple requests
// can be handled simultanesously.
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
