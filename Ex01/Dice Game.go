package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? \n> ")

	name, _ := reader.ReadString('\n')

	// Print a greeting message
	fmt.Printf("Hello, %s!\n", name)
}
