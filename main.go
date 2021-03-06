package main

import (
	"errors"
	"fmt"

	"github.com/tejasjog/webservice/models"
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
	var ii int
	ii = 42
	fmt.Println(ii)

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

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u1 user
	u1.ID = 1
	u1.FirstName = "Tejas"
	u1.LastName = "Jog"
	fmt.Println(u1)
	fmt.Println(u1.FirstName)

	u2 := user{ID: 1,
		FirstName: "Tejas",
		LastName:  "Jog",
	}
	fmt.Println(u2)

	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)

	for i := 0; i < 5; i++ {
		println(i)
		if i == 4 {
			break
		}
		if i == 3 {
			continue
		}
		println("continuing ... ")
	}

	var j int
	for {
		if j == 5 {
			break
		}
		println(j)
		j++
	}

	slice5 := []int{1, 2, 3}
	for k, val := range slice5 {
		println(k, val)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	for k, v := range wellKnownPorts {
		println(k, v)
	}

	for k := range wellKnownPorts {
		println(k)
	}

	for _, v := range wellKnownPorts {
		println(v)
	}

	user1 := user{
		ID:        1,
		FirstName: "Tejas",
		LastName:  "Jog",
	}

	user2 := user{
		ID:        2,
		FirstName: "tejas",
		LastName:  "jog",
	}

	if user1 == user2 {
		println("same user")
	} else if user1.FirstName == user2.FirstName {
		println("similar user")
	} else {
		println("not same user")
	}

	//port := 3000

	//_, err := startWebServer(port, 2)
	//fmt.Println(err)

	//controllers.RegisterController()
	//http.ListenAndServe(":3000", nil)

}

func startWebServer(port, numberOfRertries int) (bool, error) {
	fmt.Println("Starting server")
	//panic("Something bad just happened")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of rertries", numberOfRertries)
	return false, errors.New("Something went wrong")
}
