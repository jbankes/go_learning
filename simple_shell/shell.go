package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("-------------------")

	// main loop
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert crlf to lf
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hello, Justin")
		}
	}
}
