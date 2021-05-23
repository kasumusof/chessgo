package chessgo

import (
	"fmt"
)

// Board that holds the chess position
type Board struct {
	square    [64]int
	colors    [2]bitBoard
	pieces    [6]bitBoard
	turn      color
	castlings castlings
	rule50    int
	moveNo    int
	enpassant int
	history   moves
	state     states
}

func (b *Board) clear() {
	for i := A1; i <= H8; i++ {
		b.square[i] = Empty
	}
	for i := 0; i < 2; i++ {
		b.colors[i].clear()
	}
	for i := 0; i < 6; i++ {
		b.pieces[i].clear()
	}
	b.turn = WHITE
	b.castlings = 0x0
	b.rule50 = 0
	b.moveNo = 1
	b.enpassant = 0
}

func (b *Board) set(piece int, sq int) {
	col := WHITE
	p := abs(piece) - 1
	if piece < 0 {
		col = BLACK
	}
	b.unset(sq)
	b.square[sq] = piece
	b.colors[col].set(sq)
	b.pieces[p].set(sq)
}

func (b *Board) unset(sq int) {
	p12 := b.square[sq]
	if p12 == Empty {
		return
	}
	col := WHITE
	p6 := abs(p12) - 1
	if p12 < 0 {
		col = BLACK
	}
	b.square[sq] = Empty
	b.colors[col].unset(sq)
	b.pieces[p6].unset(sq)
}

func (b *Board) allPieces() bitBoard {
	return b.colors[WHITE] | b.colors[BLACK]
}

func (b *Board) bitBoard(col color, piece int) bitBoard {
	return b.colors[col] & b.pieces[piece]
}

// gen attacks

func (b *Board) strtAttacks(sq int) bitBoard {
	mv := bitBoard(0)
	var dMove bitBoard
	// to go up
	for i := sq + N; sq/8 != 7 && i <= H8; i += N {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}

	// to go down
	for i := sq + S; sq/8 != 0 && i >= A1; i += S {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}

	// to go right
	for i := sq + E; sq%8 != 7 && i <= H8 && i%8 != 0; i += E {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}

	// to go left
	for i := sq + W; sq%8 != 0 && i >= A1 && i%8 != 7; i += W {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}

	return mv
}

func (b *Board) diagAttacks(sq int) bitBoard {
	mv := bitBoard(0)
	var dMove bitBoard
	// to go north-east conditions are, in order; not from the last rank, not from the last file, not back to the first file(for west<->east movements) ,cant go beyond the last square
	for i := sq + NE; sq/8 != 7 && sq%8 != 7 && i%8 != 0 && i <= H8; i += NE {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}
	// to go south-west conditions are, in order; not from the first rank, not from the first file, not back to the last file(for west<->east movements) ,cant go beyond the first square
	for i := sq + SW; sq/8 != 0 && sq%8 != 0 && i%8 != 7 && i >= A1; i += SW {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}
	// to go north-west conditions are, in order; not from the last rank, not from the first file, not back to the last file(for west<->east movements) ,cant go beyond the last square
	for i := sq + NW; sq/8 != 7 && sq%8 != 0 && i%8 != 7 && i <= H8; i += NW {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}
	// to go south-east conditions are, in order; not from the first rank, not from the first file, not back to the last file(for west<->east movements) ,cant go beyond the first square
	for i := sq + SE; sq/8 != 0 && sq%8 != 7 && i%8 != 0 && i >= A1; i += SE {
		dMove = bitBoard(1 << i)
		if b.allPieces()&dMove != 0 {
			mv |= dMove
			break
		}
		mv |= dMove
	}
	return mv
}

// Piece Specific attacks

func (b *Board) kingAttacks(sq int) bitBoard {
	mv := bitBoard(0)
	mv |= kingAttacks[sq]
	return mv
}

func (b *Board) knightAttacks(sq int) bitBoard {
	mv := bitBoard(0)
	mv |= knightAttacks[sq]
	return mv
}

func (b *Board) rookAttacks(sq int) bitBoard {
	mv := bitBoard(0)
	mv |= b.strtAttacks(sq)
	return mv
}

func (b *Board) bishopAttacks(sq int) bitBoard {
	mv := bitBoard(0)
	mv |= b.diagAttacks(sq)
	return mv
}

func (b *Board) queenAttacks(sq int) bitBoard {
	mv := bitBoard(0)
	mv |= b.diagAttacks(sq) | b.strtAttacks(sq)
	return mv
}

func (b *Board) pawnAttacks(sq int, col color) bitBoard {
	var lAttack, rAttack bitBoard
	if col == WHITE {
		lAttack = (bitBoard(1<<sq) &^ FileA) << NW
		rAttack = (bitBoard(1<<sq) &^ FileH) << NE
	} else {
		lAttack = (bitBoard(1<<sq) &^ FileH) >> NW
		rAttack = (bitBoard(1<<sq) &^ FileA) >> NE
	}

	return lAttack | rAttack
}

// General Function to generate attacks for all pieces on a side
func (b *Board) genAttacks(col color) bitBoard {
	attack := bitBoard(0)
	for p := 0; p < 6; p++ {
		piece := b.bitBoard(col, p)
		for n := piece; n.countSet() != 0; n.nextSet() {
			sq := n.firstSet()
			switch p {
			case Pawn:
				attack |= b.pawnAttacks(sq, col)
			case Knight:
				attack |= b.knightAttacks(sq)
			case Bishop:
				attack |= b.bishopAttacks(sq)
			case Rook:
				attack |= b.rookAttacks(sq)
			case Queen:
				attack |= b.queenAttacks(sq)
			case King:
				attack |= b.kingAttacks(sq)
			}

		}
	}
	return attack
}

// to check if king is in check
func (b *Board) kingInCheck(col color) bool {
	return (b.bitBoard(col, King) & b.genAttacks(col.opp())) != 0
}

// Special checks
func (b *Board) pawnEnp(sq int, col color) bitBoard {
	mv := bitBoard(0)
	if b.enpassant != 0 {
		if col == WHITE {
			if (sq+NE == b.enpassant && sq%8 != 7) || (sq+NW == b.enpassant && sq%8 != 0) {
				mv |= bitBoard(1 << b.enpassant)
			}
		} else {
			if (sq-NE == b.enpassant && sq%8 != 0) || (sq-NW == b.enpassant && sq%8 != 7) {
				mv |= bitBoard(1 << b.enpassant)
			}
		}
	}
	return mv
}

// Moves for the pieces on the board
func (b *Board) knightMoves(sq int, color color) bitBoard {
	return b.knightAttacks(sq) & ^b.colors[color]
}

func (b *Board) bishopMoves(sq int, color color) bitBoard {
	return b.bishopAttacks(sq) & ^b.colors[color]
}

func (b *Board) queenMoves(sq int, color color) bitBoard {
	return b.queenAttacks(sq) & ^b.colors[color]
}

func (b *Board) rookMoves(sq int, color color) bitBoard {
	return b.rookAttacks(sq) & ^b.colors[color]
}

func (b *Board) kingMoves(sq int, color color) bitBoard {
	return b.kingAttacks(sq) & ^b.colors[color]
}

func (b *Board) pawnMoves(sq int, col color) bitBoard {
	var mv, one, two, p bitBoard
	p = bitBoard(1 << sq)
	allPieces := b.allPieces()

	if col == WHITE {
		one = (p << N) &^ allPieces
		two = ((one << N) & Rank4) &^ allPieces
	} else {
		one = (p >> N) &^ allPieces
		two = ((one >> N) & Rank5) &^ allPieces
	}
	mv = one | two
	mv |= b.pawnEnp(sq, col)
	return mv | (b.pawnAttacks(sq, col) & b.colors[col.opp()])
}

func (b *Board) legalMoves(piece, sq int, col color) bitBoard {
	mvs := bitBoard(0)
	switch piece {
	case Pawn:
		mvs = b.pawnMoves(sq, col)
	case Knight:
		mvs = b.knightMoves(sq, col)
	case Bishop:
		mvs = b.bishopMoves(sq, col)
	case Rook:
		mvs = b.rookMoves(sq, col)
	case Queen:
		mvs = b.queenMoves(sq, col)
	case King:
		mvs = b.kingMoves(sq, col)
	}
	return mvs
}

// to gen moves
func (b *Board) genMove(col color) moves {
	res := moves{}
	bb := b.colors[col]

	for piece := Pawn; piece <= King; piece++ {
		p := bb & b.pieces[piece]
		for p != 0 {
			fr := p.nextSet()
			mvs := b.legalMoves(piece, fr, col)
			for mvs != 0 {
				to := mvs.nextSet()
				frName := squareName[fr]
				toName := squareName[to]
				var flag int = mvQuiet

				// handling promotions for pawns
				if piece == Pawn && (bitBoard(1<<to)&Rank8 != 0 || bitBoard(1<<to)&Rank1 != 0) {
					prom := []string{"Q", "R", "N", "B"}
					for i := 0; i < 4; i++ {
						switch prom[i] {
						case "Q":
							flag = mvQProm
						case "R":
							flag = mvRProm
						case "N":
							flag = mvNProm
						case "B":
							flag = mvBProm
						}
						mv := move{
							fr:   fr,
							to:   to,
							flag: flag,
							str:  frName + toName + prom[i],
						}
						res.push(mv)
					}
					continue
				}
				mv := move{
					fr:   fr,
					to:   to,
					flag: flag,
					str:  frName + toName,
				}
				res.push(mv)
			}
		}
	}

	return res
}

// to gen legal moves
func (b *Board) stdMoves(col color) moves {
	var movs moves
	for _, mov := range b.genMove(col) {
		if err := b.move(mov.str); err == nil {
			if b.kingInCheck(col) {
				b.unmove()
				continue
			}
			b.unmove()
			movs = append(movs, mov)
		}
	}
	return movs
}

// to move
func (b *Board) move(str string) error {
	// log.Println(str)
	mv := b.stringToMove(str)
	var err error
	prev := state{
		square:    b.square,
		colors:    b.colors,
		pieces:    b.pieces,
		enpassant: b.enpassant,
		rule50:    b.rule50,
		castlings: b.castlings,
	}
	fr := mv.fr
	to := mv.to
	p12 := b.square[fr]
	col := WHITE

	if p12 < 0 {
		col = BLACK
	}
	p6 := abs(p12) - 1
	//TODO: moves validation
	//check to see if the from square is not empty
	if p12 == Empty {
		err = fmt.Errorf("from square is empty")
		return err
	}

	// check to see if the fr piece is of the right color
	if col != b.turn {
		err = fmt.Errorf("piece color is not of right color: %s turn", colors[col])
		return err
	}

	//check to see if move is available in piece moves
	if bitBoard(1<<to)&b.legalMoves(p6, fr, col) == 0 {
		err = fmt.Errorf("move is not legal")
		return err
	}

	if mv.flag == mvCastLn {
		if !b.castlingsCond(col, "L") {
			err = fmt.Errorf("%s cant castle long", colors[col])
			return err
		}
		a := A1
		d := D1
		rook := WRook
		if col == BLACK {
			a = A8
			d = D8
			rook = BRook
		}
		b.unset(a)
		b.set(rook, d)
	}

	if mv.flag == mvCastSh {
		if !b.castlingsCond(col, "S") {
			err = fmt.Errorf("%s cant castle short", colors[col])
			return err
		}
		a := H1
		d := F1
		rook := WRook
		if col == BLACK {
			a = H8
			d = F8
			rook = BRook
		}
		b.unset(a)
		b.set(rook, d)
	}

	if p6 == King {
		if col == WHITE {
			b.castlings.unset(longW)
			b.castlings.unset(shortW)
		} else {
			b.castlings.unset(longB)
			b.castlings.unset(shortB)
		}
	}

	if p6 == Rook {
		if fr == A1 {
			b.castlings.unset(longW)
		}
		if fr == H1 {
			// fmt.Println("I have lost short caslte")
			b.castlings.unset(shortW)
		}
		if fr == A8 {
			b.castlings.unset(shortB)
		}
		if fr == A8 {
			b.castlings.unset(shortB)
		}
	}

	b.unset(fr)
	b.set(p12, to)

	if mv.flag == mvEnp {
		sq := b.getEpCapSqr( /*fr,*/ col)
		// fmt.Println(squareName[sq])
		b.unset(sq)
	}

	// handling promotions
	if mv.flag == mvNProm || mv.flag == mvNPromCap {
		p6 := Knight
		p12 := p6 + 1
		if col == BLACK {
			p12 = -p12
		}
		b.set(p12, to)
	}
	if mv.flag == mvBProm || mv.flag == mvBPromCap {
		p6 := Bishop
		p12 := p6 + 1
		if col == BLACK {
			p12 = -p12
		}
		b.set(p12, to)
	}
	if mv.flag == mvRProm || mv.flag == mvRPromCap {
		p6 := Rook
		p12 := p6 + 1
		if col == BLACK {
			p12 = -p12
		}
		b.set(p12, to)
	}
	if mv.flag == mvQProm || mv.flag == mvQPromCap {
		p6 := Queen
		p12 := p6 + 1
		if col == BLACK {
			p12 = -p12
		}
		b.set(p12, to)
	}

	b.enpassant = 0
	if mv.flag&mvCapture != 0 || p6 == Pawn {
		b.rule50 = 0
	} else {
		b.rule50++
	}
	if b.isDoublePawnMove(fr, to, col) {
		b.enpassant = b.getEpSqr(to, col)
	}
	b.state.push(prev)
	b.history.push(mv)
	// incremeting the move no when it is black turn
	if col == BLACK {
		b.moveNo++
	}
	b.turn.flip()
	if b.kingInCheck(col) {
		err = fmt.Errorf("%s king in check", colors[col])
		b.unmove()
		return err
	}
	return nil
}

// unmove taking back a move
func (b *Board) unmove() error {
	prev, err := b.state.pop()
	if err != nil {
		return err
	}

	_, err = b.history.pop()
	if err != nil {
		return err
	}

	b.square = prev.square
	b.colors = prev.colors
	b.pieces = prev.pieces
	b.enpassant = prev.enpassant
	b.rule50 = prev.rule50
	b.castlings = prev.castlings
	if b.turn == WHITE {
		b.moveNo--
	}
	b.turn.flip()
	return nil
}
