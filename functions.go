package main

import "fmt"

// simple function with parameters and a return value
func add(a int, b int) int {
	return a + b
}

// function with a single string return value
func greet(name string) string {
	return "Hello, " + name
}

// function returning multiple values
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func demoFunctions() {
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	message := greet("Vishnu")
	fmt.Println(message)

	q, r := divide(17, 5)
	fmt.Println("Quotient:", q, "Remainder:", r)
}
