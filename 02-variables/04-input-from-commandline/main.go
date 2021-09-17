package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("Argument 1:", os.Args[1])
	fmt.Println("Argument 2:", os.Args[2])
	fmt.Println("Argument 3:", os.Args[3])

	fmt.Println("Number of items in os.Args:", len(os.Args))
}
