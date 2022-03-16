package token

import (
	"asap/internal/source"
	"errors"
)

var tokens = []Token{
	tokenKeyword,
	tokenStringLiteral,
	tokenWhiteSpace,
	tokenIdentifier,
	tokenOperator,
	tokenInteger,
	tokenSepparator,
	tokenOpenBracket,
	tokenCloseBracket,
}

func TokenizeSource(src *source.Source) []*TokenData {
	rawSrc := src.GetRawCode()
	offset := 0
	result := []*TokenData{}
	for offset < len(rawSrc) {
		subString := rawSrc[offset:]

		var tokenData *TokenData = nil
		for _, token := range tokens {
			tokenSize, tokenValue := token.Parse(subString)

			if tokenSize != 0 {
				col, row := getTokenPos(rawSrc, offset)
				tokenData = &TokenData{
					value:     tokenValue,
					tokenSize: tokenSize,
					token:     token,
					src:       src,
					offset:    offset,
					col:       col,
					row:       row,
				}
				break
			}
		}

		if tokenData == nil {
			col, row := getTokenPos(rawSrc, offset)
			src.ThrowError(col, row, errors.New("unknown token"))
		}
		result = append(result, tokenData)
		offset += tokenData.tokenSize
	}

	return result
}
