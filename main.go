package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error: arguments not found")
		return
	}
	name := os.Args[1]
	fmt.Println("hello,", name)
}
