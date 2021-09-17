package main

func main() {
	/*
		When to use normal declaration:
			1. If you don't know the initial value
			score := 0 // Don't
			var score int // Automatically initialized to value 0

			2. When you need a package scoped variable

			3. When you want to group variables together for 
			greater readability
			var (
				// Related:
				video string

				// Closely related:
				duration int
				current int
			)
		
		When to use short declaration:
			1. If you know the initial value
			var width, height = 100, 50// Don't
			widht, height := 100, 50 // Preferred way

			2. Use short declaration to keep code concise
			and easy to read

			3. For redeclaration
			// Don't
			width = 50
			color := red

			// Preferred way
			width, color := 50, "red"
			
	*/
}
