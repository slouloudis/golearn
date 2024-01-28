package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// return multiple things - the thing u care about and the error
func getGuess() (int, error) {
	var number int
	fmt.Scan(&number)
	// how to do template literals in go, goes in order
	// %s for string
	/// %d for digit

	if number < 1 || number > 10 {
		// create new error object
		// return in order of params.
		return 0, errors.New("must be a number between 1 and 10")
	}

	return number, nil
}

func main() {

	// get user input
	// var name string
	// fmt.Scan(&name)

	// have to use all vars 0- tell go we dont want the secnd var using _
	// // do the thing, check the error - pattern of kings
	// _, err := getGuess()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Good job!")

	number := rand.Intn(9) + 1

	var guess int
	var err error

	for number != guess {

		fmt.Println("Enter a number between 1 and 10")
		guess, err = getGuess()

		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Nice!")
}
