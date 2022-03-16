package main

import (
	"asap/internal/source"
	"asap/internal/token"
	"fmt"
)

func main() {
	src, err := source.ReadSource("test.asap")
	if err != nil {
		panic(err)
	}

	tokens := token.TokenizeSource(src)

	for _, token := range tokens {
		fmt.Printf("token: %s\n", token)
	}
}
