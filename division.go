package main

import (
	"errors"
	"fmt"
)

func main() {

	/*a := 8
	b := 0

	fmt.Println(division(a, b))*/

	div, err := Division(5, 3)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Division: ", div)

}

func Division(a int, b int) (int, error) {

	if b == 0 {
		return -1, errors.New("el divisor no puede ser 0")
	}

	return a / b, nil

	//panic: runtime error: integer divide by zero
}
