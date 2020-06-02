// Flag prints its command-line arguments.

package main

import (
	"flag"
	"fmt"
	"strings"
)

// The variables sep and n are pointers to the flag variables, which must be
// accessed indirectly as *sep and *n
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
