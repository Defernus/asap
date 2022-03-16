package token

type Token interface {
	Parse([]byte) (size int, value string)
	GetName() string
}
