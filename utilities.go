package main

import (
	"fmt"
	"math"
)

type Move struct {
	x, y, score int
	chessPiece  IChessPiece
}

func getChessGame(board [8][8]string, turn string) ChessGame {
	//any initialization modification will go here
	return ChessGame{board, turn}
}

func getAllAvailableMovesForTurn(pieces []IChessPiece, chessGame *ChessGame) map[IChessPiece][]Move {
	movableBoardSpots := chessGame.getMovableMoves()
	pieceMoves := getAvailablePieceMoves(movableBoardSpots, pieces, chessGame.board)
	return pieceMoves
}

func calculateBestMoveForPiece(piece IChessPiece, chessGame *ChessGame) Move {
	return Move{}
}

func calculateBestMoveForTurn(moves []Move) Move {
	return Move{}
}

func isWhitePiece(piece string) bool {
	if piece == EmptySpace {
		return false
	}
	if piece == WhiteBishop || piece == WhiteKing || piece == WhiteKnight ||
		piece == WhiteQueen || piece == WhiteRook || piece == WhitePawn {
		return true
	}
	return false
}

func isBlackPiece(piece string) bool {
	if piece == EmptySpace {
		return false
	}
	if piece == BlackBishop || piece == BlackKing || piece == BlackKnight ||
		piece == BlackQueen || piece == BlackRook || piece == BlackPawn {
		return true
	}
	return false
}

func getAvailablePieceMoves(moves []Move, pieces []IChessPiece, board [8][8]string) map[IChessPiece][]Move {
	var validMoves []Move
	var validMoveMapping = make(map[IChessPiece][]Move)
	//think we need to return mapping of piece to moves
	//println(len(moves))
	for _, move := range moves {
		for _, piece := range pieces {
			if piece.canMove(move, board) {
				validMoves = append(validMoves, move)
				//if I go with map, will remove break
				validMoveMapping[piece] = append(validMoveMapping[piece], move)
				//break
			}
		}
	}
	//println(len(validMoves))
	//return validMoves
	return validMoveMapping
}

func isMoveBlocked(move Move, x, y int, board [8][8]string) bool {
	//this method should not be called for knight
	yMovesBetween := math.Abs(float64(move.y)-float64(y)) - 1
	xMovesBetween := math.Abs(float64(move.x)-float64(x)) - 1
	//handles adjacent cells
	if xMovesBetween <= 0 && yMovesBetween <= 0 {
		return false
	}

	xDistance := 0
	shouldIncrementX := move.x > x
	var xMove int
	if shouldIncrementX {
		xMove = x + 1
	} else {
		xMove = x - 1
	}
	if yMovesBetween <= 0 {
		for float64(xDistance) < xMovesBetween {
			if shouldIncrementX {
				cell := board[xMove][y]
				if cell != EmptySpace {
					return true
				}
				xDistance++
				xMove++
			} else {
				cell := board[xMove][y]
				if cell != EmptySpace {
					return true
				}
				xDistance++
				xMove--
			}
		}
	}

	yDistance := 0
	shouldIncrementY := move.y > y
	var yMove int
	if shouldIncrementY {
		yMove = y + 1
	} else {
		yMove = y - 1
	}

	if xMovesBetween <= 0 {
		for float64(yDistance) < yMovesBetween {
			if shouldIncrementY {
				cell := board[x][yMove]
				if cell != EmptySpace {
					return true
				}
				yDistance++
				yMove++
			} else {
				cell := board[x][yMove]
				if cell != EmptySpace {
					return true
				}
				yDistance++
				yMove--
			}
		}
	}

	//diagonal check
	if xMovesBetween == yMovesBetween {
		diagDistance := 0
		xDistance := x + 1
		yDistance := y + 1
		//moving down and right
		if move.x > x && move.y > y {
			for float64(diagDistance) < yMovesBetween {
				piece := board[xDistance][yDistance]
				if piece != EmptySpace {
					return true
				}
				xDistance++
				yDistance++
				diagDistance++
			}
		}

		//moving up and left
		diagDistance = 0
		xDistance = x - 1
		yDistance = y - 1
		if move.x < x && move.y < y {
			for float64(diagDistance) < yMovesBetween {
				piece := board[xDistance][yDistance]
				if piece != EmptySpace {
					return true
				}
				xDistance--
				yDistance--
				diagDistance++
			}
		}

		//moving up and right
		diagDistance = 0
		xDistance = x - 1
		yDistance = y + 1
		if move.x < x && move.y > y {
			for float64(diagDistance) < yMovesBetween {
				piece := board[xDistance][yDistance]
				if piece != EmptySpace {
					return true
				}
				xDistance--
				yDistance++
				diagDistance++
			}
		}

		//moving down and to the left
		diagDistance = 0
		xDistance = x + 1
		yDistance = y - 1
		if move.x > x && move.y < y {
			for float64(diagDistance) < yMovesBetween {
				piece := board[xDistance][yDistance]
				if piece != EmptySpace {
					return true
				}
				xDistance++
				yDistance--
				diagDistance++
			}
		}
		//4 posible directions here
	}
	return false

}

func analyzeMoves(moveMapping map[IChessPiece][]Move, chessGame *ChessGame, level int, score int, originalTurn string, prevEaten IChessPiece) {
	for piece, moves := range moveMapping {
		for index, move := range moves {
			score := analyzeMove(piece, move, chessGame.board, chessGame.playerTurn, level, score, originalTurn, prevEaten)
			moves := moveMapping[piece]
			moves[index].score = score
			// go func(piece IChessPiece, move Move, board [8][8]string, turn string, level int, score int, originalTurn string, moveMapping map[IChessPiece][]Move) {
			// 	scores <- analyzeMove(piece, move, board, turn, level, score, originalTurn)
			// 	moves := moveMapping[piece]
			// 	newScore := <-scores
			// 	moves[index].score = newScore
			// }(piece, move, chessGame.board, chessGame.playerTurn, level, score, originalTurn, moveMapping)
		}
	}
}

func analyzeMove(piece IChessPiece, move Move, board [8][8]string, turn string, level int, score int, originalTurn string, lastEaten IChessPiece) int {
	//possible issue not passing along score
	//implement shouldprune and getindividual move score
	//reason about the implementation
	//score needs to take into account my color
	prune, value := shouldPrune(piece, move, board, turn == originalTurn, level, score, lastEaten)
	if prune {
		return score + (value * (MaxRecursiveLevel - level))
	}
	scoreMove := getIndividualMoveScore(piece, move, board, turn)
	if turn == originalTurn {
		score += (scoreMove * (MaxRecursiveLevel - level))
	} else {
		score -= (scoreMove * (MaxRecursiveLevel - level))
	}
	if level == MaxRecursiveLevel {
		return score
	}

	newBoard := makeBoardMove(piece, move, board)
	newTurn := getNextPlayerTurn(turn)
	newChessGame := ChessGame{newBoard, newTurn}
	pieces := newChessGame.getPiecesForTurn()
	newMovesMapping := getAllAvailableMovesForTurn(pieces, &newChessGame)
	if move.chessPiece != nil {
		lastEaten = move.chessPiece
	}
	analyzeMoves(newMovesMapping, &newChessGame, level+1, score, originalTurn, lastEaten)
	highScore, _, _ := getHighestMoveScoreFromMap(newMovesMapping)
	return highScore
}

func shouldPrune(piece IChessPiece, move Move, board [8][8]string, initialTurn bool, level int, score int, lastEaten IChessPiece) (bool, int) {
	//this function will can return high value end game
	//need to make sure smallers turns are favored, pass in recursive level
	if move.chessPiece != nil {
		if move.chessPiece.getValue() == KingScore && initialTurn {
			return true, KingScore
		}

		if move.chessPiece.getValue() == KingScore && !initialTurn && level != 1 {
			return true, -KingScore
		}

		if move.chessPiece.getValue() == KingScore && !initialTurn && level == 1 {
			return true, -1000000
		}

		if lastEaten == nil && !initialTurn && level == 1 {
			return true, -1000000
		}

		if lastEaten != nil {
			if move.chessPiece.getValue() > lastEaten.getValue() && !initialTurn && level == 1 {
				return true, -1000000
			}
		}
	}

	return false, 0
}

func getIndividualMoveScore(piece IChessPiece, move Move, board [8][8]string, turn string) int {
	//should try preventing dumb moves? Not sure score might handle it
	// var nextTurn = getNextPlayerTurn(turn)
	// var newChess = ChessGame{board, nextTurn}
	// var opponentPieces = newChess.getPiecesForTurn()
	// movesMappingOpponents := getAllAvailableMovesForTurn(opponentPieces, &newChess)

	// var targetable =
	if move.chessPiece != nil {
		//need to consider if I get eaten but opponent score will take that into account
		return move.chessPiece.getValue()
	}

	return 0
}

func makeBoardMove(piece IChessPiece, move Move, board [8][8]string) [8][8]string {
	pieceString := board[piece.xLocation()][piece.yLocation()]
	board[piece.xLocation()][piece.yLocation()] = ""
	board[move.x][move.y] = pieceString
	return board
}

func getNextPlayerTurn(currentTurn string) string {
	if currentTurn == WhiteTurn {
		return BlackTurn
	}
	return WhiteTurn
}

func getHighestMoveScoreFromMap(moveMapping map[IChessPiece][]Move) (int, IChessPiece, Move) {
	score := -10000000
	var topMove Move
	var topPiece IChessPiece
	sumScore := 0
	sumCount := 0
	for piece, moves := range moveMapping {
		for _, move := range moves {
			sumScore += move.score
			sumCount++
			if move.score > score {
				score = move.score
				topMove = move
				topPiece = piece
			}
		}
	}
	return sumScore, topPiece, topMove
}

func getPieceString(x, y int, board [8][8]string) string {
	return board[x][y]
}

func translateMove(piece IChessPiece, move Move, board [8][8]string) string {
	var pieceNotation string
	var moveNotation string
	fmt.Println(piece, move)
	//send back algebraic and basic just in case
	// pieceString := getPieceString(piece.xLocation(), piece.yLocation(), board)
	// //moveString := getPieceString(move.x, move.y, board)
	// if pieceString == "P" {
	// 	pieceString = ""
	// }
	//pieceNotation += strings.ToUpper(pieceString)
	switch piece.yLocation() {
	case 0:
		pieceNotation += "a"
	case 1:
		pieceNotation += "b"
	case 2:
		pieceNotation += "c"
	case 3:
		pieceNotation += "d"
	case 4:
		pieceNotation += "e"
	case 5:
		pieceNotation += "f"
	case 6:
		pieceNotation += "g"
	case 7:
		pieceNotation += "h"
	}

	switch piece.xLocation() {
	case 0:
		pieceNotation += "8"
	case 1:
		pieceNotation += "7"
	case 2:
		pieceNotation += "6"
	case 3:
		pieceNotation += "5"
	case 4:
		pieceNotation += "4"
	case 5:
		pieceNotation += "3"
	case 6:
		pieceNotation += "2"
	case 7:
		pieceNotation += "1"
	}

	switch move.y {
	case 0:
		moveNotation += "a"
	case 1:
		moveNotation += "b"
	case 2:
		moveNotation += "c"
	case 3:
		moveNotation += "d"
	case 4:
		moveNotation += "e"
	case 5:
		moveNotation += "f"
	case 6:
		moveNotation += "g"
	case 7:
		moveNotation += "h"
	}

	switch move.x {
	case 0:
		moveNotation += "8"
	case 1:
		moveNotation += "7"
	case 2:
		moveNotation += "6"
	case 3:
		moveNotation += "5"
	case 4:
		moveNotation += "4"
	case 5:
		moveNotation += "3"
	case 6:
		moveNotation += "2"
	case 7:
		moveNotation += "1"
	}

	//need to support the other symbols
	var sep = "-"
	// if move.chessPiece != nil {
	// 	sep = "x"
	// }
	return pieceNotation + sep + moveNotation

}
