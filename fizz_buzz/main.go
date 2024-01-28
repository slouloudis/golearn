package main

import "fmt"

// language formates itself.

// no unioin types
func isFizz(i int) bool {
	return i%3 == 0
}

func isBuzz(i int) bool {
	return i%5 == 0
}

func main() {
	// lets have fun
	fmt.Println("This is how string")

	for i := 0; i < 20; i++ {
		//
		if isFizz(i) && isBuzz(i) {
			fmt.Println("FizzBuzz")
		} else if isFizz(i) {
			fmt.Println("Fizz")
		} else if isBuzz(i) {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
