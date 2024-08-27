package caseconverter

import (
	"strings"
	"unicode"
)

var ResponseKeyCasingConverter StringCasingConverter = SnakeCaseConverter

type StringCasingConverter func(value string) string

func SplitIntoWords(value string) []string {
	var words []string
	var currentWord []rune
	for _, r := range value {
		if unicode.IsUpper(r) || unicode.IsSpace(r) || r == '-' || r == '_' || r == '.' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{}
			}
		}
		if r != '-' && !unicode.IsSpace(r) {
			currentWord = append(currentWord, unicode.ToLower(r))
		}
	}
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}

func CapitalizeFirstChar(value string) string {
	if len(value) == 0 {
		return value
	}
	runes := []rune(value)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func SnakeCaseConverter(value string) string {
	words := SplitIntoWords(value)
	return strings.Join(words, "_")
}

func CamelCaseConverter(value string) string {
	words := SplitIntoWords(value)
	for i := range words {
		if i > 0 {
			words[i] = CapitalizeFirstChar(words[i])
		}
	}
	return strings.Join(words, "")
}

func KebabCaseConverter(value string) string {
	words := SplitIntoWords(value)
	return strings.Join(words, "-")
}

func PascalCaseConverter(value string) string {
	words := SplitIntoWords(value)
	for i := range words {
		words[i] = CapitalizeFirstChar(words[i])
	}
	return strings.Join(words, "")
}

func TrainCaseConverter(value string) string {
	words := SplitIntoWords(value)
	for i := range words {
		words[i] = CapitalizeFirstChar(words[i])
	}
	return strings.Join(words, "-")
}
