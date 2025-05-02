package main

import (
	"fmt"
	"math/rand"
)

//TODO: Implementar logica de movimento
/*
	apos cada movimento adicionar um novo random tile
*/

//TODO: Implementar interface e logica de jogo
/*
   se o numero 2048 aparecer, o jogo acaba (vitoria)
   se nao houver mais movimentos possiveis, o jogo acaba (derrota)
*/

var board [4][4]int

func main() {
	fmt.Println("Hello, World!")
	initializeBoard()
	fmt.Println("Initial Board:")
	fmt.Println(board)
	moveRight()
	fmt.Println("Board after moving right:")
	fmt.Println(board)
	moveLeft()
	fmt.Println("Board after moving left:")
	fmt.Println(board)
	moveUp()
	fmt.Println("Board after moving up:")
	fmt.Println(board)
	moveDown()
	fmt.Println("Board after moving down:")
	fmt.Println(board)
	/*test sum tiles
	board[0][2] = 2
	board[0][3] = 2
	fmt.Println("Board before moving right:")
	fmt.Println(board)
	moveRight()
	fmt.Println("Board after moving right:")
	fmt.Println(board)*/

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
	//rand.Intn(n) returns a random number between 0 and n

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

func moveRight() {
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- {
			if board[i][j] == 0 && board[i][j-1] != 0 {
				board[i][j] = board[i][j-1]
				board[i][j-1] = 0
			} else if board[i][j] == board[i][j-1] {
				board[i][j] *= 2
				board[i][j-1] = 0
			}
		}
	}
}

func moveLeft() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 && board[i][j+1] != 0 {
				board[i][j] = board[i][j+1]
				board[i][j+1] = 0
			} else if board[i][j] == board[i][j+1] {
				board[i][j] *= 2
				board[i][j+1] = 0
			}
		}
	}
}

func moveUp() {
	for j := 0; j < 4; j++ {
		for i := 0; i < 3; i++ {
			if board[i][j] == 0 && board[i+1][j] != 0 {
				board[i][j] = board[i+1][j]
				board[i+1][j] = 0
			} else if board[i][j] == board[i+1][j] {
				board[i][j] *= 2
				board[i+1][j] = 0
			}
		}
	}
}

func moveDown() {
	for j := 0; j < 4; j++ {
		for i := 3; i > 0; i-- {
			if board[i][j] == 0 && board[i-1][j] != 0 {
				board[i][j] = board[i-1][j]
				board[i-1][j] = 0
			} else if board[i][j] == board[i-1][j] {
				board[i][j] *= 2
				board[i-1][j] = 0
			}
		}
	}
}
