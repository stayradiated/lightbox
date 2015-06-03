package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {

	var a, b float64

	flag.Float64Var(&a, "a", 0, "a")
	flag.Float64Var(&b, "b", 0, "b")
	flag.Parse()

	var ratio = a / b

	for i := float64(1); i < b; i++ {
		c := i / ratio
		if math.Mod(c, 1) == 0 {
			fmt.Println(i, c)
		}
	}

}
