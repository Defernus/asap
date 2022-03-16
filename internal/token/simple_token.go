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
	match: *regexp.MustCompile("\\b(if|else|while|return|var|const|include|typedef)\\b"),
	name:  "keyword",
}

var tokenType = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("(\\b|(\\*|\\s)*)(int|char|void|float)\\b"),
	name:  "type",
}

var tokenStringLiteral = &simpleToken{
	begin: regexp.MustCompile("\""),
	end:   regexp.MustCompile("\""),
	match: *regexp.MustCompile("(\\\\\"|\\\\\\w|[^\\\"\\\\])*"),
	name:  "string-literal",
}

var tokenWhiteSpace = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("\\s"),
	name:  "white-space",
}

var tokenIdentifier = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("\\b[a-zA-Z_][\\w_]*"),
	name:  "identifier",
}

var tokenOperator = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("\\+\\+|\\+|\\-\\-|\\-|\\*|<=|>=|<|>|==|=|!|\\|\\||\\||&&|&|:"),
	name:  "operator",
}

var tokenInteger = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("(\\+|\\-)?\\d+"),
	name:  "integer",
}

var tokenSepparator = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile(";|,|\\."),
	name:  "sepparator",
}

var tokenOpenBracket = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("\\(|\\{|\\["),
	name:  "open-bracket",
}

var tokenCloseBracket = &simpleToken{
	begin: nil,
	end:   nil,
	match: *regexp.MustCompile("\\)|\\}|\\]"),
	name:  "close-bracket",
}
