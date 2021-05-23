package chessgo

import "fmt"

/*
	Exported fucntions
*/

// NewBoardFromFEN to create a board from a fen string
func NewBoardFromFEN(fen string) *Board {
	return fenToBoard(fen)

}

// SquareNameToInt to convert square names to integers
func SquareNameToInt(sq string) int {
	return squareNameToInt[sq]
}

// NewBoard function creates an empty board
func NewBoard() *Board {
	b := Board{}
	b.clear()
	return &b
}

// NewStdBoard function creates board with std chess position in place
func NewStdBoard() *Board {
	// chase := "5r1k/1pR3p1/4Q3/4P1pp/3P2P1/5PKP/2Pq4/5R2 b - - 0 35"
	// chase := "8/3R4/4K3/3nN3/3Pkr2/8/8/8 b - - 9 56"
	// chase := "rn2kb1r/pQp1pppp/5n2/8/8/2N1q3/PPP3PP/1K1b1BNR w kq - 1 10" // bug
	// chase := "3rk2N/p5pp/1pP5/4p3/QKBn1q2/2P5/PP1Pb1PP/RNB4R b - - 1 17"
	// chase := "7r/p2QP2p/kp4pR/2p5/2n1PP2/2q5/2R5/2K5 b - - 1 38"
	// chase := "1k1r4/5p2/2Q1p3/P1P5/2Pbp3/P2q4/1B6/KR6 b - - 2 29"
	// chase := "5rk1/pb3p1p/1pn2Q2/5N2/8/2br4/PP3PPP/6K1 w - - 3 24" //bug
	// chase := "r1bqr1k1/pp3p2/8/6p1/1b2p3/2Q1PNB1/PP3PP1/R3K2R w KQ - 2 19"
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// chase :=
	// b := fenToBoard(chase)
	b := fenToBoard(StartPos)
	return b
}

// GetDisplay give the unicode characters of pieces for display purposes
func GetDisplay(sq int) int {
	return display[sq]
}

// SquareName gives the corresponding square name of a square
func SquareName(sq int) string {
	return squareName[sq]
}

// DrawBoard To draw a playing board
func DrawBoard(b *Board) {
	drawBoard(b)
}

// // DrawBitboard To draw a playing bitboard
// func DrawBitboard(b *Board) {
// 	drawBitBoard(b.colors[WHITE] | b.colors[BLACK])
// }

// func GenAttacks(b *Board, color color) bitBoard {
// 	return b.genAttacks(color)
// }

// func DrawAttacksBB(b *Board, color color) {
// 	drawBitBoard(b.pawnAttacks(C2, color))
// }

// func Dbb(a uint) {
// 	drawBitBoard(bitBoard(a))
// }

// func Move(b *Board, str string) {
// 	b.move(str)
// }

/*
	Exported methods
*/

// TurnToMove give the player to move
func (b *Board) TurnToMove() color {
	return b.turn
}

// ToFEN give the fen of the current board position
func (b *Board) ToFEN() string {
	return b.toFEN()
}

// GetPiece12 give the corresponding piece12 integer of a piece
func (b *Board) GetPiece12(sq int) int {
	return b.square[sq]
}

// GenMove gives the available moves for the color
func (b *Board) GenMove(col color) []string {
	var moves []string
	for _, mov := range b.stdMoves(col) {
		moves = append(moves, mov.str)
	}
	return moves
}

// UnMove gives the available moves for the color
func (b *Board) UnMove() error {
	return b.unmove()
}

// Move gives the available moves for the color
func (b *Board) Move(mov string) error {
	return b.move(mov)
}

// CheckMate checks if the current board posiition is a checkmate
func (b *Board) CheckMate() bool {
	return b.kingInCheck(b.turn) && len(b.stdMoves(b.turn)) == 0
}

// StaleMate checks if the current board posiition is a checkmate
func (b *Board) StaleMate() bool {
	return !b.kingInCheck(b.turn) && len(b.stdMoves(b.turn)) == 0
}

// PlayoutMoves to play out a slice of moves
func (b *Board) PlayoutMoves(str []string) error {
	if len(str) == 0 {
		return fmt.Errorf("empty move list")
	}
	for _, s := range str {
		if err := b.move(string(s)); err != nil {
			return err
		}
	}
	return nil
}

// //
// func (b *Board) MoveStr(mov string)

func (b *Board) KingInCheck(col color) bool {
	return b.kingInCheck(col)
}
