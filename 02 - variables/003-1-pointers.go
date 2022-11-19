package main

// import "fmt"
// pass by value
func main() {

	count := 10

	println("count:\t value of [", count, "]\tAddr of[",&count, "]")

	increment(count)

	println("count:\t value of [", count, "]\tAddr of[",&count, "]")
}

func increment(inc int) {
	inc++
	println("count:\t value of [", inc, "]\tAddr of[",&inc, "]")
}