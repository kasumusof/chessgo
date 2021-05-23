package chessgo

// constants for ranks and files
const (
	Rank1 = bitBoard(0xff)
	Rank2 = bitBoard(0xff00)
	Rank3 = bitBoard(0xff0000)
	Rank4 = bitBoard(0xff000000)
	Rank5 = bitBoard(0xff00000000)
	Rank6 = bitBoard(0xff0000000000)
	Rank7 = bitBoard(0xff000000000000)
	Rank8 = bitBoard(0xff00000000000000)

	FileA = bitBoard(0x0101010101010101)
	FileB = bitBoard(0x0202020202020202)
	FileC = bitBoard(0x0404040404040404)
	FileD = bitBoard(0x0808080808080808)
	FileE = bitBoard(0x1010101010101010)
	FileF = bitBoard(0x2020202020202020)
	FileG = bitBoard(0x4040404040404040)
	FileH = bitBoard(0x8080808080808080)
)

// constants for square names
const (
	A1 = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1

	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2

	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3

	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4

	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5

	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6

	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7

	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)
// piece12 = (piece6+1) * +-1
// constants for pieces(6)
const (
	Pawn = iota
	Knight
	Bishop
	Rook
	Queen
	King
)
// piece6 = abs(piece12)-1
// constants for pieces(12)
const (
	Empty = iota
	WPawn
	WKnight
	WBishop
	WRook
	WQueen
	WKing

	BPawn   = -WPawn
	BKnight = -WKnight
	BBishop = -WBishop
	BRook   = -WRook
	BQueen  = -WQueen
	BKing   = -WKing
)

// constants for colors
const (
	WHITE = color(0)
	BLACK = color(1)
)

/*
constants for directions
*/
// directions for peices excluding knights
const (
	N  = 8
	S  = -8
	E  = 1
	W  = -1
	NE = N + E
	NW = N + W
	SE = S + E
	SW = S + W
)

// directions for knights
const (
	NEN = N + E + N
	NEE = N + E + E
	NWN = N + W + N
	NWW = N + W + W
	SES = S + E + S
	SEE = S + E + E
	SWS = S + W + S
	SWW = S + W + W
)

// StartPos starting Position of standard chess fen
const StartPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const funnyPos = "r2q2k1/1bp1nppp/p7/1p2r3/8/P3n1Q1/1P2NPPP/3RKB1R w K - 0 1"
const funnyPos2 = "3rk3/2P5/8/8/8/8/1PPPPPP1/RNBQKBNR w KQ - 0 1"
