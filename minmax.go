package main

import "fmt"

type PieceMove struct {
	move  Move
	piece IChessPiece
	score int
}

var kingKills = 0

func minMax(moveMapping map[IChessPiece][]Move, board [8][8]string, turn string, level int, initialTurn string, parentMove Move, parentPiece IChessPiece, alpha, beta int) (int, Move, IChessPiece) {
	if isKingKilled(board, turn) {
		if turn == initialTurn {
			return -10000 * (level + 1), parentMove, parentPiece
		}
		return 10000 * (level + 1), parentMove, parentPiece
	}

	if level == 0 {
		score := analyzeBoard(board, turn, initialTurn)
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
				resultMax, _, _ := makeMoveAndGenerate(move, piece, turn, board, level-1, initialTurn, alpha, beta)
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
				resultMin, _, _ := makeMoveAndGenerate(move, piece, turn, board, level-1, initialTurn, alpha, beta)
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

func analyzeBoard(board [8][8]string, turn string, initialTurn string) int {
	count += 1
	//Give Greater Value to same score diff but less pieces on board vs more pieces on the board
	//Take into Account pawn progression on board, pieces defended
	//Whether King in check or not
	//If Any pieces defended vs not defended?
	//stacked pawns are bad
	//more available moves in comparison to other side is better
	//look to add concurrency to increase depth
	//always analyze current turn here

	friendlyGame := ChessGame{board, initialTurn}
	enemyTurn := getNextPlayerTurn(initialTurn)
	enemyGame := ChessGame{board, enemyTurn}
	friendlyPieces := friendlyGame.getPiecesForTurn()
	enemyPieces := enemyGame.getPiecesForTurn()

	friendlyValue := friendlyGame.getBoardValueForPieces(friendlyPieces)
	enemyValue := enemyGame.getBoardValueForPieces(enemyPieces)

	friendMoves := getAllAvailableMovesForTurn(friendlyPieces, &friendlyGame)
	enemyMoves := getAllAvailableMovesForTurn(enemyPieces, &enemyGame)

	return friendlyValue - enemyValue + availableMoveScore(friendMoves, enemyMoves)
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

func makeMoveAndGenerate(move Move, piece IChessPiece, turn string, board [8][8]string, level int, initialTurn string, alpha, beta int) (int, Move, IChessPiece) {
	newBoard := makeBoardMove(piece, move, board)
	newTurn := getNextPlayerTurn(turn)
	newChessGame := ChessGame{newBoard, newTurn}
	pieces := newChessGame.getPiecesForTurn()
	newMapping := getAllAvailableMovesForTurn(pieces, &newChessGame)
	return minMax(newMapping, newBoard, newTurn, level, initialTurn, move, piece, alpha, beta)
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
