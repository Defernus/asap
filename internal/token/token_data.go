package token

import (
	"asap/internal/source"
	"fmt"
)

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

func (tokenData *TokenData) GetPath() string {
	return fmt.Sprintf("%v:%v:%v", tokenData.src.GetSourcePath(), tokenData.col+1, tokenData.row+1)
}

func (tokenData *TokenData) String() string {
	return fmt.Sprintf(
		"%v at %v (%v size) \"%v\"",
		tokenData.token.GetName(),
		tokenData.GetPath(),
		tokenData.tokenSize,
		tokenData.value,
	)
}

func (tokenData *TokenData) GetCol() int {
	return tokenData.col
}

func (tokenData *TokenData) GetRow() int {
	return tokenData.row
}

func (tokenData *TokenData) GetOffset() int {
	return tokenData.offset
}

func (tokenData *TokenData) GetValue() string {
	return tokenData.value
}

func (tokenData *TokenData) GetToken() Token {
	return tokenData.token
}
