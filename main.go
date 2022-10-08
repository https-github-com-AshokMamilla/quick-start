package main

import (
	"fmt"
)

func main() {

	//a := 5
	//b := 4
	//you should use fmt.capital letter for any function in fmt or any other package
	//fmt.println(a + b)
	var a int
	var b int
	var c int

	fmt.Print("what is the value a: ")
	//fmt.Scanf("" & a)
	fmt.Scan(&a)
	fmt.Print("what is the value b: ")
	//fmt.Scanf(&b)
	fmt.Scan(&b)
	c = a + b
	//fmt.Scan(&c)
	//fmt	.Print("value of c")
	fmt.Print("value of c=", c)
	//fmt.Scan(&c)
	//fmt.Scan(&c)
}
