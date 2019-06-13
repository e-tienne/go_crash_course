package main

import (
	"fmt"
)

func main() {

	// int
	// bool
	// byte
	// rune
	// float32, float64
	// complex64, complex128

	// var name = "Etienne"
	var age = 36
	const isCool = true

	// isCool = false #we can't assign false to isCool because we defined it as a const

	// shorthand
	// name := "Etienne"

	name, email := "Etienne", "esimard@redhat.com"
	size := 1.3

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)

}
