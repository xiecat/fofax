package goflags

import (
	"strings"
)

// StringSlice is a slice of strings
type StringSlice []string

func (stringSlice StringSlice) String() string {
	return stringSlice.createStringArrayDefaultValue()
}

// Set appends a value to the string slice.
func (stringSlice *StringSlice) Set(value string) error {
	*stringSlice = append(*stringSlice, value)
	return nil
}

func (stringSlice *StringSlice) createStringArrayDefaultValue() string {
	defaultBuilder := &strings.Builder{}
	defaultBuilder.WriteString("[")
	for i, k := range *stringSlice {
		defaultBuilder.WriteString("\"")
		defaultBuilder.WriteString(k)
		defaultBuilder.WriteString("\"")
		if i != len(*stringSlice)-1 {
			defaultBuilder.WriteString(", ")
		}
	}
	defaultBuilder.WriteString("]")
	return defaultBuilder.String()
}
