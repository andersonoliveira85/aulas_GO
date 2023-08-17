package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a = "uma string"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(c, b)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	var num01 int = 10
	fmt.Println(num01)

	num02 := 5.75
	fmt.Println(num02)

	f := "banana"
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(num02)
	fmt.Println(reflect.TypeOf(num02))

}
