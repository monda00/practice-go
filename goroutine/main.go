package main

import (
	"fmt"
)

func Say(s string) {
	fmt.Println(s)
}

func main() {
	go Say("hello")
	go Say("world")
	go Say("!!")

	fmt.Scanln()
}
