package main

import (
	"math"
)

type IChessPiece interface {
	canMove(move Move) bool
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

func (pawn Pawn) canMove(move Move) bool {
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

func (rook Rook) canMove(move Move) bool {
	return false
}

func (bishop Bishop) canMove(move Move) bool {
	return false
}

func (knight Knight) canMove(move Move) bool {
	return false
}

func (queen Queen) canMove(move Move) bool {
	return false
}

func (king King) canMove(move Move) bool {
	return false
}
