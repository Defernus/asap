package source

import (
	"fmt"
	"os"
	"strings"
)

func (src *Source) ThrowError(col, row int, err error) {
	lines := strings.Split(src.GetRawCode(), "\n")
	line := lines[col]
	fmt.Printf("at %v:%v:%v error: %v \n%v\n", src.filePath, col, row, err, line)
	fmt.Printf("%v^\n", strings.Repeat(" ", row))
	os.Exit(1)
}
