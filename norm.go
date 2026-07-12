// Package norm normalizes user-entered strings, the writing companion to the
// predicate module github.com/goloop/is: where is reads (is this an email?),
// norm cleans a value toward the canonical form of that type and validates the
// result.
//
// # Typed normalizers
//
// Each typed normalizer mirrors an is predicate and returns the cleaned value
// together with whether it is valid:
//
//	value, ok := norm.Email("  Example @ Gmail.com ")
//	// value == "Example@gmail.com", ok == true
//
// The returned string is the best-effort normalized value even when ok is
// false, so a UI can offer a "did you mean" suggestion. Covered types: Email,
// URL, UUID, MAC, IBAN, BankCard, E164 and the hex/hash family (Hex, MD5, SHA1,
// SHA256, SHA512).
//
// Normalization is conservative and type-specific. Every value is first
// scrubbed of control and invisible characters (never valid inside these
// formats); safe canonical steps (trim, lowercasing a case-insensitive domain
// or scheme) are always applied; a more aggressive repair (removing internal
// whitespace) is only attempted when the value is still invalid. Every result
// is confirmed with an is predicate, so a value is never reported valid unless
// is agrees. norm never strips arbitrary punctuation and never invents data
// (for example a URL scheme or a phone country code).
//
// # Character toolkit
//
// For the cases that are not a fixed external format, norm offers character
// group tools driven by a Class bitmask:
//
//	norm.Trim(s, norm.Spaces|norm.Controls) // remove groups from both ends
//	norm.Remove(s, norm.Invisible)          // remove a group everywhere
//	norm.Keep(s, norm.Letters|norm.Digits)  // keep only these groups
//	norm.CollapseSpace(s)                   // runs of whitespace -> one space
//	norm.Lower(s), norm.Upper(s)            // Unicode case
//
// These deliberately replace mirrors like a hypothetical norm.Alpha or
// norm.Digit, which would imply a canonical form but really perform lossy
// filtering; Keep and Remove make that explicit.
//
// See DOC.md (English) and DOC.UK.md (Ukrainian) for the full reference.
package norm

import "strings"

// stripAny removes every rune contained in cutset from s.
func stripAny(s, cutset string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(cutset, r) {
			return -1
		}
		return r
	}, s)
}

// lowerDomain lowercases the domain part of an email (after the last '@'),
// leaving the local part untouched. It returns s unchanged when there is no '@'.
func lowerDomain(s string) string {
	i := strings.LastIndexByte(s, '@')
	if i < 0 {
		return s
	}
	return s[:i+1] + strings.ToLower(s[i+1:])
}
