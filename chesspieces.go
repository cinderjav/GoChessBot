package main

import (
	"math"
)

//Need to handle pessant scenario and castling

type IChessPiece interface {
	canMove(move Move, board [8][8]string) bool
}

type ChessPiece struct {
	x, y, value int
	description string
}

type Pawn struct {
	ChessPiece
}
type Rook struct {
	ChessPiece
}
type Bishop struct {
	ChessPiece
}
type Knight struct {
	ChessPiece
}
type Queen struct {
	ChessPiece
}
type King struct {
	ChessPiece
}

func (pawn Pawn) canMove(move Move, board [8][8]string) bool {
	whitePawn := pawn.description == WhitePawn
	blackPawn := pawn.description == BlackPawn
	yDistance := math.Abs(float64(move.y) - float64(pawn.y))
	xDistance := math.Abs(float64(move.x) - float64(pawn.x))

	//zDistance := math.Abs(float64(move.x)-float64(pawn.x)) + math.Abs(float64(move.y)-float64(pawn.y))
	if move.chessPiece != nil && (move.y == pawn.y) {
		return false
	}
	if whitePawn && (move.x > pawn.x) {
		return false
	}
	if blackPawn && (move.x < pawn.x) {
		return false
	}

	if isMoveBlocked(move, pawn.x, pawn.y, board) {
		return false
	}

	if (whitePawn && pawn.x == 6) && (pawn.x-move.x <= 2) && (pawn.y == move.y) {
		return true
	}
	if (blackPawn && pawn.x == 1) && (move.x-pawn.x) <= 2 && (pawn.y == move.y) {
		return true
	}

	if xDistance == 1 && yDistance == 0 && move.chessPiece == nil {
		return true
	}

	if xDistance == 1 && yDistance == 1 && move.chessPiece != nil {
		return true
	}

	return false
}

func (rook Rook) canMove(move Move, board [8][8]string) bool {
	yDistance := math.Abs(float64(move.y) - float64(rook.y))
	xDistance := math.Abs(float64(move.x) - float64(rook.x))
	if yDistance > 0 && xDistance > 0 {
		return false
	}

	if isMoveBlocked(move, rook.x, rook.y, board) {
		return false
	}

	if (yDistance > 0 && xDistance == 0) || (xDistance > 0 && yDistance == 0) {
		return true
	}

	return false
}

func (bishop Bishop) canMove(move Move, board [8][8]string) bool {
	yDistance := math.Abs(float64(move.y) - float64(bishop.y))
	xDistance := math.Abs(float64(move.x) - float64(bishop.x))

	if xDistance != yDistance {
		return false
	}

	if isMoveBlocked(move, bishop.x, bishop.y, board) {
		return false
	}

	return true
}

func (knight Knight) canMove(move Move, board [8][8]string) bool {
	yDistance := math.Abs(float64(move.y) - float64(knight.y))
	xDistance := math.Abs(float64(move.x) - float64(knight.x))

	if (xDistance == 2 && yDistance == 1) || (yDistance == 2 && xDistance == 1) {
		return true
	}
	return false
}

func (queen Queen) canMove(move Move, board [8][8]string) bool {
	yDistance := math.Abs(float64(move.y) - float64(queen.y))
	xDistance := math.Abs(float64(move.x) - float64(queen.x))
	//not on diagonal and not adjacent
	if (xDistance != yDistance) && xDistance > 1 && yDistance > 1 {
		return false
	}

	if isMoveBlocked(move, queen.x, queen.y, board) {
		return false
	}
	return true
}

func (king King) canMove(move Move, board [8][8]string) bool {
	yDistance := math.Abs(float64(move.y) - float64(king.y))
	xDistance := math.Abs(float64(move.x) - float64(king.x))

	if xDistance > 1 || yDistance > 1 {
		return false
	}
	return true
}
