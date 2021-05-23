package chessgo

var (
	pieceValues = make(map[int]int)

	pawn = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	knight = []int{
		-3, 0, 0, 0, 0, 0, 0, -3,
	}
	bishop = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	rook = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	queen = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	king = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	pawnTable = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		50, 50, 50, 50, 50, 50, 50, 50,
		10, 10, 20, 30, 30, 20, 10, 10,
		5, 5, 10, 25, 25, 10, 5, 5,
		0, 0, 0, 20, 20, 0, 0, 0,
		5, -5, -10, 0, 0, -10, -5, 5,
		5, 10, 10, -20, -20, 10, 10, 5,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	pawnTableEndGame = []int{
		900, 900, 900, 900, 900, 900, 900, 900,
		500, 500, 500, 500, 500, 500, 500, 500,
		300, 300, 300, 300, 300, 300, 300, 300,
		90, 90, 90, 100, 100, 90, 90, 90,
		70, 70, 70, 85, 85, 70, 70, 70,
		20, 20, 20, 20, 20, 20, 20, 20,
		-10, -10, -10, -10, -10, -10, -10, -10,
		0, 0, 0, 0, 0, 0, 0, 0,
	}

	knightTable = []int{
		-50, -40, -30, -30, -30, -30, -40, -50,
		-40, -20, 0, 0, 0, 0, -20, -40,
		-30, 0, 10, 15, 15, 10, 0, -30,
		-30, 5, 15, 20, 20, 15, 5, -30,
		-30, 0, 15, 20, 20, 15, 0, -30,
		-30, 5, 10, 15, 15, 10, 5, -30,
		-40, -20, 0, 5, 5, 0, -20, -40,
		-50, -40, -30, -30, -30, -30, -40, -50,
	}

	bishopTable = []int{
		-20, -10, -10, -10, -10, -10, -10, -20,
		-10, 0, 0, 0, 0, 0, 0, -10,
		-10, 0, 5, 10, 10, 5, 0, -10,
		-10, 5, 5, 10, 10, 5, 5, -10,
		-10, 0, 10, 10, 10, 10, 0, -10,
		-10, 10, 10, 10, 10, 10, 10, -10,
		-10, 5, 0, 0, 0, 0, 5, -10,
		-20, -10, -10, -10, -10, -10, -10, -20,
	}

	rookTable = []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		5, 10, 10, 10, 10, 10, 10, 5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		-5, 0, 0, 0, 0, 0, 0, -5,
		0, 0, 0, 5, 5, 0, 0, 0,
	}

	queenTable = []int{
		-20, -10, -10, -5, -5, -10, -10, -20,
		-10, 0, 0, 0, 0, 0, 0, -10,
		-10, 0, 5, 5, 5, 5, 0, -10,
		-5, 0, 5, 5, 5, 5, 0, -5,
		0, 0, 5, 5, 5, 5, 0, -5,
		-10, 5, 5, 5, 5, 5, 0, -10,
		-10, 0, 5, 0, 0, 0, 0, -10,
		-20, -10, -10, -5, -5, -10, -10, -20,
	}

	kingEndGameTable = []int{
		-175, -175, -175, -175, -175, -175, -175, -175,
		-175, -50, -50, -50, -50, -50, -50, -175,
		-175, -50, 50, 50, 50, 50, -50, -175,
		-175, -50, 50, 150, 150, 50, -50, -175,
		-175, -50, 50, 100, 100, 50, -50, -175,
		-175, -50, 50, 50, 50, 50, -50, -175,
		-175, -50, -50, -50, -50, -50, -50, -175,
		-175, -175, -175, -175, -175, -175, -175, -175,
	}

	kingMiddleGameTable = []int{
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-20, -30, -30, -40, -40, -30, -30, -20,
		-10, -20, -20, -20, -20, -20, -20, -10,
		20, 20, 0, 0, 0, 0, 20, 20,
		20, 30, 10, 0, 0, 10, 30, 20,
	}

	pieceSquareTables = [][]int{
		pawnTable,
		knightTable,
		bishopTable,
		rookTable,
		queenTable,
		kingMiddleGameTable,
	}
)

func pieceEval(b *Board, col color) int {
	var ans int
	for i := Pawn; i <= King; i++ {
		piecebb := b.pieces[i] & b.colors[col]
		for j := piecebb; j.countSet() != 0; j.nextSet() {
			a := j.firstSet()
			if col == BLACK {
				v := pieceSquareTables[i][a]
				ans += v
			} else {
				var p []int = pieceSquareTables[i]
				reverse(p)
				v := p[a]
				ans += v
				reverse(p)
			}
		}

	}
	return ans
}

func pieceValue(b *Board, col color) int {
	var ans int
	mult := 1
	if col == BLACK {
		mult = -1
	}
	for i := Pawn; i <= King; i++ {
		bb := b.bitBoard(col, i)
		ans += pieceValues[i] * bb.countSet()
	}
	return ans * mult
}

func Evaluate(b *Board) int {
	if b.CheckMate() {
		// log.Println("in here")
		if b.turn == WHITE {
			return 10000
		}
		return -10000
	}
	if b.StaleMate() {
		return 5000
	}
	ans := 0
	// ans += int(pieceValue(b, WHITE) / 100)
	// ans += int(pieceValue(b, BLACK) / 100)
	ans += int(pieceEval(b, WHITE) / 1000)
	ans += int(pieceEval(b, BLACK) / 1000)
	return ans
}

func Search(depth int, board *Board) move {
	var bestMoveFound move
	col := board.TurnToMove()
	var maxScore int
	for _, move := range board.stdMoves(col) {
		board.move(move.str)
		score := negamax(board, depth, -999999, 999999, col)
		if score >= maxScore {
			bestMoveFound = move
			maxScore = score
		}
		board.unmove()
	}
	return bestMoveFound
}

func negamax(board *Board, depth, alpha, beta int, col color) int {
	if depth == 0 || len(board.stdMoves((col))) == 0 {
		return Evaluate(board)
	}
	for _, move := range board.stdMoves(board.TurnToMove()) {
		board.move(move.str)
		alpha = max(alpha, -negamax(board, depth, -beta, -alpha, -col.opp()))
		board.unmove()
		if alpha >= beta {
			break
		}
	}
	return alpha
}
