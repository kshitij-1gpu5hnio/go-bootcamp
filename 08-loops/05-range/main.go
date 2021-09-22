package main

import (
	"fmt"
	"os"
)

func main() {
	// for i, v := range os.Args {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(v)
	// }
	
	for _, v := range os.Args[1:] {
		fmt.Println(v)
	}
}