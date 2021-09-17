package main

import "fmt"

func main() {
	var s string
	s = "how are you?"
	s = `how are you?`

	fmt.Println(s)

	s = "<html>\n\t<body>\"hello\"</body>\n</html>"
	fmt.Println(s)

	s = `
<html>
	<bode>"hello"</body>
</html>`
	fmt.Println(s)

	fmt.Println("c:\\my\\dir\\file")  // String literal (Single line, processed)
	fmt.Println(`c:\my\dir\file`) // Raw string literal (Multi line unprocessed)
}