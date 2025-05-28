package main

import (
	"fmt"
	"regexp"
)

// Technical note: this might be a bit confusing as met in code, so a quick explaination:
// a - is name actually in skip list
// b - is reverse mode active
//
// a      b        shouldSkip?
// true,  true  -> false
// true,  false -> true
// false, true  -> true
// false, false -> false
//
// So we use: a != b to check this cond.

func (c *Context) ShouldSkipFunc(f CIdentifier) bool {
	return c.Skip.Funcs.ShouldSkip(f) != c.preset.ReverseMode
}

func (c *Context) ShouldSkipTypedef(t CIdentifier) bool {
	return c.Skip.Typedefs.ShouldSkip(t) != c.preset.ReverseMode
}

func (c *Context) ShouldSkipStruct(t CIdentifier) bool {
	return c.Skip.Structs.ShouldSkip(t) != c.preset.ReverseMode
}

func (c *Context) ShouldSkipMethod(t CIdentifier) bool {
	return c.Skip.Methods.ShouldSkip(t) != c.preset.ReverseMode
}

func (c *Context) ShouldSkipFile(t string) bool {
	return c.Skip.Files.ShouldSkip(t) != c.preset.ReverseMode
}

// Skipper is an advanced method of skipping functions/types basing on a set of regular expressions.
// Regular expressions must be compatible with the `regexp` package in Go.
type Skipper[T ~string] struct {
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
