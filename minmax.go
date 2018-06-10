package main

import "fmt"

type PieceMove struct {
	move  Move
	piece IChessPiece
	score int
}

var kingKills = 0

func minMax(moveMapping map[IChessPiece][]Move, board [8][8]string, turn string, level int, initialTurn string, parentMove Move, parentPiece IChessPiece, alpha, beta, initialScore, initialPiecesCount int) (int, Move, IChessPiece) {

	if isKingKilled(board, turn) {
		if turn == initialTurn {
			return -10000 * (level + 1), parentMove, parentPiece
		}
		return 10000 * (level + 1), parentMove, parentPiece
	}

	if level == 0 {
		score := analyzeBoard(board, turn, initialTurn, initialScore, initialPiecesCount)
		return score, parentMove, parentPiece
	}

	prunedMap := pruneMinMax(moveMapping, board, turn)

	if turn == initialTurn {
		var maxMove Move
		var maxPiece IChessPiece
		//var bestMoveChoice Move
		var breakOuter bool
		for piece, moves := range prunedMap {
			for _, move := range moves {
				resultMax, _, _ := makeMoveAndGenerate(move, piece, turn, board, level-1, initialTurn, alpha, beta, initialScore, initialPiecesCount)
				if level == MaxRecursiveLevel {
					fmt.Println(resultMax, move, piece)
				}
				if alpha < beta {
					alpha, maxMove, maxPiece = max(alpha, resultMax, maxMove, move, maxPiece, piece)
				} else {
					breakOuter = true
					break
				}
			}
			if breakOuter {
				break
			}
		}
		return alpha, maxMove, maxPiece
	} else {
		var minMove Move
		var minPiece IChessPiece
		var breakmin bool
		for piece, moves := range prunedMap {
			for _, move := range moves {
				resultMin, _, _ := makeMoveAndGenerate(move, piece, turn, board, level-1, initialTurn, alpha, beta, initialScore, initialPiecesCount)
				if beta > alpha {
					beta, minMove, minPiece = min(beta, resultMin, minMove, move, minPiece, piece)
				} else {
					breakmin = true
					break
				}
			}
			if breakmin {
				break
			}
		}
		return beta, minMove, minPiece
	}
}

func max(currentmax, newValue int, currentMove, move Move, currentPiece, piece IChessPiece) (int, Move, IChessPiece) {
	if newValue > currentmax {
		return newValue, move, piece
	}

	if currentPiece == nil {
		currentPiece = piece
		currentMove = move
	}

	return currentmax, currentMove, currentPiece
}

func min(currentmin, newValue int, currentMove, move Move, currentPiece, piece IChessPiece) (int, Move, IChessPiece) {
	if newValue < currentmin {
		return newValue, move, piece
	}

	if currentPiece == nil {
		currentPiece = piece
		currentMove = move
	}

	return currentmin, currentMove, currentPiece
}

var count = 0

func analyzeBoard(board [8][8]string, turn string, initialTurn string, initialScore, initialPiecesCount int) int {
	count += 1
	//Give Greater Value to same score diff but less pieces on board vs more pieces on the board
	//Take into Account pawn progression on board, pieces defended
	//Whether King in check or not
	//If Any pieces defended vs not defended?
	//always analyze current turn here

	friendlyGame := ChessGame{board, initialTurn}
	enemyTurn := getNextPlayerTurn(initialTurn)
	enemyGame := ChessGame{board, enemyTurn}
	friendlyPieces := friendlyGame.getPiecesForTurn()
	enemyPieces := enemyGame.getPiecesForTurn()

	friendlyValue := friendlyGame.getBoardValueForPieces(friendlyPieces)
	enemyValue := enemyGame.getBoardValueForPieces(enemyPieces)
	pieceValues := friendlyValue - enemyValue
	if pieceValues >= initialScore && len(friendlyPieces) <= initialPiecesCount {
		pieceValues += 2
	}
	friendMoves := getAllAvailableMovesForTurn(friendlyPieces, &friendlyGame)
	enemyMoves := getAllAvailableMovesForTurn(enemyPieces, &enemyGame)

	pawnScoreFriend := getPawnScorePosition(friendlyPieces)
	pawnScoreEnemy := getPawnScorePosition(enemyPieces)
	pawnScoreDiff := pawnScoreEnemy - pawnScoreFriend
	knightScoreFriend, friendknightCount := getKnightPositionScore(friendlyPieces)
	knightScoreEnemy, enemyknightCount := getKnightPositionScore(enemyPieces)
	friendbishopCount := getBishopPositionScore(friendlyPieces)
	enemybishopCount := getBishopPositionScore(enemyPieces)
	knightScoreDiff := knightScoreEnemy - knightScoreFriend
	// friendKing := getKing(initialTurn, board)
	// enemyKing := getKing(enemyTurn, board)
	if friendknightCount == 2 && friendbishopCount == 0 {
		pieceValues -= 1
	}
	if friendbishopCount == 2 && friendknightCount == 0 {
		pieceValues += 1
	}
	if enemyknightCount == 2 && enemybishopCount == 0 {
		pieceValues += 1
	}
	if friendbishopCount == 2 && friendknightCount == 0 {
		pieceValues -= 1
	}
	//xvalue check
	// if friendKing.xLocation() != 0 || friendKing.xLocation() != 7 {
	// 	pieceValues -= 1
	// }
	// if enemyKing.xLocation() != 0 || enemyKing.xLocation() != 7 {
	// 	pieceValues += 1
	// }
	//if unable to castle take away a point
	//when ahead the less pieces the better

	scoreValue := pieceValues + availableMoveScore(friendMoves, enemyMoves) + pawnScoreDiff + knightScoreDiff
	return scoreValue
}

func availableMoveScore(friendlen, enemylen map[IChessPiece][]Move) int {
	var friendCount int
	var enemyCount int
	//  postMoveBoard := makeBoardMove(piece, move, board)
	// defending := isEnemyDefendingMove(move, turn, postMoveBoard)
	//  //Remove any move that will cause King to die next turn
	//  //should be handled in does my move lead to checkmate
	//  pruneKingChecked := pruneMoveKingChecked(piece, turn, board, move, defending)
	// if pruneKingChecked {
	// 	return true, false
	// }
	for piecesfr, _ := range friendlen {
		friendCount += len(friendlen[piecesfr])
	}
	for piecesen, _ := range enemylen {
		enemyCount += len(enemylen[piecesen])
	}
	if friendCount == 0 {
		return -10000
	}
	if enemyCount == 0 {
		return 10000
	}
	diff := friendCount - enemyCount
	return diff / 4
}

func getPawnScorePosition(pieces []IChessPiece) int {
	//check double pawns, pawns near edges weaker
	//Think about adding pawn progression scores
	var zeroCol int
	var oneCol int
	var twoCol int
	var threeCol int
	var fourCol int
	var fiveCol int
	var sixCol int
	var sevenCol int
	for _, piece := range pieces {
		if piece.getValue() == PawnScore {
			switch {
			case piece.yLocation() == 0:
				zeroCol++
			case piece.yLocation() == 1:
				oneCol++
			case piece.yLocation() == 2:
				twoCol++
			case piece.yLocation() == 3:
				threeCol++
			case piece.yLocation() == 4:
				fourCol++
			case piece.yLocation() == 5:
				fiveCol++
			case piece.yLocation() == 6:
				sixCol++
			case piece.yLocation() == 7:
				sevenCol++
			}
		}
	}

	var score int
	if zeroCol > 0 || sevenCol > 0 {
		score += 1
	}
	if oneCol > 1 || twoCol > 1 || threeCol > 1 || fourCol > 1 || fiveCol > 1 || sixCol > 1 || sevenCol > 1 {
		score += 1
	}

	return score

}

func getKnightPositionScore(pieces []IChessPiece) (int, int) {
	var score int
	var knightCount int
	for _, piece := range pieces {
		if piece.getValue() == KnightScore {
			knightCount += 1
			if piece.yLocation() == 0 || piece.yLocation() == 7 {
				score += 1
			}
		}
	}
	return score, knightCount
}

func getBishopPositionScore(pieces []IChessPiece) int {
	var bishopCount int
	for _, piece := range pieces {
		if piece.getValue() == KnightScore {
			bishopCount += 1
		}
	}
	return bishopCount
}

func makeMoveAndGenerate(move Move, piece IChessPiece, turn string, board [8][8]string, level int, initialTurn string, alpha, beta, initialScore, initialPiecesCount int) (int, Move, IChessPiece) {
	newBoard := makeBoardMove(piece, move, board)
	newTurn := getNextPlayerTurn(turn)
	newChessGame := ChessGame{newBoard, newTurn}
	pieces := newChessGame.getPiecesForTurn()
	newMapping := getAllAvailableMovesForTurn(pieces, &newChessGame)
	return minMax(newMapping, newBoard, newTurn, level, initialTurn, move, piece, alpha, beta, initialScore, initialPiecesCount)
}

func pruneMinMax(moveMapping map[IChessPiece][]Move, board [8][8]string, turn string) map[IChessPiece][]Move {
	prunedMap := make(map[IChessPiece][]Move)
	var pruneList []Move
	for piece, moves := range moveMapping {
		for _, move := range moves {
			prune, pruneAll := pruneMoveMinMax(piece, move, board, turn)
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

	return prunedMap
	// for _, badMove := range pruneMoves {
	// 	delete(moveMapping, badPiece)
	// }

}

func pruneMoveMinMax(piece IChessPiece, move Move, board [8][8]string, turn string) (bool, bool) {
	//change defending moves to simulate the move first
	postMoveBoard := makeBoardMove(piece, move, board)
	defending := isEnemyDefendingMove(move, turn, postMoveBoard)
	//Remove any move that will cause King to die next turn
	//should be handled in does my move lead to checkmate
	pruneKingChecked := pruneMoveKingChecked(piece, turn, board, move, defending)
	if pruneKingChecked {
		return true, false
	}

	if isCheckmateMove(piece, move, turn, board, defending, -1) {

		//prune all other moves in map, this is clearly the best move
		//idea is to limit score for this path, should result in much lower score propagated up to original move
		return false, true
	}
	//dont kill a piece that is defended and lower value than you

	//does my move lead to a checkmate?
	// newb := makeBoardMove(piece, move, board)
	// next := getNextPlayerTurn(turn)
	// oppChess := ChessGame{newb, next}
	// oppopieces := oppChess.getPiecesForTurn()
	// newMovesMapping := getAllAvailableMovesForTurn(oppopieces, &oppChess)
	// for piecekey, movesvalue := range newMovesMapping {
	// 	for _, moveValue := range movesvalue {
	// 		postMoveInner := makeBoardMove(piecekey, moveValue, newb)
	// 		if isKingKilled(postMoveInner, turn) {
	// 			return true, false
	// 		}
	// 		//is king on board check for black here
	// 		//keep condition below also
	// 		// defendingInner := isEnemyDefendingMove(moveValue, next, postMoveInner)
	// 		// if isCheckmateMove(piecekey, moveValue, next, newb, defendingInner, -1) {
	// 		// 	return true, false
	// 		// }
	// 	}
	// }

	//Moving piece will not lead to invalid or checkmate under this comment, and I do not have a checkmate available

	if piece.getValue() == PawnScore && !defending && (move.x == 0 || move.x == 7) {
		return false, true
	}

	if move.chessPiece != nil {
		if move.chessPiece.getValue() == QueenScore && !defending {
			return false, true
		}
		if move.chessPiece.getValue() < piece.getValue() && defending {
			return true, false
		}
		// if move.chessPiece.getValue() > piece.getValue() {
		// 	return false, true
		// }

		//Moving to defended location would be bad if my side is not defending, do this
	}
	//do I want to choose this over everything, even killing enemy queen?

	return false, false
}
