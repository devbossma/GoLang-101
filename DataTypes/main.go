package main

import "fmt"

func main(){

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

	// A: Unsigned Integers

	var u uint // uint: Platform-dependent (32 or 64 bits).
	var u8 uint8 // uint8 (alias byte): 0 to 255
	var u16 uint16 // uint16: 0 to 65,535
	var u32 uint32 // uint32: 0 to 4,294,967,295
	var u64 uint64 // uint64: 0 to 18,446,744,073,709,551,615
	var uintlarge uintptr    //: An unsigned integer type large enough to store the bit pattern of any pointer. 
	// Value Assignment
	u 	= 255
	u8  = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 1844674473709551615

	fmt.Println("Signed Integers", i, i8, i16, i32, i64)
	fmt.Println("USigned Integers", u, u8, u16, u32, u64)
	


}