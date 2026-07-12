package norm

import (
	"strings"

	"github.com/goloop/is/v2"
)

// MAC normalizes a MAC address to lowercase, colon-separated form and reports
// whether the result is valid. It accepts the common separators (":", "-", ".")
// and the separatorless form, and regroups an EUI-48 (12 hex digits) or EUI-64
// (16 hex digits) into colon-separated pairs. No other punctuation is stripped.
func MAC(s string) (string, bool) {
	s = scrub(s)

	hex := strings.ToLower(stripAny(s, ":-. "))
	if (len(hex) == 12 || len(hex) == 16) && isHexString(hex) {
		if v := group(hex, 2, ':'); is.MAC(v) {
			return v, true
		}
	}

	// Fall back to a lowercased best-effort for anything else.
	v := strings.ToLower(s)
	return v, is.MAC(v)
}

// group inserts sep after every n bytes of s.
func group(s string, n int, sep byte) string {
	var b strings.Builder
	b.Grow(len(s) + len(s)/n)
	for i := 0; i < len(s); i++ {
		if i > 0 && i%n == 0 {
			b.WriteByte(sep)
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

// isHexString reports whether s is a non-empty run of lowercase hex digits.
func isHexString(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
			return false
		}
	}
	return true
}
