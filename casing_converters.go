package problem

import (
	"strings"
	"unicode"
)

type StringCasingConverter func(value string) string

func splitIntoWords(value string) []string {
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

func capitalizeFirstChar(value string) string {
	if len(value) == 0 {
		return value
	}
	runes := []rune(value)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func SnakeCaseConverter(value string) string {
	words := splitIntoWords(value)
	return strings.Join(words, "_")
}

func CamelCaseConverter(value string) string {
	words := splitIntoWords(value)
	for i := range words {
		if i > 0 {
			words[i] = capitalizeFirstChar(words[i])
		}
	}
	return strings.Join(words, "")
}

func KebabCaseConverter(value string) string {
	words := splitIntoWords(value)
	return strings.Join(words, "-")
}

func PascalCaseConverter(value string) string {
	words := splitIntoWords(value)
	for i := range words {
		words[i] = capitalizeFirstChar(words[i])
	}
	return strings.Join(words, "")
}

func TrainCaseConverter(value string) string {
	words := splitIntoWords(value)
	for i := range words {
		words[i] = capitalizeFirstChar(words[i])
	}
	return strings.Join(words, "-")
}
