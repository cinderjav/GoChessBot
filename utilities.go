package main

import (
	"fmt"
	"math"
)

type Move struct {
	x, y, score, avgScore, moveCount int
	chessPiece                       IChessPiece
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

func getKing(turn string, board [8][8]string) IChessPiece {
	var newChess = getChessGame(board, turn)
	var turnPieces = newChess.getPiecesForTurn()
	for _, piece := range turnPieces {
		if piece.getValue() == KingScore {
			return piece.getCopy(piece.xLocation(), piece.yLocation(), turn)
		}
	}
	return nil
}

func pruneMoveKingChecked(piece IChessPiece, turn string, board [8][8]string, move Move, defending bool) bool {
	if piece.getValue() != KingScore {
		KingCopy := getKing(turn, board)
		if KingCopy != nil {
			isKingChecked := isEnemyDefendingMove(Move{KingCopy.xLocation(), KingCopy.yLocation(), 0, 0, 0, KingCopy}, turn, board)
			if isKingChecked {
				simBoard := makeBoardMove(piece, move, board)
				isKingCheckedAfterSim := isEnemyDefendingMove(Move{KingCopy.xLocation(), KingCopy.yLocation(), 0, 0, 0, KingCopy}, turn, simBoard)
				if isKingCheckedAfterSim {
					return true
				}
			}
		} else {
			return true
		}
	}
	if piece.getValue() == KingScore && defending {
		return true
	}
	return false
}

func isCheckmateMove(piece IChessPiece, move Move, turn string, board [8][8]string, defending bool, level int) bool {
	//get next turn to get opponent king
	newTurn := getNextPlayerTurn(turn)
	//get opponent king
	opponentKingCopy := getKing(newTurn, board)
	//simulate a move on the original turn to move piece to spot
	simBoard := makeBoardMove(piece, move, board)
	//checks to see if any of my guys can move to opponent king after I move, therefore making opponent king checked
	isOpponentKingCheckedWithMove := isEnemyDefendingMove(Move{opponentKingCopy.xLocation(), opponentKingCopy.yLocation(), 0, 0, 0, opponentKingCopy}, newTurn, simBoard)
	//now check to see if I am targetable by the enemy once I get the king checked
	//problem is king is defending move, but his move is also defended
	// if level == 0 {
	// 	fmt.Println(piece, move, !defending && isOpponentKingCheckedWithMove, defending, isOpponentKingCheckedWithMove, opponentKingCopy)
	// }
	if !defending && isOpponentKingCheckedWithMove {
		//see if another piece from an enemy move, may cause king to be unchecked
		//make it the opponents turn now, after sim moves my piece
		var newChess = ChessGame{simBoard, newTurn}
		//get all pieces for opponent
		var opponentPieces = newChess.getPiecesForTurn()
		//get all the opponents potential moves in respond to my move
		movesMappingOpponents := getAllAvailableMovesForTurn(opponentPieces, &newChess)
		//for each piece, check his moves, I need to see if they can remove the king being checked
		//King should not be allowed to eat a defended move to clear check
		for pieceEnemy, moves := range movesMappingOpponents {
			//var countSafe = 0
			for _, moveEnemy := range moves {
				//need to sim enemy making available move to determine if it would remove check
				//what if straight up block sacrifice?
				simBoardNew := makeBoardMove(pieceEnemy, moveEnemy, simBoard)
				//if king I want to move opponent king copy
				if pieceEnemy.getValue() == KingScore {
					opponentKingCopy = getKing(newTurn, simBoardNew)
				}
				//create a new chess object with the newboard after the move
				//change control of turn back to original
				var chessTwo = ChessGame{simBoardNew, turn}
				//see if king is still checked by iterating through and see if i can still attack the king
				//if i find a piece that can attack the king, then that piece was not sufficient in saving checkmate
				//if I cant find a piece that can attack the king, then no checkmate is present
				var originalPieces = chessTwo.getPiecesForTurn()
				var countMoveKing = 0
				for _, pieceAgain := range originalPieces {
					//check to see if any of pieces can still attack king, if any are found, then we have a checkmate for this move
					//we need to go through all moves, and if we find a move that does not allow king to be checked then no checkmate
					//if I cant move there, that is a good thing for the enemy no checkmate
					if pieceAgain.canMove(Move{opponentKingCopy.xLocation(), opponentKingCopy.yLocation(), 0, 0, 0, opponentKingCopy}, simBoardNew) {
						countMoveKing++
					}
					//isKingChecked()
					//if any remove check break
				}
				//if counrMoveKing > 1, not a life saving move for enemy
				//keep checking
				//if count = 0, doesn't mean it is over, it means he founds a safe spot

				if countMoveKing == 0 && len(originalPieces) > 0 {
					return false
				}
			}
			//check here if piece could not find empty spot
		}
		return true
	}
	return false
}

func isKingKilled(board [8][8]string, turn string) bool {
	var king string
	if turn == WhiteTurn {
		king = "K"
	} else {
		king = "k"
	}

	for _, row := range board {
		for _, col := range row {
			if col == king {
				return false
			}
		}
	}
	return true

}

func isEnemyDefendingMove(move Move, turn string, board [8][8]string) bool {
	//problem is king is defending move, but his move is also defended
	var nextTurn = getNextPlayerTurn(turn)
	var newChess = ChessGame{board, nextTurn}

	var opponentPieces = newChess.getPiecesForTurn()
	for _, piece := range opponentPieces {
		//want to check if the piece is king, and they can move there,
		if (piece.canMove(move, board)) && (piece.getValue() == KingScore) {
			//get
			var boardAfterMove = makeBoardMove(piece, move, board)
			var newChessKing = ChessGame{boardAfterMove, nextTurn}
			var myPieces = newChessKing.getPiecesForTurn()
			for _, pieceKingMove := range myPieces {
				//if king will die if he makes move, it is not defensible
				if pieceKingMove.canMove(move, boardAfterMove) {
					return false
				}
			}
		} else if piece.canMove(move, board) {
			return true
		}
	}
	return false
}

func isDefendingMove(turn string, board [8][8]string, decisionPiece IChessPiece, move Move, currentMove bool) (bool, []IChessPiece, []IChessPiece, int) {
	//get all black moves
	//simulate me moving cells
	//see if I can attack one of blacks available moves
	var nextTurn = getNextPlayerTurn(turn)
	var newChess = ChessGame{board, nextTurn}
	var opponentPieces = newChess.getPiecesForTurn()
	movesMappingOpponents := getAllAvailableMovesForTurn(opponentPieces, &newChess)
	result := false
	var copyPiece IChessPiece
	if !currentMove {
		copyPiece = decisionPiece.getCopy(move.x, move.y, turn)
	} else {
		copyPiece = decisionPiece
	}
	var piecesDefendedAgainst []IChessPiece
	var piecesProtecting []IChessPiece
	value := -1
	for piece, moves := range movesMappingOpponents {
		for _, mov := range moves {
			if copyPiece.canMove(mov, board) && mov.chessPiece != nil {
				result = true
				piecesDefendedAgainst = append(piecesDefendedAgainst, piece)
				piecesProtecting = append(piecesProtecting, mov.chessPiece)
			} else if copyPiece.canMove(mov, board) {
				result = true
			}
		}
	}
	if result && len(piecesProtecting) == 1 {
		value = 1
	}
	if result && len(piecesProtecting) == 0 {
		value = 0
	}
	if result && len(piecesProtecting) >= 2 {
		value = 2
	}
	return result, piecesDefendedAgainst, piecesProtecting, value

	//get all
}

func makeBoardMove(piece IChessPiece, move Move, board [8][8]string) [8][8]string {
	pieceString := board[piece.xLocation()][piece.yLocation()]
	var queenPromotion string
	if pieceString == BlackPawn && move.x == 7 {
		queenPromotion = BlackQueen
	} else if pieceString == WhitePawn && move.x == 0 {
		queenPromotion = WhiteQueen
	}
	board[piece.xLocation()][piece.yLocation()] = ""
	if queenPromotion == WhiteQueen || queenPromotion == BlackQueen {
		board[move.x][move.y] = queenPromotion
	} else {
		board[move.x][move.y] = pieceString
	}

	return board
}

func getNextPlayerTurn(currentTurn string) string {
	if currentTurn == WhiteTurn {
		return BlackTurn
	}
	return WhiteTurn
}

func getHighestMoveScoreFromMap(moveMapping map[IChessPiece][]Move) (int, IChessPiece, Move, int, int) {
	score := -1000000000
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
	if sumCount == 0 {
		return 0, nil, Move{}, 0, sumCount
	}
	average := sumScore / sumCount
	return sumScore, topPiece, topMove, average, sumCount
}

func getPieceString(x, y int, board [8][8]string) string {
	return board[x][y]
}

func translateMove(piece IChessPiece, move Move, board [8][8]string) string {
	var pieceNotation string
	var moveNotation string
	//send back algebraic and basic just in case
	// pieceString := getPieceString(piece.xLocation(), piece.yLocation(), board)
	// //moveString := getPieceString(move.x, move.y, board)
	// if pieceString == "P" {
	// 	pieceString = ""
	// }
	//pieceNotation += strings.ToUpper(pieceString)
	//not sure here
	fmt.Println(move, piece)
	if piece == nil {
		return "stalemate"
	}
	//fmt.Println(move, piece)
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
	var queenPromotion string
	pieceString := board[piece.xLocation()][piece.yLocation()]
	if pieceString == BlackPawn && move.x == 7 {
		queenPromotion = BlackQueen
	} else if pieceString == WhitePawn && move.x == 0 {
		queenPromotion = WhiteQueen
	}

	if queenPromotion == WhiteQueen || queenPromotion == BlackQueen {
		return pieceNotation + sep + moveNotation + "/" + queenPromotion
	}
	return pieceNotation + sep + moveNotation

}
