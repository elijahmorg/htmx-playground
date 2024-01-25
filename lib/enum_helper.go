package lib

import (
	"strings"
	"unicode"
)

// MakeDisplayString converts input of form "word_next_word" into "Word next word"
func MakeDisplayString(input string) string {
	output := strings.ReplaceAll(input, "_", " ")
	r := []rune(output)
	r[0] = unicode.ToUpper(r[0])
	s := string(r)
	return s
}
