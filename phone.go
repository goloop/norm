package norm

import (
	"github.com/goloop/is/v2"
)

// E164 normalizes a phone number toward the E.164 format and reports whether the
// result is valid. It removes spaces, dashes, parentheses and dots. It is
// non-magical: it never guesses a country code, so a national number without a
// leading "+" stays invalid.
func E164(s string) (string, bool) {
	v := stripAny(scrub(s), " -().")
	return v, is.E164(v)
}

// Phone normalizes a phone number by removing spaces, parentheses, dashes and
// dots, and reports whether the result is valid. Like E164 it is non-magical: a
// number without a leading "+" and country code stays invalid.
func Phone(s string) (string, bool) {
	v := stripAny(scrub(s), " ()-.")
	return v, is.Phone(v)
}

// IMSI normalizes an International Mobile Subscriber Identity (15 digits) by
// removing spaces and dashes, and reports whether the result is valid.
func IMSI(s string) (string, bool) {
	v := stripAny(scrub(s), " -")
	return v, is.IMSI(v)
}
