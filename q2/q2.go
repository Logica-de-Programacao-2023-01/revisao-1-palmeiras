package q2

import (
	"errors"
	"strings"
	"unicode"
)

func AverageLettersPerWord(text string) (float64, error) {
	if strings.TrimSpace(text) == "" {
		return 0, errors.New("text cannot be empty")
	}

	var wordCount, letterCount int
	inWord := false
	for _, c := range text {
		if unicode.IsLetter(c) {
			if !inWord {
				inWord = true
				wordCount++
			}
			letterCount++
		} else {
			inWord = false
		}
	}

	if wordCount == 0 {
		return 0, errors.New("text must contain at least one word")
	}

	return float64(letterCount) / float64(wordCount), nil
}
