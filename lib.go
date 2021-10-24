package lecture2lib

import "unicode"

func ChangeCharCase(char rune) rune {
	if unicode.IsLower(char) {
		return unicode.ToUpper(char)
	} else {
		return unicode.ToUpper(char)
	}
}
