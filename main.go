package main

import (
	"fmt"
)

func main() {
	// Create a Secret Number
	secret := getRandomNumber()
	// Looping for the last step
	for matching := false; !matching; {
		// Get Input Number
		guess := getInputNumber()
		//fmt.Println(secret, guess)
		// Make Comaprison
		matching = isMatching(secret, guess)
		fmt.Println(matching)
	}
}
