package main

import "fmt"

func adicionar(x, y int) int {
	return x + y
}
func main() {
	res := adicionar(2, 45)
	fmt.Println("O  resultado da soma é: ", res)
}
