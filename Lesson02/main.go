package main

import (
	"Lesson02/fibonacci"
	"fmt"
)

func main() {

	var number int

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)

	for i := 0; i <= number; i++ {
		fmt.Print(fibonacci.FibMap(i))
		fmt.Print(" ")
	}
	fmt.Println(" ")

}
