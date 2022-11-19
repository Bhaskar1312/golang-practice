package main

const size = 1024 //run with 1(same addresses) and then 1024(different addresses)
func main() {
	s := "HELLO"
	stackcopy(&s, 0, [size]int{})
}

func stackcopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackcopy(s, c, a)
}

// os - 4M
// go co-routines or go-routine-2K for stack
// go is about integrity and then minimizing resources
// so go copies, all pointers on stack to another large stack(25% larger)
