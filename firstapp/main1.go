package main

import "fmt"

var i int = 2000

func main() {
	var i int
	i = 42
	var j float64 = 28
	k := 999.
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("%v, %T", k, k)
}
