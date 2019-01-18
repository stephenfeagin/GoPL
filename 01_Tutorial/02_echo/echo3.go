package main

import (
	"fmt"
	"os"
	"strings"
)

// Re-creating new strings with s += sep + arg can be costly if done repeatedly
// It has to re-allocate the memory for the string and the old string will have
// to be garbage collected. Using strings.Join is more efficient.
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
