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
