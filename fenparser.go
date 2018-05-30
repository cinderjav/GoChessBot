package main

import (
	"log"
	"strconv"
)

func fenParser(fen string) [8][8]string {
	var chessboard [8][8]string

	row := 0
	col := 0
	terminate := false
	for _, char := range fen {
		//fmt.Println(char, string(char), row, col)
		isLetter := char >= 65
		isNumber := char >= 49 && char <= 56
		isSlash := char == 47

		switch {
		case isLetter:
			chessboard[row][col] = string(char)
			col++
		case isNumber:
			data := string(char)
			conversion, err := strconv.Atoi(data)
			if err != nil {
				log.Fatal(err)
			}
			col += conversion
		case isSlash:
			row++
			col = 0
		default:
			terminate = true
		}

		if terminate {
			break
		}
	}
	return chessboard

	/*Rank: left to right on string, picture chess board top to bottom. White uppercase, black lowercase
	8 refers to empty row
	4P3 example means, 4 spaces pawn 3 spaces
	1 means empty square
	*/
	//Turn: w means it is whites turn
	//Castling: - is used if castling not possible KQ refer to white king queen
	//passant square
	//half move clock
	//total round count, increments after black moves
}
