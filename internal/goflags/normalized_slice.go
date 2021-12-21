package goflags

import (
	"strings"

	"github.com/pkg/errors"
)

var quotes = []rune{'"', '\'', '`'}

// NormalizedStringSlice is a slice of strings
type NormalizedStringSlice []string

func (normalizedStringSlice NormalizedStringSlice) String() string {
	return normalizedStringSlice.createStringArrayDefaultValue()
}

//Set appends a value to the string slice.
func (normalizedStringSlice *NormalizedStringSlice) Set(value string) error {
	if slice, err := ToNormalizedStringSlice(value); err != nil {
		return err
	} else {
		*normalizedStringSlice = append(*normalizedStringSlice, slice...)
		return nil
	}
}

func (normalizedStringSlice *NormalizedStringSlice) createStringArrayDefaultValue() string {
	defaultBuilder := &strings.Builder{}
	defaultBuilder.WriteString("[")
	for i, k := range *normalizedStringSlice {
		defaultBuilder.WriteString("\"")
		defaultBuilder.WriteString(k)
		defaultBuilder.WriteString("\"")
		if i != len(*normalizedStringSlice)-1 {
			defaultBuilder.WriteString(", ")
		}
	}
	defaultBuilder.WriteString("]")
	return defaultBuilder.String()
}

func ToNormalizedStringSlice(value string) ([]string, error) {
	var result []string

	addPartToResult := func(part string) {
		if strings.TrimSpace(part) != "" {
			result = append(result, strings.TrimSpace(strings.Trim(strings.TrimSpace(strings.ToLower(part)), string(quotes))))
		}
	}

	index := 0
	for index < len(value) {
		char := rune(value[index])
		if isQuote, quote := isQuote(char); isQuote {
			quoteFound, part := searchPart(value[index+1:], quote)

			if !quoteFound {
				return nil, errors.New("Unclosed quote in path")
			}

			index += len(part) + 2

			addPartToResult(part)
		} else {
			commaFound, part := searchPart(value[index:], ',')

			if commaFound {
				index += len(part) + 1
			} else {
				index += len(part)
			}

			addPartToResult(part)
		}
	}

	return result, nil
}

func isQuote(char rune) (bool, rune) {
	for _, quote := range quotes {
		if quote == char {
			return true, quote
		}
	}
	return false, 0
}

func searchPart(value string, stop rune) (bool, string) {
	var result string
	for _, char := range value {
		if char != stop {
			result += string(char)
		} else {
			return true, result
		}
	}
	return false, result
}
