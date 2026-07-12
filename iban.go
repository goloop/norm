package norm

import (
	"github.com/goloop/is/v2"
)

// IBAN normalizes an International Bank Account Number and reports whether the
// result is valid. It removes spaces and uppercases the value (an IBAN is
// defined over uppercase letters and digits), then validates it strictly,
// including the checksum. No other characters are removed.
func IBAN(s string) (string, bool) {
	v := Upper(Remove(scrub(s), Spaces))
	return v, is.IBAN(v, true)
}
