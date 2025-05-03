package main

import (
	"fmt"
)

func multiply(a, b int) {
	fmt.Println("Результат:", a*b)
}
func div(a, b int) {
	if b == 0 {
		fmt.Println("Деление на ноль")
	} else {
		fmt.Println("Result:", a/b)
	}
}
func main() {

}
