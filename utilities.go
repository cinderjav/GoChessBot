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
	println(len(moves))
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
	println(len(validMoves))
	fmt.Println(validMoveMapping)
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
		yMove = x + 1
	} else {
		yMove = x - 1
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
