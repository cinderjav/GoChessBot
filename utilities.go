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

func pruneMove(piece IChessPiece, move Move, board [8][8]string, turn string, level int, initialTurn bool, lastEaten IChessPiece) (bool, bool) {
	//change defending moves to simulate the move first
	postMoveBoard := makeBoardMove(piece, move, board)
	defending := isEnemyDefendingMove(move, turn, postMoveBoard)
	pruneKingChecked := pruneMoveKingChecked(piece, turn, board, move, defending)
	if pruneKingChecked {
		return true, false
	}

	if isCheckmateMove(piece, move, turn, board, defending, level) {
		if level == 0 {
			fmt.Println("Check mate move found!")
		}

		//prune all other moves in map, this is clearly the best move
		//idea is to limit score for this path, should result in much lower score propagated up to original move
		return false, true
	}

	//prune any move that results in checkmate after I move

	if move.chessPiece != nil {
		//think about if these conditions should be a score hit or a prune
		//ideally would need to find a way to propagate up, score might be only way

		//why do it need this? This is actually good thing, move to score
		if move.chessPiece.getValue() == KingScore && initialTurn {
			return false, true
		}

		if move.chessPiece.getValue() == KingScore && !initialTurn {
			return false, true
		}

		//test
		//and not defended, this condition is ok if piece if defended and its value is less than attacker
		movePieceValue := move.chessPiece.getValue()
		diffValue := movePieceValue - piece.getValue()
		if (lastEaten == nil && !initialTurn && level == 1) && (!defending || diffValue > 0) {
			return false, false
		}

		if lastEaten != nil {
			lastEatenValue := lastEaten.getValue()
			combinedValue := lastEatenValue + piece.getValue()
			if move.chessPiece.getValue() > lastEaten.getValue() && !initialTurn && level == 1 && !defending {
				return false, true
			}
			if defending && combinedValue < move.chessPiece.getValue() && !initialTurn && level == 1 {
				return false, true
			}
		}
	}

	return false, false
}

func pruneMap(moveMapping map[IChessPiece][]Move, board [8][8]string, turn string, level int, originalTurn string, prevEaten IChessPiece) map[IChessPiece][]Move {
	prunedMap := make(map[IChessPiece][]Move)
	var pruneList []Move
	for piece, moves := range moveMapping {
		for _, move := range moves {
			prune, pruneAll := pruneMove(piece, move, board, turn, level, turn == originalTurn, prevEaten)
			if pruneAll {
				prunedMapAll := make(map[IChessPiece][]Move)
				prunedMapAll[piece] = append(prunedMapAll[piece], move)
				return prunedMapAll
			}
			if prune {
				pruneList = append(pruneList, move)
			} else {
				prunedMap[piece] = append(prunedMap[piece], move)
			}
		}
	}

	//moveMapping = prunedMap
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
	prune, value := shouldPrune(piece, move, board, turn == originalTurn, level, score, lastEaten, turn)
	if prune {
		return score + (value * ((MaxRecursiveLevel + 1) - level)), 0, 0
	}
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

func shouldPrune(piece IChessPiece, move Move, board [8][8]string, initialTurn bool, level int, score int, lastEaten IChessPiece, turn string) (bool, int) {
	//this function will can return high value end game
	//need to make sure smallers turns are favored, pass in recursive level
	//A move which results in an opponent king having no moves, and that move is not defended is a good move?
	if move.chessPiece != nil {
		defending := isEnemyDefendingMove(move, turn, board)

		if move.chessPiece.getValue() == KingScore && initialTurn {
			return true, KingScore
		}

		if move.chessPiece.getValue() == KingScore && !initialTurn && level != 1 {
			return true, -KingScore
		}

		if move.chessPiece.getValue() == KingScore && !initialTurn && level == 1 {
			return true, -KingScore * 1000
		}

		//test
		//and not defended, this condition is ok if piece if defended and its value is less than attacker
		movePieceValue := move.chessPiece.getValue()
		diffValue := movePieceValue - piece.getValue()
		if (lastEaten == nil && !initialTurn && level == 1) && (!defending || diffValue > 0) {
			return true, -KingScore * 100
		}

		if lastEaten != nil {
			lastEatenValue := lastEaten.getValue()
			combinedValue := lastEatenValue + piece.getValue()
			if move.chessPiece.getValue() > lastEaten.getValue() && !initialTurn && level == 1 && !defending {
				return true, -KingScore * 100
			}
			if defending && combinedValue < move.chessPiece.getValue() && !initialTurn && level == 1 {
				return true, -KingScore * 100
			}
		}
	}

	return false, 0
}

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
