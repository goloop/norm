package norm

import "github.com/goloop/is/v2"

// Domain normalizes a domain name (scrub, trim and lowercase, since domains are
// case-insensitive) and reports whether the result is valid. It performs no
// IDN/punycode conversion.
func Domain(s string) (string, bool) {
	v := Lower(scrub(s))
	return v, is.Domain(v)
}

// Hostname normalizes a hostname (scrub, trim and lowercase) and reports
// whether the result is valid.
func Hostname(s string) (string, bool) {
	v := Lower(scrub(s))
	return v, is.Hostname(v)
}
