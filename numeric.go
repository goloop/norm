package norm

import "github.com/goloop/is/v2"

// Bin normalizes a binary string by removing whitespace (binary is often
// space-grouped) and lowercasing, so an uppercase "0B" prefix becomes "0b", and
// reports whether the result is valid.
func Bin(s string) (string, bool) {
	v := Lower(Remove(scrub(s), Spaces))
	return v, is.Bin(v)
}

// Decimal normalizes a run of decimal digits by removing whitespace (digits are
// sometimes space-grouped) and reports whether the result is valid.
func Decimal(s string) (string, bool) {
	v := Remove(scrub(s), Spaces)
	return v, is.Decimal(v)
}

// Float normalizes a floating-point literal and reports whether the result is
// valid. It only scrubs and trims: the sign and decimal point are grammar, so
// no characters are removed from the middle (which could silently change the
// value).
func Float(s string) (string, bool) {
	v := scrub(s)
	return v, is.Float(v)
}
