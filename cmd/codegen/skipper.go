package main

import (
	"fmt"
	"regexp"
)

// Skipper is an advanced method of skipping functions/types basing on a set of regular expressions.
// Regular expressions must be compatible with the `regexp` package in Go.
type Skipper struct {
	data []regexp.Regexp
}

func CreateSkipper[T ~string](data []T) (*Skipper[T], error) {
	if data == nil {
		return &Skipper[T]{}, nil
	}

	result := &Skipper[T]{data: make([]regexp.Regexp, len(data))}

	for i, d := range data {
		re, err := regexp.Compile(string(d))
		if err != nil {
			return nil, fmt.Errorf("Cannot compile regular expression #%d \"%s\": %w", i, d, err)
		}

		result.data[i] = *re
	}

	return result, nil
}

func (s *Skipper[T]) ShouldSkip(name T) bool {
	for _, re := range s.data {
		if re.MatchString(string(name)) {
			return true
		}
	}

	return false
}
