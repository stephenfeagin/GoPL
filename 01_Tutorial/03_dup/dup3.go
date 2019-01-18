// Dup3 reads from a list of named files, and prints the lines that are
// repeated along with their counts.
// While Dup1 and Dup2 streamed their input, Dup3 reads the text all at once
// and processes it.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

        // ReadFile returns a byte slice, so we have to convert it to string in
        // order to use `strings.Split()`.
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
