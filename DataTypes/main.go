package main

import "fmt"

func main(){

	/* 
	** Go's basic (primitive) types are the building blocks for all other data structures in the language. They are categorized into three main groups: Boolean, Numeric, and String types. 
	** The specific primitive types are listed below: 
	** 
	** ## bool: Represents a boolean value, which is either true or false.
	** ## Numeric Types:
	** ### > Integers:
	**   	* Signed integers: int8, int16, int32, int64, and int. The plain int type is generally recommended unless a specific size is required; its size is platform-dependent (either 32 or 64 bits).
	**   	* Unsigned integers: uint8, uint16, uint32, uint64, and uint.
	** ### > Aliases: byte is an alias for uint8, and rune is an alias for int32 (used to represent a Unicode code point).
	** ### > Other: uintptr is an unsigned integer type large enough to store the uninterpreted bits of a pointer value.
	**
	** ### > Floating-point numbers: float32 and float64.
	**
	** ## > Complex numbers: complex64 and complex128.
	** ## > string: A sequence of bytes, conventionally but not necessarily representing UTF-8 encoded text. String values are immutable in Go. 
	** 
	*/

	// Section1: Integers in Go
	// Declaration: var var_nam type


	// A: Signed Integers

	var i int //  Platform-dependent (32 or 64 bits), generally preferred for most uses. 
	var i8 int8 // int8: -128 to 127
	var i16 int16 // int16: -32,768 to 32,767
	var i32 int32 // int32 (alias rune): -2,147,483,648 to 2,147,483,647
	var i64 int64 // int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

	// Value Assignment
	i 	= -182
	i8  = 127
	i16 = -3278
	i32 = 2147483647
	i64 = -3987509348750913

	// B: Unsigned Integers

	var u uint // uint: Platform-dependent (32 or 64 bits).
	var u8 uint8 // uint8 (alias byte): 0 to 255
	var u16 uint16 // uint16: 0 to 65,535
	var u32 uint32 // uint32: 0 to 4,294,967,295
	var u64 uint64 // uint64: 0 to 18,446,744,073,709,551,615
	// var uintlarge uintptr    //: An unsigned integer type large enough to store the bit pattern of any pointer. 
	// Value Assignment
	u 	= 255
	u8  = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 1844674473709551615

	fmt.Println("Signed Integers", i, i8, i16, i32, i64)
	fmt.Println("USigned Integers", u, u8, u16, u32, u64)
	

	// Section2: Float in Go

	//  A: Float32: Used for less precise calculations
	var f32 float32 = 10.06 // Approximately -3.4e+38 to 3.4e+38
	var LP float32 = 10123456789012345
	fmt.Println("Float32: ", f32)
	//  B: Float64: Used for more precise calculations
	var f64 float64 = 654.80876760860860 // Approximately -1.7e+308 to +1.7e+308
	var HP float64 = 10123456789012345
	fmt.Println("Float64: ", f64)

	// Low Precesion Vs Hight Precesion
	fmt.Println("LP", LP) // LP 1.0123457e+16
	fmt.Println("HP", HP) // HP 1.0123456789012344e+16
	
	// Section3: Boolean Types

	var on bool = true
	var off bool = false
	fmt.Println("is Active", on)
	fmt.Println("is Active", off)


	// Section4: Complex Data Type
	var CN1 complex128 = complex(5, 10)
	var CN2 complex64 = complex(2, 7)

	fmt.Println("CN1", CN1) // CN1 (5+10i)
	fmt.Println("CN2", CN2) // CN2 (2+7i)

	// Using a literal: c := 3 + 4i. An imaginary literal (like 4i) denotes a complex number with a zero real component. 
	c1 := 2 + 3i
    c2 := 4 + 5i
    fmt.Println("Addition:", c1 + c2)       // Output: Addition: (6+8i)
    fmt.Println("Multiplication:", c1 * c2) // Output: Multiplication: (-7+22i)
    fmt.Println("Real part:", real(c1 * c2))  // Output: Real part: -7
    fmt.Println("Imaginary part:", imag(c1 * c2)) // Output: Imaginary part: 22

	// Section5: String Data Type
	var name string = "Yassine"

	// formatting but not printing to the console
	var greating string = fmt.Sprintf("Hello... %s!", name)
    fmt.Println(greating)
	// formatting and printing to the console but no new line
	// fmt.Printf("Hello... %s!", name) // no new line

	// Bonus of using a short hand declaratin and assignments
	// (:=) : is the short variable declaration operator. It is a shorthand for declaring and initializing a variable in a single statement, with the compiler automatically inferring the variable's type based on the assigned value. 

	age := 32
	// age = "Yassine" // cannot use "Yassine" (untyped string constant) as int value in assignment
	fmt.Println("My Age is not", age)
}