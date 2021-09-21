package main

import "fmt"

func main() {
	score, valid := 2, true

	if score > 3 && valid {
		fmt.Printf("good\n")
	} else if score == 3 {
		fmt.Printf("on the edge\n")
	} else if score == 2 {
		fmt.Printf("meh...\n")
	} else {
		fmt.Printf("low\n")
	}
}