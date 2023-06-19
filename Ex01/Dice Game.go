package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? \n> ")

	name, _ := reader.ReadString('\n')

	// Print a greeting message
	fmt.Printf("Hello, %s!\n", name)
	
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate two random numbers from 1 to 6
	die1 := rand.Intn(6) + 1
	die2 := rand.Intn(6) + 1
	total := die1 + die2

	fmt.Println("Rolling dice...")
	fmt.Println("Die 1:", die1)
	fmt.Println("Die 2:", die2)
	fmt.Println("Total value:", total)

	if total > 7 {
		fmt.Println("You won")
	} else {
		fmt.Println("You lost")
	}
}
