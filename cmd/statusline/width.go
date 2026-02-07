// Package main provides terminal width calculation utilities
package main

import "github.com/mattn/go-runewidth"

// StringWidth returns the terminal display width of a string
// It correctly handles emoji (2 cells), CJK characters (2 cells),
// and combining characters (0 cells) which are counted incorrectly
// by len([]rune()) or utf8.RuneCountInString()
func StringWidth(s string) int {
	return runewidth.StringWidth(s)
}
