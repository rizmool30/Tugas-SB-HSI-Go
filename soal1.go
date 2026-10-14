package main

import (
	"fmt"
)

func main() {

	var a int
	var b int
	var c int

	fmt.Print("masukan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("masukan nilai b: ")
	fmt.Scan(&b)

	fmt.Print("masukan nilai c: ")
	fmt.Scan(&c)

	if (a*a)+(b*b) == (c * c) {
		fmt.Println("pytagoras")
	} else {
		fmt.Println("not pytagoras")
	}

}
	