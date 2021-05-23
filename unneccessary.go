package chessgo



func (b *Board) stringToMove(str string) move {
	fr := squareNameToInt[string(str[:2])]
	to := squareNameToInt[string(str[2:4])]
	mv := move{
		fr:   fr,
		to:   to,
		flag: 0,
		str:  str,
	}
	return mv
}

func (b *Board) castlingsCond(col color, t string) bool {
	if b.kingInCheck(col) {
		return false
	}
	allPieces := b.allPieces()
	attacks := b.genAttacks(col.opp())
	if col == WHITE {
		switch t {
		case "L":
			if b.castlings&longW == 0 {
				return false
			}
			if allPieces&bitBoard(1<<C1) != 0 {
				return false
			}
			if attacks&bitBoard(1<<C1) != 0 {
				return false
			}
			if allPieces&bitBoard(1<<D1) != 0 {
				return false
			}
			if attacks&bitBoard(1<<D1) != 0 {
				return false
			}
		case "S":
			if b.castlings&0b1000 == 0 {
				return false
			}
			if allPieces&bitBoard(1<<G1) != 0 {
				return false
			}
			if attacks&bitBoard(1<<G1) != 0 {
				return false
			}
			if allPieces&bitBoard(1<<F1) != 0 {
				return false
			}
			if attacks&bitBoard(1<<F1) != 0 {
				return false
			}
		}
	} else {
		switch t {
		case "L":
			if b.castlings&0b0001 == 0 {
				return false
			}
			if allPieces&bitBoard(1<<C8) != 0 {
				return false
			}
			if attacks&bitBoard(1<<C8) != 0 {
				return false
			}
			if allPieces&bitBoard(1<<D8) != 0 {
				return false
			}
			if attacks&bitBoard(1<<D8) != 0 {
				return false
			}
		case "S":
			if b.castlings&0b0010 == 0 {
				return false
			}
			if allPieces&bitBoard(1<<G8) != 0 {
				return false
			}
			if attacks&bitBoard(1<<G8) != 0 {
				return false
			}
			if allPieces&bitBoard(1<<F8) != 0 {
				return false
			}
			if attacks&bitBoard(1<<F8) != 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) isTypeEnpassant(fr, to int, col color) bool {
	var attack, opp bitBoard
	var a = 0
	if b.enpassant != to {
		return false
	}
	if col == WHITE {
		a = to + S
	} else {
		a = to + N
	}

	if a >= A1 && a <= H8 {
		opp = bitBoard(1 << a)
	}

	attack = opp & b.bitBoard(col.opp(), Pawn) & ^bitBoard(1<<to)
	return attack != 0
}

func (b *Board) isTypeCapture(piece, fr, to int, col color) bool {
	var attack bitBoard
	opp := b.colors[col.opp()] & bitBoard(1<<to)
	switch piece {
	case Pawn:
		attack = b.pawnMoves(fr, col)
		return b.isTypeEnpassant(fr, to, col) || attack&opp != 0
	default:
		attack = b.legalMoves(piece, fr, col)
	}
	return attack&opp != 0
}

func (b *Board) isTypeProm(fr, to int, col color) bool {

	if col == WHITE {
		return bitBoard(1<<to)&Rank8 != 0
	}
	return bitBoard(1<<to)&Rank1 != 0
}

func (b *Board) isDoublePawnMove(fr, to int, col color) bool {
	if col == WHITE {
		return fr/8 == 1 && fr+N+N == to
	}
	return fr/8 == 6 && fr+S+S == to
}

func (b *Board) isShortCaslte(fr, to int, col color) bool {
	if col == WHITE {
		return bitBoard(1<<fr)&Rank1 != 0 && bitBoard(1<<to)&FileG != 0
	}
	return bitBoard(1<<fr)&Rank8 != 0 && bitBoard(1<<to)&FileG != 0
}

func (b *Board) isLongCaslte(fr, to int, col color) bool {
	if col == WHITE {
		return bitBoard(1<<fr)&Rank1 != 0 && bitBoard(1<<to)&FileC != 0
	}
	return bitBoard(1<<fr)&Rank8 != 0 && bitBoard(1<<to)&FileC != 0
}


// get the square that is set for enpassant if a pawn move double
func (b *Board) getEpSqr(to int, col color) int {
	if col == WHITE {
		return to - N
	}
	return to + N
}

// get the square that the opponent pawn is on when enpassant was set
func (b *Board) getEpCapSqr(col color) int {
	if col == WHITE {
		return b.enpassant - N
	}
	return b.enpassant + N
}