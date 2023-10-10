package decompress

import (
	"fmt"
	"strings"
)

var numbers = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func isDigit(val rune) bool {
	_, ok := numbers[val]
	return ok
}

func digitOf(val rune) int {
	res, _ := numbers[val]
	return res
}

func Decompress(src string) (string, error) {
	repr := []rune(src)
	var sb strings.Builder

	var count = 0
	var lastRune rune
	var firstLetter = true

	for i := 0; i < len(repr); {

		if repr[i] == '\\' {
			if !firstLetter {
				sb.WriteRune(lastRune)
				i++
			} else {
				firstLetter = false
			}
			lastRune = repr[i]
			i++
			continue
		}

		if !isDigit(repr[i]) {
			if !firstLetter {
				sb.WriteRune(lastRune)
			} else {
				firstLetter = false
			}
			lastRune = repr[i]
			i++
			continue
		}

		for i < len(repr) && isDigit(repr[i]) {
			count = count*10 + digitOf(repr[i])
			i++
		}

		if firstLetter {
			return "", fmt.Errorf("Expected char before number at %v", i)
		} else {
			for count > 0 {
				sb.WriteRune(lastRune)
				count--
			}
			firstLetter = true
		}
	}

	if !firstLetter {
		sb.WriteRune(lastRune)
	}
	return sb.String(), nil
}
