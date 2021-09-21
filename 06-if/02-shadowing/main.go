package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n int
		err error
	)

	if args := os.Args; len(args) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(args[1]); err != nil {
		fmt.Printf("Cannot convert %q.\n", args[1])
	} else {
		fmt.Printf("%s * 2 = %d\n", args[1], n*2)
	}

	fmt.Printf("n is %d. You've been shadowed.", n)
}
