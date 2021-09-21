package main

import "fmt"

func main() {
	var brand = "Google"

	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand)

	var (
		planet = "venus"
		distance = 261
		orbital = 224.701
		hasLife = false
	)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Orbital Period: %.3f days\n", orbital)
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
	fmt.Printf("Does %s has life? %t\n", planet, hasLife)
	

	fmt.Printf(
		"%v is %v away. Think %[2]v kms! %[1]v OMG!\n",
		planet, distance,
	)
}
