package norm

import (
	"github.com/goloop/is/v2"
)

// Email normalizes an email address and reports whether the result is valid.
//
// It trims surrounding whitespace and lowercases the domain (domains are
// case-insensitive; the local part is preserved, as it is technically
// case-sensitive). If the address is still invalid, it removes internal
// whitespace as a best-effort repair. A valid input, including a quoted local
// part that legitimately contains spaces, is preserved because the repair runs
// only when the milder result does not validate.
//
// It does not perform IDN/punycode conversion: a Unicode domain is accepted
// only as far as is.Email accepts it.
func Email(s string) (string, bool) {
	s = scrub(s)
	if v := lowerDomain(s); is.Email(v) {
		return v, true
	}
	v := lowerDomain(Remove(s, Spaces))
	return v, is.Email(v)
}
