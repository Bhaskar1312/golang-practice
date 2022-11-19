package main

import (
	"fmt"
	"time"
)

const (
	maxInt = 9223372036854775807 //64-bit arch, max int 
	bigger = 922337203685477580712345
)

type Duration int64
const (
	Nanosecond Duration = 1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second = 1000 * Millisecond
	Minute = 60 * Second
	Hour = 60 * Minute
)

func main() {

	// compiler can implicitly perform conversions of untyped constants(unlike varibles)
	// untyped constants
	const ui = 12345   // kind: integer
	const uf = 3.1415  //kind: floating-point


	const ti int = 12345       
	const tf float64 = 2.71828

	// const myUint8 unit8 = 1000 //overflows

	//var is of type float64
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	println(answer)

	const zero = 1/3 // KindInt(1)/KindInt(3)


	// max values
	fmt.Println("Will Compile")


	now := time.Now()
	
	literal := now.Add(-5) //subtact 5 nanoseconds

	const timeout = 5 * time.Second
	constant := now.Add(-timeout)

	// minusFive := int64(-5)
	// variable := now.Add(minusFive)
	// .go:55:22: cannot use minusFive (variable of type int64) as type time.Duration in argument to now.Add

	fmt.Println(now)
	fmt.Printf("Now        : %v\n", now)
	fmt.Printf("Literal    : %v\n", literal)
	fmt.Printf("Constant   : %v\n", constant)
	// fmt.Println("Variable   : %v \n", variable)


	const ( // blocks for constants, variables, import
		A1 = iota //0 : start at 0
		B1 = iota //1 
		C1 = iota //2
	)
	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota //0 : start at 0
		B2
		C2
	)
	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1: start at 0 + 1
		B3
		C3
	)
	fmt.Println("3:", A3, B3, C3)

	const (
		Ldate      = 1<< iota // 1 : Shift 1 to left 0. 0000 0001
		Ltime                 // 2 : Shift 1 to left 1. 0000 0010
		Lmicroseconds         // 4 : Shift 1 to left 2. 0000 0100
		Llongfile             // 8 
		Lshortfile            //16
		LUTC                  //32
	)
	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}

// consts - 256 bits of precision - unnamed consts - Kind Promotion
// consts only exist at compile time