package main

import "fmt"

func main() {
	// Variables in Go
	fmt.Println("======= Variables in Go =======")

	// Declaring A variable name and Data Type (Static Type declaration)
	var age int

	// Assign value to the declared Variable
	age = 37

	fmt.Println(age)

	// Declaring Variable Nmae and Data Type as well ass assigning the value
	var name = "Yassine"
	fmt.Println(name)

	// shorthand notation ( := ) =>  declaration + assignment + type inference (Dynamic Type declaration)
	id := "id-98768y8"
	fmt.Println(id)

	// Declaring multiple variables in the same line:
	country, city, number := "Morocco", "Casablanca", 23

	adress := fmt.Sprintf("Country: %s, City: %s, Number %d", country, city, number) // Country: Morocco, City: Casablanca, Number 23

	fmt.Println(adress)

	// Get Variable Type:
	price := 54.98
	fmt.Printf("Price is of type %T\n", price) // Price is of type float64

	// Constants, unchangebale variables in runtime.

	// const PI float32 // missing init expr for PI cannot separate declaration and assignment for constants variables
	const PI float32 = 3.141
	// PI = 87.9 // cannot assign to PI an other value
	fmt.Println(PI)

	// Integer literals
	// An integer litirals can be decimal, octal or hexadecimal

	const (
		DECIMAL     = 255  // decimal with no prefix
		OCTAL       = 0377 // octal with leading 0
		HEXADECIMAL = 0xff // hexadecimal with leading 0x
	)

	fmt.Printf("Decimal: %d\nOctal: %d\nHexadecimal: %d\n", DECIMAL, OCTAL, HEXADECIMAL)

	// Floating-point literals
	/*
	** A floating-ponit literal can have an integer part,
	** a decimal poin, a fraction part and an exponent part.
	 */

	const DISCOUNT = 6.564e23
	fmt.Println("Discount = ", DISCOUNT)

	// Escape sequences in String literals

	const GREATING = "Hello\tYassine!\n\"\aGo Language\" is Simple\n"
	print(GREATING)

}
