package main

import ( "fmt"; "path" )

func main() {
	// var dir, file string
	// var file string

	// dir, file = path.Split("css/main.css")
	// _, file = path.Split("css/main.css") // Using blank identifier to discard dir
	_, file := path.Split("css/main.css") // Using short declaration

	// fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}
