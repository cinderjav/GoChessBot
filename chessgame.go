package main

import (
	"strings"
)

type ChessGame struct {
	board      [8][8]string
	playerTurn string
}

func (chessGame *ChessGame) movePiece(move Move) {

}

func (chessGame *ChessGame) getPiece(piece string, x, y int) IChessPiece {
	pieceValue := strings.ToLower(piece)
	switch pieceValue {
	case "p":
		return Pawn{ChessPiece{x, y, 1, piece}}
	case "r":
		return Rook{ChessPiece{x, y, 3, piece}}
	case "b":
		return Bishop{ChessPiece{x, y, 2, piece}}
	case "n":
		return Knight{ChessPiece{x, y, 2, piece}}
	case "q":
		return Queen{ChessPiece{x, y, 5, piece}}
	case "k":
		return King{ChessPiece{x, y, 100, piece}}
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
					move := Move{i, j, piece}
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
					move := Move{i, j, piece}
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

func (chessGame *ChessGame) executeMove() Move {
	pieces := chessGame.getPiecesForTurn()
	moves := getAllAvailableMovesForTurn(pieces, chessGame)
	//given current available moves, calls a recursive function to get the best move from these available moves
	//specify a depth level I want to search for start small
	//fmt.Println(moves)
	return moves[0]
}

func (chessGame *ChessGame) getPiecesForTurn() []IChessPiece {
	if chessGame.playerTurn == WhiteTurn {
		return chessGame.getWhitePieces()
	}
	return chessGame.getBlackPieces()
}
