package main

import (
	"fmt"
)

func main() {
	n := 1 == 1
	m := 1 == 2
	var isEmployee bool = false
	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", m, m)

	fmt.Printf("%v, %T\n", isEmployee, isEmployee)
}
