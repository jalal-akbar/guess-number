package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

func getInputNumber() int {
	var input int
	fmt.Print("Please Input Your Number : ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("failed to parse your input")
	} else {
		fmt.Println("You Guessed", input)
	}
	return input
}

func isMatching(secret, guess int) bool {
	if guess < secret {
		fmt.Println("Your Guess is too Smaller")
		return false
	} else if guess > secret {
		fmt.Println("Your Guess is  too Bigger")
		return false
	} else {
		fmt.Println("OK: Number Matching")
		return true
	}
}
