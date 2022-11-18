package main

import "fmt"

func main() {
	//default set to zero value
	var a int 
	var b string //use var to get default 0 value (empty string)
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a);
	fmt.Printf("var b string \t %T [%v]\n", b, b);
	fmt.Printf("var c float64 \t %T [%v]\n", c, c);
	fmt.Printf("var d bool \t %T [%v]\n", d, d);



	//short variable declaration
	aa := 0
	bb := "hello"
	cc := 3.14159
	dd := true
	fmt.Printf("var aa int \t %T [%v]\n", aa, aa);
	fmt.Printf("var bb string \t %T [%v]\n", bb, bb);
	fmt.Printf("var cc float64 \t %T [%v]\n", cc, cc);
	fmt.Printf("var dd bool \t %T [%v]\n", dd, dd);

	aaa := int32(10)
	bbb := string("hello")
	ccc := float64(1.23344)
	
}