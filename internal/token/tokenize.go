package token

import (
	"asap/internal/source"
	"errors"
)

var tokenParsers = []func(string) *Token{}

func TokenizeSource(src *source.Source) []*Token {
	rawSrc := src.GetRawCode()
	offset := 0
	result := []*Token{}
	for offset < len(rawSrc) {
		subString := rawSrc[offset:]

		var token *Token = nil
		for _, parse := range tokenParsers {
			token = parse(subString)
			if token != nil {
				break
			}
		}

		if token == nil {
			col, row := getTokenPos(rawSrc, offset)
			src.ThrowError(col, row, errors.New("unknown token"))
		}
		result = append(result, token)
	}

	return result
}
