package main

import (
	"fmt"
	"os"
	"strconv"

	"PropertyFinder-Patika/week-1-homework-1-denizcamalan/code/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s \n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
	}
}

// input :
// ./cf 45 67 89 100

// output :
// 45°F = 7.222222222222222°C, 45°C = 113°F
// 67°F = 19.444444444444443°C, 67°C = 152.6°F
// 89°F = 31.666666666666668°C, 89°C = 192.2°F
// 100°F = 37.77777777777778°C, 100°C = 212°F
