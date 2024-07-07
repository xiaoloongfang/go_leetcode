package main

import "unicode"

func detectCapitalUse(word string) bool {
	return checkAllLowerOrUpper(word, true) ||
		checkAllLowerOrUpper(word, false) ||
		(unicode.IsUpper(rune(word[0])) && checkAllLowerOrUpper(word[1:], true))
}

func checkAllLowerOrUpper(word string, lower bool) bool {
	for _, c := range word {
		if (unicode.IsLower(c) && lower) || (unicode.IsUpper(c) && !lower) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {

}
