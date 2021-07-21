package main

import (
	"fmt"
)

const (
	first  = iota
	second = iota
)

const (
	third = iota
	fourth
)

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
	fmt.Println(float32(co) + 1.2)

	fmt.Println(first, second, third, fourth)

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr[:]
	arr[1] = 18
	slice[2] = 79
	fmt.Println(arr, slice)

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	slice1 = append(slice1, 4, 5)
	fmt.Println(slice1)

	slice2 := slice1[1:]
	slice3 := slice1[:2]
	slice4 := slice1[1:2]

	fmt.Println(slice2, slice3, slice4)

	m := map[string]int{"a": 1}
	fmt.Println(m)
	fmt.Println(m["a"])

	m["a"] = 9
	fmt.Println(m["a"])

	delete(m, "a")
	fmt.Println(m)

}
