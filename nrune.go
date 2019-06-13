package piscine

import"unicode/utf8"

func FirstRune(s string, n int) rune{
	ss,_:=utf8.DecodeRune([]byte(s[n+1:]))
	return ss
}
