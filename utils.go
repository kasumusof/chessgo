package chessgo

import (
	"fmt"
	"regexp"
	"strings"
)

var upper = func(text string) string {
	text = strings.ToUpper(text)
	return text
}
var lower = func(text string) string {
	text = strings.ToLower(text)
	return text
}

var trimSpaces = func(text string) string {
	text = strings.TrimSpace(text)
	space := regexp.MustCompile(`\s+`)
	text = space.ReplaceAllString(text, " ")
	return text
}

func drawBitBoard(bb bitBoard) {
	// start := time.Now()

	bitString := fmt.Sprintf("%064b", bb)
	var c, d string
	// fmt.Println("______Drawing_BitBoard_Started______")

	fmt.Println("    A   B   C   D   E   F   G   H")

	fmt.Println("  +---|---|---|---|---|---|---|---+")
	for i := 0; i < 8; i++ {
		// c = string(bitString[i*8 : (i+1)*8])
		c = string(bitString[i*8 : (i+1)*8])
		for j := 7; j >= 0; j-- {
			d = string(c[j])
			if j == 0 {
				fmt.Printf("%2s |%2v", d, 8-i)
				continue
			}
			if j == 7 {
				fmt.Printf("%1v |%2s |", 8-i, d)
				continue
			}
			fmt.Printf("%2s |", d)
		}
		fmt.Println()
		fmt.Println("  +---|---|---|---|---|---|---|---+")

	}
	fmt.Println("    A   B   C   D   E   F   G   H")
	// fmt.Println("_______Drawing_BitBoard_Ended_______\n", time.Since(start))

}

func drawBoard(b *Board) {
	// start := time.Now()
	var c []int
	var d string
	// fmt.Println("_______Drawing_Board_Started_______")

	fmt.Println("    A   B   C   D   E   F   G   H")

	fmt.Println("  +---|---|---|---|---|---|---|---+")
	for i := 7; i >= 0; i-- {
		c = b.square[i*8 : (i+1)*8]
		for j := 0; j < 8; j++ {
			d = string(display[c[j]])
			// d = string(c[i])
			// d = "p"

			if j == 7 {
				fmt.Printf("%2s |%2v", d, i+1)
				continue
			}
			if j == 0 {
				fmt.Printf("%1v |%2s |", i+1, d)
				continue
			}
			fmt.Printf("%2s |", d)
		}
		fmt.Println()
		fmt.Println("  +---|---|---|---|---|---|---|---+")

	}
	fmt.Println("    A   B   C   D   E   F   G   H")
	// fmt.Println("_______Drawing_Board_Ended_______\n", time.Since(start))

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func reverse(any []int) {
	for i, j := 0, len(any)-1; i < j; i, j = i+1, j-1 {
		any[i], any[j] = any[j], any[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
