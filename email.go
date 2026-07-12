package norm

import (
	"strings"

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

// EmailFold is Email with the whole address lower-cased, for using an email as
// a case-insensitive identity: a login key, a uniqueness constraint or a
// deduplication key. Email preserves the local part's case because it is
// technically case-sensitive; EmailFold folds it too, matching how identity
// systems and mail providers treat addresses in practice, so "Boss@Acme.io"
// and "boss@acme.io" compare and store as the same key.
//
// Use Email when exact deliverability matters and EmailFold when the address is
// an identifier you look up or compare.
func EmailFold(s string) (string, bool) {
	v, ok := Email(s)
	return strings.ToLower(v), ok
}
