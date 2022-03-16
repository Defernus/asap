package token

import (
	"regexp"
)

type simpleToken struct {
	begin *regexp.Regexp
	end   *regexp.Regexp
	match regexp.Regexp
	name  string
}

func (token *simpleToken) Parse(str []byte) (int, string) {
	start := 0
	if token.begin != nil {
		beginLoc := token.begin.FindIndex(str)
		if beginLoc == nil || beginLoc[0] != 0 {
			return 0, ""
		}

		start = beginLoc[1]
	}

	matchLoc := token.match.FindIndex(str[start:])
	if matchLoc == nil || matchLoc[0] != 0 {
		return 0, ""
	}

	end := matchLoc[1] + start

	size := end
	if token.end != nil {
		endLoc := token.end.FindIndex(str[end:])
		if endLoc == nil || endLoc[0] != 0 {
			return 0, ""
		}
		size = endLoc[1] + end
	}

	return size, string(str[start:end])
}

func (token *simpleToken) GetName() string {
	return token.name
}

var tokenKeyword = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("const|var"),
	name:  "keyword",
}

var tokenStringLiteral = &simpleToken{
	begin: regexp.MustCompile("\""),
	end:   regexp.MustCompile("\""),
	match: *regexp.MustCompile("(\\\\\"|\\\\\\w|[^\\\"\\\\])*"),
	name:  "string-literal",
}
