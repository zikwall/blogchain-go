package utils

import "strings"

// It's possible that your "\n" is actually the escaped version of a line break character.
// You can replace these with real line breaks by searching for the escaped version
// and replacing with the non escaped version
func EscapeNewLine(s string) string {
	return strings.ReplaceAll(s, `\n`, "\n")
}
