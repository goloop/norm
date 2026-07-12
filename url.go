package norm

import (
	"strings"

	"github.com/goloop/is/v2"
)

// URL normalizes a URL and reports whether the result is valid.
//
// It trims surrounding whitespace and lowercases the scheme and host only; the
// userinfo, path, query and fragment are left byte-for-byte untouched (no
// re-encoding). It never invents a missing scheme, so a schemeless input such
// as "example.com" stays invalid (and is returned as the best-effort value with
// false).
func URL(s string) (string, bool) {
	s = scrub(s)
	v := lowerSchemeHost(s)
	return v, is.URL(v)
}

// lowerSchemeHost lowercases the scheme and host of a "scheme://authority/..."
// URL by rewriting only those spans of the original string, leaving everything
// else verbatim. A string without "://" is returned unchanged.
func lowerSchemeHost(s string) string {
	i := strings.Index(s, "://")
	if i < 0 {
		return s
	}
	scheme := strings.ToLower(s[:i])
	rest := s[i+3:]

	// The authority ends at the first '/', '?' or '#'.
	end := len(rest)
	if j := strings.IndexAny(rest, "/?#"); j >= 0 {
		end = j
	}
	authority, tail := rest[:end], rest[end:]

	// Preserve userinfo (which may be case-sensitive); lowercase only the host.
	userinfo, host := "", authority
	if at := strings.LastIndexByte(authority, '@'); at >= 0 {
		userinfo, host = authority[:at+1], authority[at+1:]
	}

	return scheme + "://" + userinfo + strings.ToLower(host) + tail
}
