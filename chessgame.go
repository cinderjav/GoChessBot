package main

import (
	"fmt"
	"strings"
)

type ChessGame struct {
	board      [8][8]string
	playerTurn string
}

func (chessGame *ChessGame) getPiece(piece string, x, y int) IChessPiece {
	pieceValue := strings.ToLower(piece)
	switch pieceValue {
	case "p":
		return Pawn{ChessPiece{x, y, PawnScore, piece}}
	case "r":
		return Rook{ChessPiece{x, y, RookScore, piece}}
	case "b":
		return Bishop{ChessPiece{x, y, BishopScore, piece}}
	case "n":
		return Knight{ChessPiece{x, y, KnightScore, piece}}
	case "q":
		return Queen{ChessPiece{x, y, QueenScore, piece}}
	case "k":
		return King{ChessPiece{x, y, KnightScore, piece}}
	default:
		return nil
	}
}

func (chessGame *ChessGame) getWhitePieces() []IChessPiece {
	var whitePieces []IChessPiece
	for i := 0; i < len(chessGame.board); i++ {
		for j := 0; j < len(chessGame.board[i]); j++ {
			if isWhitePiece(chessGame.board[i][j]) {
				piece := chessGame.getPiece(chessGame.board[i][j], i, j)
				whitePieces = append(whitePieces, piece)
			}
		}
	}
	return whitePieces
}

func (chessGame *ChessGame) getBoardValueForPieces(pieces []IChessPiece) int {
	var totalScore = 0
	for _, piece := range pieces {
		totalScore += piece.getValue()
	}
	return totalScore
}

func (chessGame *ChessGame) getBlackPieces() []IChessPiece {
	var blackPieces []IChessPiece
	for i := 0; i < len(chessGame.board); i++ {
		for j := 0; j < len(chessGame.board[i]); j++ {
			if isBlackPiece(chessGame.board[i][j]) {
				piece := chessGame.getPiece(chessGame.board[i][j], i, j)
				blackPieces = append(blackPieces, piece)
			}
		}
	}
	return blackPieces
}

func (chessGame *ChessGame) getMovableMoves() []Move {
	var availableMoves []Move
	if chessGame.playerTurn == BlackTurn {
		for i := 0; i < len(chessGame.board); i++ {
			for j := 0; j < len(chessGame.board[i]); j++ {
				locationString := chessGame.board[i][j]
				if isWhitePiece(locationString) || locationString == EmptySpace {
					piece := chessGame.getPiece(locationString, i, j)
					move := Move{i, j, 0, 0, 0, piece}
					availableMoves = append(availableMoves, move)
				}
			}
		}
	} else {
		for i := 0; i < len(chessGame.board); i++ {
			for j := 0; j < len(chessGame.board[i]); j++ {
				locationString := chessGame.board[i][j]
				if isBlackPiece(locationString) || locationString == EmptySpace {
					piece := chessGame.getPiece(locationString, i, j)
					move := Move{i, j, 0, 0, 0, piece}
					availableMoves = append(availableMoves, move)
				}
			}
		}
	}
	return availableMoves
}

func (chessGame *ChessGame) getScore() map[string]int {
	var scoreMap map[string]int
	return scoreMap
}

func (chessGame *ChessGame) getQueens(pieces []IChessPiece) []IChessPiece {
	var queenPieces []IChessPiece
	for _, piece := range pieces {
		if piece.getValue() == QueenScore {
			queenPieces = append(queenPieces, piece)
		}
	}
	return queenPieces
}

func (chessGame *ChessGame) getRooks(pieces []IChessPiece) []IChessPiece {
	var rookPieces []IChessPiece
	for _, piece := range pieces {
		if piece.getValue() == RookScore {
			rookPieces = append(rookPieces, piece)
		}
	}
	return rookPieces
}

func (chessGame *ChessGame) getBishops(pieces []IChessPiece) []IChessPiece {
	var bishopPieces []IChessPiece
	for _, piece := range pieces {
		if piece.getValue() == BishopScore {
			bishopPieces = append(bishopPieces, piece)
		}
	}
	return bishopPieces
}

func (chessGame *ChessGame) getKnights(pieces []IChessPiece) []IChessPiece {
	var knightPieces []IChessPiece
	for _, piece := range pieces {
		if piece.getValue() == KnightScore {
			knightPieces = append(knightPieces, piece)
		}
	}
	return knightPieces
}

func (chessGame *ChessGame) executeMoveMinMax() (string, int) {
	pieces := chessGame.getPiecesForTurn()
	// if len(pieces) < 9 {
	// 	MaxRecursiveLevel = MaxRecursiveLevel + 1
	// }
	println(MaxRecursiveLevel)
	movesMapping := getAllAvailableMovesForTurn(pieces, chessGame)

	enemyTurn := getNextPlayerTurn(chessGame.playerTurn)
	enemyGame := ChessGame{chessGame.board, enemyTurn}
	enemyPieces := enemyGame.getPiecesForTurn()

	friendlyValue := chessGame.getBoardValueForPieces(pieces)
	enemyValue := enemyGame.getBoardValueForPieces(enemyPieces)

	initialPiecesCount := len(pieces)
	valueDiff := friendlyValue - enemyValue

	score, move, pieceMove := minMax(movesMapping, chessGame.board, chessGame.playerTurn, MaxRecursiveLevel, chessGame.playerTurn, Move{}, Pawn{}, -1000000000, 1000000000, valueDiff, initialPiecesCount)
	moveTranslation := translateMove(pieceMove, move, chessGame.board)
	fmt.Println(moveTranslation)
	return moveTranslation, score
}

func (chessGame *ChessGame) getPiecesForTurn() []IChessPiece {
	if chessGame.playerTurn == WhiteTurn {
		return chessGame.getWhitePieces()
	}
	return chessGame.getBlackPieces()
}
