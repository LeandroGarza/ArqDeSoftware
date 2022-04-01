package main

import "fmt"

func main() {

	a := 8
	b := 0

	fmt.Println(division(a, b))

	/*div, err := division(5, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Division:", div)*/

}

func division(a int, b int) int {
	return a / b

	//panic: runtime error: integer divide by zero
}
