package token

import "asap/internal/source"

type Token struct {
	src       *source.Source
	offset    int
	tokenSize int
	col       int
	row       int
}

func getTokenPos(src string, offset int) (col, row int) {
	col = 0
	row = 0

	for pos, char := range src {
		if pos == offset {
			break
		}

		if char == '\n' {
			row = 0
			col++
		} else {
			row++
		}
	}

	return col, row
}

func NewToken(src *source.Source, offset, size int) *Token {
	col, row := getTokenPos(src.GetRawCode(), offset)

	return &Token{
		src:       src,
		offset:    offset,
		tokenSize: size,
		col:       col,
		row:       row,
	}
}
