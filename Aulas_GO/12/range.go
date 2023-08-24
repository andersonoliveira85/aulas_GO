package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}

	sum := 0

	//range no array/slice
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	//range em Map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
