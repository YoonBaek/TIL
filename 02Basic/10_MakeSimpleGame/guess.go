// Author : Github Yoon Baek
// We are gonna make random number guessing game.
// Grammar in here : For Loop
// Golang has  only one loop.

package main

import (
	"bufio"
	. "fmt"
	"log"
	"math/rand" // we import by path/packagename and use by packagename
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := rand.Intn(100) + 1
	// Int'n' means we can set range by setting maximum limit n
	Println("randInt before setting seed :", target)
	// You can see "target" should change randomly, but it doesn't

	// make time value as int type.
	sec := time.Now().Unix()
	Println("Set seed by Now :", sec)
	// Set random seed.
	rand.Seed(sec)
	// Check target
	target = rand.Intn(100) + 1
	Println("randInt after setting seed :", target)

	// Real Guessing game
	// Reset Seed.
	time.Sleep(1)
	sec = time.Now().Unix()
	rand.Seed(sec)
	// make rand int
	target = rand.Intn(2000) + 1
	Println("I've chosen a random number between 1 and 2000.")
	Println("Can you guess it?")

	//Get inputs
	/*
		Buffer has bigger memory, and is faster than scanner.
		 if size of input is pretty big.
	*/
	reader := bufio.NewReader(os.Stdin)

	/*
		for loop also has a block.
		So if Varaiables declared inside the for block,
		It's scope is inside for block.
	*/
	// For Loop
	// if guesses become 0, exit the loop.
	var isSuccess bool = false
	for guesses := 10; guesses >= 1; guesses-- {
		Println("You have", guesses, "guesses left.")
		Println("Make a guess :")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
		}
		input = strings.TrimSpace(input)
		// Use strconv.Atoi. It converts string to int. A means ASCIII
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal()
		}
		if guess > target {
			Println("Your guess was", guess, "Too HIGH")
		} else if guess < target {
			Println("Your guess was", guess, "Too LOW")
		} else {
			isSuccess = true
			Println("Congrats! You guessed the answer.")
			break
			/*
				continue : skip to the next loop.
				break : end the loop
			*/
		}
	}
	// if failed, tell user the answer.
	if !isSuccess {
		Println("Sorry, No chances are left. The answer was :", target)
	}
}
