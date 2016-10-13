package releaseinfo

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/dlclark/regexp2"
)

func isNullOrWhiteSpace(s string) bool {
	return strings.TrimSpace(s) == ""
}

func isNonSpacingMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
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
