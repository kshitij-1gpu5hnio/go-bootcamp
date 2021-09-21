package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}


	i := 0

	// Using boolean express in case condition
	switch {
	case i > 0:
		fmt.Println("Postive")
	case i < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}