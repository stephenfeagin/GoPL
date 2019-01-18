// Dup1 prints the text of each line that appears more than once in the
// standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// `map`s provide constant time operations to store, retrieve, or test for
	// an item in a set.
	// Key must be a type that can be compared with `==`, but value can be
	// anything.
	// In the next line, we create a map with strings as keys and ints as
	// values.
	counts := make(map[string]int)

	// `bufio.Scanner` reads inputs and breaks it into lines or words
	input := bufio.NewScanner(os.Stdin)

	// `.Scan()` returns a boolean for if there is another line to read. If
	// there is, it advances one spot to read that item. It also removes the
	// newline character from the end of the string.
	for input.Scan() {
		// It doesn't matter that there is no pair in counts that corresponds
		// to `input.Text()`. As soon as it's referenced, it is initialized
		// with its type's zero value.
		counts[input.Text()]++
	}

	// NOTE: Ignoring potential errors from `input.Err()`.
	// Range over a map yields key, value.
	for line, n := range counts {
		if n > 1 {
			// Print a formatted string from a list of expressions, using
			// "verbs" to specify how they should be formatted.
            // Printf does not print an ending newline by default.
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
