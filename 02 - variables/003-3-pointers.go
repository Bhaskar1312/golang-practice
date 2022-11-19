package main

type user struct {
	name string
	email string
}

// escape analysis - determines value cannot be on stack, allocates on heap
// stack is self-healing, heap requires garbage collection
func main() {
	u1 := createUserV1()
	u2 := createUserV2()
	u3 := createUserV3()

	println("V1", &u1)
	println("V2", &u2)
	println("V3", &u3)
}

func createUserV1() user {
	u := user {
		name: "Bill",
		email: "bill@ardanlabs.com",
	}
	println("V1", &u)
	return u
}

//value semantic construction, then pointer
func createUserV2() *user {
	u := user {
		name: "Bill",
		email: "bill@ardanlabs.com",
	}
	println("V2", &u)
	return &u
}

//pointer semantic construction; only use at return or func input param, not like here
func createUserV3() *user {
	u := &user {
		name: "Bill",
		email: "bill@ardanlabs.com",
	}
	println("V3", &u)
	return u
}

//  go build -gcflags "-m -m" // for profiling, why go allocates into heap etc..