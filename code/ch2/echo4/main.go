package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

/*
			Usage of ./echo4:
			-n omit trailing newline
			-s string
				separator (default " ")
*/

// input  : ./echo4 -s / a bmmb ff
// output : a/bmmb/ff
