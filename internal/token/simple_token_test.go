package token

import (
	"testing"
)

var constInitialization = []byte("const zero = 0")
var strChar = []byte("\"abc kek \\n \\\" abc \"[2]")

func TestSimpleTokenKeywordParse(t *testing.T) {
	size, value := tokenKeyword.Parse(constInitialization)

	expectedValue := "const"
	expectedSize := len(expectedValue)

	if size != expectedSize {
		t.Errorf("got size %v, wanted %v", size, expectedSize)
	}

	if value != expectedValue {
		t.Errorf("got value %v, wanted %v", value, expectedValue)
	}
}

func TestSimpleTokenKeywordOnStringParse(t *testing.T) {
	size, value := tokenKeyword.Parse(strChar)

	expectedValue := ""
	expectedSize := 0

	if size != expectedSize {
		t.Errorf("got size %v, wanted %v", size, expectedSize)
	}

	if value != expectedValue {
		t.Errorf("got value %v, wanted %v", value, expectedValue)
	}
}

func TestSimpleTokenStringParse(t *testing.T) {
	size, value := tokenStringLiteral.Parse(strChar)

	expectedValue := "abc kek \\n \\\" abc "
	expectedSize := len(expectedValue) + 2

	if size != expectedSize {
		t.Errorf("got size %v, wanted %v", size, expectedSize)
	}

	if value != expectedValue {
		t.Errorf("got value %v, wanted %v", value, expectedValue)
	}
}

func TestSimpleTokenStringFailedParse(t *testing.T) {
	size, value := tokenStringLiteral.Parse(constInitialization)

	expectedValue := ""
	expectedSize := 0

	if size != expectedSize {
		t.Errorf("got size %v, wanted %v", size, expectedSize)
	}

	if value != expectedValue {
		t.Errorf("got value %v, wanted %v", value, expectedValue)
	}
}
