package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Tejas"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(3), imag(4)
	fmt.Println(r, im)

	var lastName *string = new(string)
	*lastName = "Jog"
	fmt.Println(lastName)
	fmt.Println(*lastName)

	middleName := "Prakash"
	fmt.Println(middleName)
	ptr := &middleName
	fmt.Println(ptr, *ptr)
	middleName = "Prakash1"
	fmt.Println(ptr, *ptr)

	const pi = 3.14 // has to be initialized when declared and the value should be determinable at compile time
	fmt.Println(pi)

	const co int = 3
	fmt.Println(co + 3)
	//do something
	fmt.Println(co + 1.2)

}
