package main

import (
	"fmt"
)

func main() {

	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)

	}

	// Fizz buzz (Loop through 100, for every multiple of 3, output the word fizz, multiple of 5, buzz and 15 fizzbuzz, else output the iteration)

	for i := 1; i <= 100; i++ {

		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
