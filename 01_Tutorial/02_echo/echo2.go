// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

    // Go does not allow for unused variables, so we can't do something like
    // `for i, arg := range ...`
    // Instead, the blank identifier, `_`, discards the unwanted variable
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
