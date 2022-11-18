package main

import "fmt"

type example struct {
	flag    bool // 1 byte
	counter int16
	pi      float32
	// flag2 bool (now 7 bytes of padding, as next one should start at 8 multiple)
	// counter2 int64
	//if we want to optimize, then order from largest to smallest : in64, int32, bool
}
type bill struct {
	flag    bool // 1 byte
	counter int16
	pi      float32
}
type alice struct {
	flag    bool // 1 byte
	counter int16
	pi      float32
}

//also struct has to align based on largest value(here int64) so 8 byte alignment

//paddings and alignments,
//2 byte should end at 0, 2, 4, ...
//4 byte at 0, 4, 8, ...
//so flag at 0, then empty(1 byte padding), int(2 bytes), float(4 bytes)

func main() {
	//create value of type example
	var e1 example
	fmt.Printf("%+v\n", e1)

	// declare using struct literal construction
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14159, //add comma here too
	}
	fmt.Printf("Flag\n", e2.flag)

	var e3 = example{} //emmpty literal construction
	fmt.Printf("%+v\n", e3)

	// declare variable of anonymous type and init using struct literal
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi: 3.128,
	}
	fmt.Printf("%+v\n", e4)

	//explicit conversion
	var b bill
	var a alice
	b = bill(a)
	fmt.Println(a, b) 

	//implicit conversion
	// b = a 
	// fmt.Println(a, b) //compiler error

	//above in implicit example, a is named type, e4 is not named type below
	var b2 bill
	b2 = e4 //compiler allows
	fmt.Println(b2, e4)
}
