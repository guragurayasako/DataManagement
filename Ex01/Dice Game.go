package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate two random numbers from 1 to 6
	die1 := rand.Intn(6) + 1
	die2 := rand.Intn(6) + 1

	fmt.Println("Rolling dice...")
	fmt.Println("Die 1:", die1)
	fmt.Println("Die 2:", die2)
	fmt.Println("Total value:", die1+die2)
}
