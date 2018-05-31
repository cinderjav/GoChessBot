package main

import "fmt"

func main() {
	board, turn := fenParser("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	chessGame := getChessGame(board, turn)
	executedMove := chessGame.executeMove()
	fmt.Println(executedMove)
}
