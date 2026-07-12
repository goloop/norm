package norm

import (
	"github.com/goloop/is/v2"
)

// BankCard normalizes a payment card number and reports whether the result is
// valid. It removes spaces and dashes (the usual grouping separators), then
// validates the number with the Luhn check and, when kinds are given, against
// those card kinds. No other characters are removed.
//
// Security: the returned value is a full card number (PAN). Do not log or store
// it unnecessarily, and remember that a Luhn-valid number is not proof of a
// real or intended card.
func BankCard(s string, kinds ...is.CardKind) (string, bool) {
	v := stripAny(scrub(s), " -")
	return v, is.BankCard(v, kinds...)
}
