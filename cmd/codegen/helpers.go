package main

import (
	"sort"
	"strings"
)

func HasPrefix[s ~string](str s, prefix string) bool {
	return strings.HasPrefix(string(str), prefix)
}

func TrimPrefix[s ~string](str s, prefix string) s {
	return s(strings.TrimPrefix(string(str), prefix))
}

func HasSuffix[s ~string](str s, suffix string) bool {
	return strings.HasSuffix(string(str), suffix)
}

func TrimSuffix[s ~string](str s, suffix string) s {
	return s(strings.TrimSuffix(string(str), suffix))
}

func ReplaceAll[s ~string](str s, old, new string) s {
	return s(strings.ReplaceAll(string(str), old, new))
}

func Contains[s ~string](str s, substr string) bool {
	return strings.Contains(string(str), substr)
}

func Split[s ~string](str s, sep string) []s {
	var ret []s
	for _, x := range strings.Split(string(str), sep) {
		ret = append(ret, s(x))
	}
	return ret
}

func Replace[s, ot, nt ~string](str s, old ot, new nt, n int) s {
	return s(strings.Replace(string(str), string(old), string(new), n))
}

func Join[s ~string](strs []s, sep string) s {
	return s(strings.Join(func() (ret []string) {
		for _, x := range strs {
			ret = append(ret, string(x))
		}
		return
	}(), sep))
}

func ContainsAny[s ~string](str s, chars string) bool {
	return strings.ContainsAny(string(str), chars)
}

func Index[s ~string](str s, substr string) int {
	return strings.Index(string(str), substr)
}

func Capitalize[s ~string](str s) s {
	return s(strings.ToUpper(string(str[0]))) + str[1:]
}

func SortStrings[s ~string](str []s) {
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
}
