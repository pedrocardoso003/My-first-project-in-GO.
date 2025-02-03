package main

import (
	"fmt"
)

type board [4][4]int

func main() {
	fmt.Println("Hello, World!")
}

func createBoard() board {
	return board{}
}
