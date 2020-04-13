package main

import (
	"fmt"
	"math/rand" //generates random number
	"time"      //
)

func main() {
	fmt.Println("Hi!")
	fmt.Println("Guess the number between 0-10!") //println adds new line vs print which does not

	source := rand.NewSource(time.Now().UnixNano()) // Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run
	number := rand.New(source)
	randomNumber := number.Intn(15) //setting limit 0 - whatever you set greater than 0 returns int so sources match int and int not int and rand*

	var userGuess int

	fmt.Printf("you guessed %d\n", userGuess) // println printed out "you guess %d" after, added\n to get rid of % sign

	for { // only looping in go is a for loop
		fmt.Println("Input your guess:")
		fmt.Scanln(&userGuess) //do & to pass memory to Scan or it just came out to nothing ... Scan line takes user input
		if userGuess > randomNumber {
			fmt.Println("Too big please try guessing again!")
		} else if userGuess < randomNumber {
			fmt.Println("Too small please try guessing again!")
		} else {
			fmt.Println("you won!")
			break // end looping
		}
	}
}
