package token

import "asap/internal/source"

type TokenData struct {
	value     string
	token     Token
	src       *source.Source
	offset    int
	tokenSize int
	col       int
	row       int
}

func getTokenPos(src []byte, offset int) (col, row int) {
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

func NewToken(src *source.Source, offset, size int) *TokenData {
	col, row := getTokenPos(src.GetRawCode(), offset)

	return &TokenData{
		src:       src,
		offset:    offset,
		tokenSize: size,
		col:       col,
		row:       row,
	}
}
