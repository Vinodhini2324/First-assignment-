package main

import (
	"fmt"
)

func main() {
	a := 13
	// var b *int = &a
	var b = &a
	*b = 10
	fmt.Println(a)
	fmt.Println(*b)

	pointerWithStruct()
}

type details struct {
	name string
}

func pointerWithStruct() {
	// var d *details = &details{}
	var d *details
	// d = &details{"vinoen"}
	d = new(details)
	(*d).name = "Vinoen"
	fmt.Println((*d).name)
	// fmt.Println(*d)
}
