package main

// import "fmt"
// shared memory, pointers
func main() {

	count := 10

	println("count:\t value of [", count, "]\tAddr of[",&count, "]")

	increment(&count)

	println("count:\t value of [", count, "]\tAddr of[",&count, "]")
}

func increment(inc *int) { // * - pointer to int type variable, inc being unnamed

	// increment the value of count that pointer points to
	*inc++
	println("count:\t value of [", inc, "]\tAddr of[",&inc, "]\t value points to[",*inc,"]")
}