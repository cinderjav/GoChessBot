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

func getKing(turn string, board [8][8]string) IChessPiece {
	var newChess = ChessGame{board, turn}
	var turnPieces = newChess.getPiecesForTurn()
	for _, piece := range turnPieces {
		if piece.getValue() == KingScore {
			return piece.getCopy(piece.xLocation(), piece.yLocation(), turn)
		}
	}
	return nil
}

func pruneMove(piece IChessPiece, move Move, board [8][8]string, turn string, level int, initialTurn bool, lastEaten IChessPiece) bool {
	// if piece.getValue() == KingScore || piece.getValue() == QueenScore {
	// 	return true
	// }
	// return false
	defending := isEnemyDefendingMove(move, turn, board)
	if piece.getValue() != KingScore {
		KingCopy := getKing(turn, board)
		if KingCopy != nil {
			isKingChecked := isEnemyDefendingMove(Move{KingCopy.xLocation(), KingCopy.yLocation(), 0, 0, 0, KingCopy}, turn, board)
			if isKingChecked {
				simBoard := makeBoardMove(piece, move, board)
				isKingCheckedAfterSim := isEnemyDefendingMove(Move{KingCopy.xLocation(), KingCopy.yLocation(), 0, 0, 0, KingCopy}, turn, simBoard)
				if level == 0 {
					fmt.Println(piece)
				}
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

	newTurn := getNextPlayerTurn(turn)
	opponentKingCopy := getKing(newTurn, board)
	simBoard := makeBoardMove(piece, move, board)
	//checks to see if any of my guys can move to king after I move
	isOpponentKingCheckedWithMove := isEnemyDefendingMove(Move{opponentKingCopy.xLocation(), opponentKingCopy.yLocation(), 0, 0, 0, opponentKingCopy}, newTurn, simBoard)
	//now check to see if I am targetable by the enemy
	if !defending && isOpponentKingCheckedWithMove {
		//see if another piece from an enemy move, may cause king to be unchecked
		var newChess = ChessGame{simBoard, newTurn}
		var opponentPieces = newChess.getPiecesForTurn()
		movesMappingOpponents := getAllAvailableMovesForTurn(opponentPieces, &newChess)
		for pieceEnemy, moves := range movesMappingOpponents {
			for _, moveEnemy := range moves {
				simBoardNew := makeBoardMove(pieceEnemy, moveEnemy, simBoard)
				var chessTwo = ChessGame{simBoardNew, newTurn}
				var opponentPiecesTwo = chessTwo.getPiecesForTurn()
				for pieceAgain := range opponentPiecesTwo {
					//isKingChecked()
					//if any remove check break
				}
			}
		}
	}

	//prune out other moves if checkmate
	//think about how to bubble up negative scores from before
	//you have a move, that puts opponent king in check, he can't eat you or safely move
	//No other move for opponent will remove check from king
	if move.chessPiece != nil {

		if move.chessPiece.getValue() == KingScore && initialTurn {
			return true
		}

		if move.chessPiece.getValue() == KingScore && !initialTurn {
			return true
		}

		//test
		//and not defended, this condition is ok if piece if defended and its value is less than attacker
		movePieceValue := move.chessPiece.getValue()
		diffValue := movePieceValue - piece.getValue()
		if (lastEaten == nil && !initialTurn && level == 1) && (!defending || diffValue > 0) {
			return true
		}

		if lastEaten != nil {
			lastEatenValue := lastEaten.getValue()
			combinedValue := lastEatenValue + piece.getValue()
			if move.chessPiece.getValue() > lastEaten.getValue() && !initialTurn && level == 1 && !defending {
				return true
			}
			if defending && combinedValue < move.chessPiece.getValue() && !initialTurn && level == 1 {
				return true
			}
		}
	}

	return false
}

func pruneMap(moveMapping map[IChessPiece][]Move, board [8][8]string, turn string, level int, originalTurn string, prevEaten IChessPiece) map[IChessPiece][]Move {
	prunedMap := make(map[IChessPiece][]Move)
	var pruneList []Move
	for piece, moves := range moveMapping {
		for _, move := range moves {
			prune := pruneMove(piece, move, board, turn, level, turn == originalTurn, prevEaten)
			if prune {
				pruneList = append(pruneList, move)
			} else {
				prunedMap[piece] = append(prunedMap[piece], move)
			}
		}
	}

	//moveMapping = prunedMap
	if level == 0 {
		fmt.Println(prunedMap)
		fmt.Println(pruneList)
	}
	return prunedMap
	// for _, badMove := range pruneMoves {
	// 	delete(moveMapping, badPiece)
	// }

}

func analyzeMoves(moveMapping map[IChessPiece][]Move, chessGame *ChessGame, level int, score int, originalTurn string, prevEaten IChessPiece, avgScore int, lenMoves int) map[IChessPiece][]Move {
	prunedMap := pruneMap(moveMapping, chessGame.board, chessGame.playerTurn, level, originalTurn, prevEaten)
	for piece, moves := range prunedMap {
		for index, move := range moves {
			score, averageScore, lenScore := analyzeMove(piece, move, chessGame.board, chessGame.playerTurn, level, score, originalTurn, prevEaten, avgScore, lenMoves)
			moves := prunedMap[piece]
			moves[index].score = score
			moves[index].avgScore = averageScore
			moves[index].moveCount = lenScore
			// go func(piece IChessPiece, move Move, board [8][8]string, turn string, level int, score int, originalTurn string, moveMapping map[IChessPiece][]Move) {
			// 	scores <- analyzeMove(piece, move, board, turn, level, score, originalTurn)
			// 	moves := moveMapping[piece]
			// 	newScore := <-scores
			// 	moves[index].score = newScore
			// }(piece, move, chessGame.board, chessGame.playerTurn, level, score, originalTurn, moveMapping)
		}
	}
	return prunedMap
}

func analyzeMove(piece IChessPiece, move Move, board [8][8]string, turn string, level int, score int, originalTurn string, lastEaten IChessPiece, avgScore int, lenMoves int) (int, int, int) {
	//possible issue not passing along score
	//implement shouldprune and getindividual move score
	//reason about the implementation
	//score needs to take into account my color
	// prune, value := shouldPrune(piece, move, board, turn == originalTurn, level, score, lastEaten, turn)
	// if prune {
	// 	return score + (value * ((MaxRecursiveLevel + 1) - level))
	// }
	scoreMove := getIndividualMoveScore(piece, move, board, turn)
	if turn == originalTurn {
		score += (scoreMove * ((MaxRecursiveLevel + 1) - level))
	} else {
		score -= (scoreMove * ((MaxRecursiveLevel + 1) - level))
	}
	if level == MaxRecursiveLevel {
		return score, avgScore, lenMoves
	}

	newBoard := makeBoardMove(piece, move, board)
	newTurn := getNextPlayerTurn(turn)
	newChessGame := ChessGame{newBoard, newTurn}
	pieces := newChessGame.getPiecesForTurn()
	newMovesMapping := getAllAvailableMovesForTurn(pieces, &newChessGame)
	if move.chessPiece != nil {
		lastEaten = move.chessPiece
	}
	prunedMap := analyzeMoves(newMovesMapping, &newChessGame, level+1, score, originalTurn, lastEaten, avgScore, lenMoves)
	highScore, _, _, averageScore, lenMoves := getHighestMoveScoreFromMap(prunedMap)
	return highScore, averageScore, lenMoves
}

// func shouldPrune(piece IChessPiece, move Move, board [8][8]string, initialTurn bool, level int, score int, lastEaten IChessPiece, turn string) (bool, int) {
// 	//this function will can return high value end game
// 	//need to make sure smallers turns are favored, pass in recursive level
// 	if move.chessPiece != nil {
// 		defending := isEnemyDefendingMove(move, turn, board)

// 		if move.chessPiece.getValue() == KingScore && initialTurn {
// 			return true, KingScore
// 		}

// 		if move.chessPiece.getValue() == KingScore && !initialTurn && level != 1 {
// 			return true, -KingScore
// 		}

// 		if move.chessPiece.getValue() == KingScore && !initialTurn && level == 1 {
// 			return true, -10000000
// 		}

// 		if level == 0 && piece.getValue() == KingScore && defending {
// 			return true, -100000000
// 		}
// 		//test
// 		//and not defended, this condition is ok if piece if defended and its value is less than attacker
// 		movePieceValue := move.chessPiece.getValue()
// 		diffValue := movePieceValue - piece.getValue()
// 		if (lastEaten == nil && !initialTurn && level == 1) && (!defending || diffValue > 0) {
// 			return true, -6000000
// 		}

// 		if lastEaten != nil {
// 			lastEatenValue := lastEaten.getValue()
// 			combinedValue := lastEatenValue + piece.getValue()
// 			if move.chessPiece.getValue() > lastEaten.getValue() && !initialTurn && level == 1 && !defending {
// 				return true, -6000000
// 			}
// 			if defending && combinedValue < move.chessPiece.getValue() && !initialTurn && level == 1 {
// 				return true, -6000000
// 			}
// 		}
// 	}

// 	return false, 0
// }

func getIndividualMoveScore(piece IChessPiece, move Move, board [8][8]string, turn string) int {
	//Moving to a location which defends, should be 0 and empty space should be -1
	//value moving to a cell that is defended should also be rewarded
	//blocking pieces from getting killed with less valuable pieces that are defended
	if move.chessPiece != nil {
		return move.chessPiece.getValue()
	}

	defending, _, _, value := isDefendingMove(turn, board, piece, move, false)
	if defending {
		return value
	}

	return -1
}

func isEnemyDefendingMove(move Move, turn string, board [8][8]string) bool {
	//prob want to make the move here then check
	var nextTurn = getNextPlayerTurn(turn)
	var newChess = ChessGame{board, nextTurn}
	var opponentPieces = newChess.getPiecesForTurn()
	for _, piece := range opponentPieces {
		if piece.canMove(move, board) {
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
		for _, move := range moves {
			if copyPiece.canMove(move, board) && move.chessPiece != nil {
				result = true
				piecesDefendedAgainst = append(piecesDefendedAgainst, piece)
				piecesProtecting = append(piecesProtecting, move.chessPiece)
			} else if copyPiece.canMove(move, board) {
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
