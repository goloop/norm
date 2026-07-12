package norm

import (
	"github.com/goloop/is/v2"
)

// Hex normalizes a hexadecimal string (trim and lowercase) and reports whether
// the result is valid hex.
func Hex(s string) (string, bool) { return lowerTrim(s, is.Hex) }

// MD5 normalizes an MD5 digest (trim and lowercase) and reports validity.
func MD5(s string) (string, bool) { return lowerTrim(s, is.MD5) }

// SHA1 normalizes a SHA-1 digest (trim and lowercase) and reports validity.
func SHA1(s string) (string, bool) { return lowerTrim(s, is.SHA1) }

// SHA256 normalizes a SHA-256 digest (trim and lowercase) and reports validity.
func SHA256(s string) (string, bool) { return lowerTrim(s, is.SHA256) }

// SHA512 normalizes a SHA-512 digest (trim and lowercase) and reports validity.
func SHA512(s string) (string, bool) { return lowerTrim(s, is.SHA512) }

// lowerTrim trims surrounding whitespace, lowercases s and validates it with ok.
// It suits hexadecimal digests, whose canonical form is lowercase.
func lowerTrim(s string, ok func(string) bool) (string, bool) {
	v := Lower(scrub(s))
	return v, ok(v)
}
