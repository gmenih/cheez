package engine

// FENString is a type alias for FENString
type FENString string

const (
	piecePlacement uint8 = iota
	sideToMove
	castlingAbility
	enPassantTarget
	halfMoveClock
	fullMoveCounter
)

// NewGameFENString represents a new game
const NewGameFENString FENString = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var piecesMap map[rune]Piece = map[rune]Piece{
	'b': Dark | Bishop,
	'k': Dark | King,
	'q': Dark | Queen,
	'n': Dark | Knight,
	'p': Dark | Pawn,
	'r': Dark | Rook,

	'B': Light | Bishop,
	'K': Light | King,
	'Q': Light | Queen,
	'N': Light | Knight,
	'P': Light | Pawn,
	'R': Light | Rook,
}

// Parse parses the FEN string into a Board
func (f FENString) Parse() Board {
	board := Board{}
	// op := piecePlacement
	var x, y uint8

	for _, r := range f {
		if r == ' ' {
			// lets stop here for now
			break
		}
		if r == '/' {
			x = 0
			y++
			continue
		}

		if r >= '1' && r <= '8' {
			v := uint8(r - '0')
			if v+x <= 8 {
				x += v
				continue
			} else {
				panic("WTF!")
			}
		}

		if v, ok := piecesMap[r]; ok {
			board[x][y] = v
			x++
		}
	}

	return board
}
