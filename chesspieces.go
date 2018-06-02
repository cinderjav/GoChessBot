package main

import (
	"math"
)

//Need to handle pessant scenario and castling

type IChessPiece interface {
	canMove(move Move, board [8][8]string) bool
	getValue() int
	xLocation() int
	yLocation() int
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

func (pawn Pawn) getValue() int {
	return PawnScore
}

func (pawn Pawn) xLocation() int {
	return pawn.x
}
func (pawn Pawn) yLocation() int {
	return pawn.y
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

func (rook Rook) getValue() int {
	return RookScore
}

func (rook Rook) xLocation() int {
	return rook.x
}
func (rook Rook) yLocation() int {
	return rook.y
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

func (bishop Bishop) getValue() int {
	return BishopScore
}

func (bishop Bishop) xLocation() int {
	return bishop.x
}
func (bishop Bishop) yLocation() int {
	return bishop.y
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

func (knight Knight) getValue() int {
	return KnightScore
}

func (knight Knight) xLocation() int {
	return knight.x
}
func (knight Knight) yLocation() int {
	return knight.y
}

func (knight Knight) canMove(move Move, board [8][8]string) bool {
	yDistance := math.Abs(float64(move.y) - float64(knight.y))
	xDistance := math.Abs(float64(move.x) - float64(knight.x))

	if (xDistance == 2 && yDistance == 1) || (yDistance == 2 && xDistance == 1) {
		return true
	}
	return false
}

func (queen Queen) getValue() int {
	return QueenScore
}

func (queen Queen) xLocation() int {
	return queen.x
}
func (queen Queen) yLocation() int {
	return queen.y
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

func (king King) getValue() int {
	return KingScore
}

func (king King) xLocation() int {
	return king.x
}
func (king King) yLocation() int {
	return king.y
}

func (king King) canMove(move Move, board [8][8]string) bool {
	yDistance := math.Abs(float64(move.y) - float64(king.y))
	xDistance := math.Abs(float64(move.x) - float64(king.x))

	if xDistance > 1 || yDistance > 1 {
		return false
	}
	return true
}
