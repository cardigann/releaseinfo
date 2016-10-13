package releaseinfo

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/dlclark/regexp2"
)

func isNullOrWhiteSpace(s string) bool {
	return removeSpace(s) == ""
}

func isNonSpacingMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func removeSpace(s string) string {
	return strings.TrimSpace(strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || r == '\x00' {
			return ' '
		}
		return r
	}, s))
}

func removeAccent(s string) string {
	b := make([]byte, len(s))
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isNonSpacingMark), norm.NFC)

	if _, _, e := t.Transform(b, []byte(s), true); e != nil {
		return s
	}

	return string(b)
}

func optionalReplace(r *regexp2.Regexp, from, to string) string {
	result, err := r.Replace(from, to, 0, -1)
	if err != nil {
		return from
	}
	return result
}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
