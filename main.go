package main

import (
	"fmt"
	"os"
	"strconv"
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
func add(a, b int) {
	fmt.Println("Результат:", a+b)
}
func sub(a, b int) {
	fmt.Println("Результат", a-b)
}

func main() {
	cmd := os.Args[1]
	a, _ := strconv.Atoi(os.Args[2])
	b, _ := strconv.Atoi(os.Args[3])

	if cmd == "add" {
		add(a, b)
	} else if cmd == "sub" {
		sub(a, b)
	} else if cmd == "mul" {
		mul(a, b)
	} else if cmd == "div" {
		div(a, b)
	} else {
		fmt.Println("Error")
	}
}
