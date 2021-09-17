package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "kshitij"
	name2 := "İnanç"

	// len returns number of bytes in a string not actual number of characters
	fmt.Println("Length of", name, "with len:", len(name))
	fmt.Println("Length of", name2, "with len:", len(name2))

	/*
	returns actual number of characters in a string,
	a rune can represent English and Non-English characters as well
	*/
	fmt.Println("Length of", name, "with RuneCountInString:", utf8.RuneCountInString(name))
	fmt.Println("Length of", name2, "with RuneCountInString:", utf8.RuneCountInString(name2))
}
