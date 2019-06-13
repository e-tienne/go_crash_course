package main

import (
	"fmt"
)

func main() {

	// // Arrays - need to be a fixed lenght and typed
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign values in 1 step / 1 line instead of 2 steps / 3 lines
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	// Slice example - not a fixed lenght

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
