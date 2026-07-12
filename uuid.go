package norm

import (
	"strings"

	"github.com/goloop/is/v2"
)

// UUID normalizes a UUID and reports whether the result is valid. It trims
// surrounding whitespace, strips a surrounding pair of braces and a leading
// "urn:uuid:" prefix, and lowercases the value. It does not remove any internal
// characters other than the recognized wrappers.
func UUID(s string) (string, bool) {
	s = scrub(s)
	if len(s) >= 2 && strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		s = s[1 : len(s)-1]
	}
	if len(s) >= 9 && strings.EqualFold(s[:9], "urn:uuid:") {
		s = s[9:]
	}
	s = strings.ToLower(s)
	return s, is.UUID(s)
}
