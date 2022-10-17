package main

import "fmt"

func addition() {
	fmt.Println("Please enter 2 Numbers: ")
	var num1, num2 float64
	fmt.Scan(&num1, &num2)
	var sum float64 = num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)
}

func subtraction() {
	fmt.Println("Please enter 2 numbers: ")
	var num1, num2 float64
	fmt.Scan(&num1, &num2)
	var difference float64 = num1 - num2
	fmt.Println("The difference between", num1, "and", num2, "is", difference)
}

func multiplication() {
	fmt.Println("Please enter 2 numbers: ")
	var num1, num2 float64
	fmt.Scan(&num1, &num2)
	var product float64 = num1 * num2
	fmt.Println("The product of", num1, "and", num2, "is", product)
}

func division() {
	fmt.Println("Please enter 2 numbers: ")
	var num1, num2 float64
	fmt.Scan(&num1, &num2)
	var quotent float64 = num1 / num2
	fmt.Println("The quotent of", num1, "and", num2, "is", quotent)
}

func power() {
	var exponent, base int
	fmt.Print("Enter base: ")
	fmt.Scan(&base)
	fmt.Print("Enter exponent: ")
	fmt.Scan(&exponent)

	output := 1
	for exponent != 0 {
		output *= base
		exponent -= 1
	}
	fmt.Printf("The output of power calculation is %d", output)
}

func main() {
	fmt.Println("Welcome to the Interactive Calculator")
	fmt.Println("Enter in an operation: (+, -, *, /, ** (power))")

	// User Input
	var operation string
	fmt.Scan(&operation)

	// Conditionals for Operations
	if operation == "+" {
		addition()
	} else if operation == "-" {
		subtraction()
	} else if operation == "*" {
		multiplication()
	} else if operation == "/" {
		division()
	} else if operation == "**" {
		power()
	} else {
		fmt.Println("Invalid Operator")
	}

}
