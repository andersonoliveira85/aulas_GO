package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("É manhã")
	default:
		fmt.Println("É tarde")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("É tinal de semana")
	default:
		fmt.Println("É semana")
	}

}
