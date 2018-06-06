package main

import (
	"math"
)

//Need to handle pessant scenario and castling and promotions

type IChessPiece interface {
	canMove(move Move, board [8][8]string) bool
	getValue() int
	xLocation() int
	yLocation() int
	getCopy(x, y int, turn string) IChessPiece
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

func (pawn Pawn) getCopy(x, y int, turn string) IChessPiece {
	if turn == WhiteTurn {
		return Pawn{ChessPiece{x, y, PawnScore, WhitePawn}}
	}
	return Pawn{ChessPiece{x, y, PawnScore, BlackPawn}}
}

func (pawn Pawn) canMove(move Move, board [8][8]string) bool {
	whitePawn := pawn.description == WhitePawn
	blackPawn := pawn.description == BlackPawn
	yDistance := math.Abs(float64(move.y) - float64(pawn.y))
	xDistance := math.Abs(float64(move.x) - float64(pawn.x))

	if move.x == pawn.x && move.y == pawn.y {
		return false
	}
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

func (rook Rook) getCopy(x, y int, turn string) IChessPiece {
	if turn == WhiteTurn {
		return Rook{ChessPiece{x, y, RookScore, WhiteRook}}
	}
	return Rook{ChessPiece{x, y, RookScore, BlackRook}}
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

	if move.x == rook.x && move.y == rook.y {
		return false
	}

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

func (bishop Bishop) getCopy(x, y int, turn string) IChessPiece {
	if turn == WhiteTurn {
		return Bishop{ChessPiece{x, y, BishopScore, WhiteBishop}}
	}
	return Bishop{ChessPiece{x, y, BishopScore, BlackBishop}}
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

	if move.x == bishop.x && move.y == bishop.y {
		return false
	}

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

func (knight Knight) getCopy(x, y int, turn string) IChessPiece {
	if turn == WhiteTurn {
		return Knight{ChessPiece{x, y, KnightScore, WhiteKnight}}
	}
	return Knight{ChessPiece{x, y, KnightScore, BlackKnight}}
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

	if move.x == knight.x && move.y == knight.y {
		return false
	}

	if (xDistance == 2 && yDistance == 1) || (yDistance == 2 && xDistance == 1) {
		return true
	}
	return false
}

func (queen Queen) getValue() int {
	return QueenScore
}

func (queen Queen) getCopy(x, y int, turn string) IChessPiece {
	if turn == WhiteTurn {
		return Queen{ChessPiece{x, y, QueenScore, WhiteQueen}}
	}
	return Queen{ChessPiece{x, y, QueenScore, BlackQueen}}
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

	if move.x == queen.x && move.y == queen.y {
		return false
	}
	//not on diagonal and not adjacent
	if (xDistance != yDistance) && xDistance >= 1 && yDistance >= 1 {
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

func (king King) getCopy(x, y int, turn string) IChessPiece {
	if turn == WhiteTurn {
		return King{ChessPiece{x, y, KingScore, WhiteKing}}
	}
	return King{ChessPiece{x, y, KingScore, BlackKing}}
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

	if move.x == king.x && move.y == king.y {
		return false
	}

	if xDistance > 1 || yDistance > 1 {
		return false
	}
	return true
}
