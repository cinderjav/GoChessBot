package main

import (
	"log"
	"strconv"
	"strings"
)

func fenParser(fen string) ([8][8]string, string) {
	var chessboard [8][8]string
	row := 0
	col := 0
	fenArray := strings.Split(fen, " ")
	for _, char := range fenArray[0] {
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
		}
	}
	colorTurn := fenArray[1]
	return chessboard, colorTurn

	//Castling: - is used if castling not possible KQ refer to white king queen
	//passant square
	//half move clock
	//total round count, increments after black moves
}
