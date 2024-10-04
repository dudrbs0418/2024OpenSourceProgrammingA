package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, b, i, s)
	fmt.Printf("%f%t%s%d", f, b, s, i)
	f = 2.7
	i = 3
	fmt.Print(f > float64(i), "\n")
}
