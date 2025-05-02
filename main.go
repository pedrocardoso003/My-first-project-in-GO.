package main

import (
	"fmt"
	"math/rand"
)

//TODO: Implementar logica de movimento
/*
	numeros para uma direcao
	se dois numeros iguais coliderem eles somam-se
	apos cada movimento adicionar um novo random tile
*/

var board [4][4]int

func main() {
	fmt.Println("Hello, World!")
	initializeBoard()
	fmt.Println("Initial Board:")
	fmt.Println(board)
}

func initializeBoard() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}
	for i := 0; i < 2; i++ {
		addRandomTile()
	}
}

func addRandomTile() {
	//rand.intn(n) returns a random number between 0 and n

	randomNumber := rand.Intn(2) //only 2 or 4

	if randomNumber == 0 {
		randomNumber = 2
	} else {
		randomNumber = 4
	}

	//make sure the random number is added to an empty cell
	for {
		if board[rand.Intn(4)][rand.Intn(4)] == 0 {
			board[rand.Intn(4)][rand.Intn(4)] = randomNumber
			break
		}
	}
}
