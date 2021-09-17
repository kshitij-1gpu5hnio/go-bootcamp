package main

import "fmt"

func main() {
	// iota is constant generator which generates ever increasing numbers
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	// EST = -5 MST = -7 PST = -8
	// blank identifier with iota
 	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)
}