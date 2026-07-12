package norm

import (
	"strings"
	"unicode"
)

// Class is a bitmask selecting one or more groups of characters. Combine groups
// with the bitwise OR operator, for example Spaces|Controls.
type Class uint

const (
	// Spaces are Unicode White_Space characters, including the regular space,
	// tab, newline and no-break space (U+00A0).
	Spaces Class = 1 << iota

	// Controls are Unicode control characters (category Cc).
	Controls

	// Invisible are zero-width and formatting characters that render to nothing
	// yet can corrupt input or hide content: Unicode format characters
	// (category Cf, such as the zero-width space, byte-order mark, soft hyphen
	// and word joiner), variation selectors and the combining grapheme joiner.
	Invisible

	// Punct are Unicode punctuation characters (category P).
	Punct

	// Symbols are Unicode symbol characters (category S), such as math and
	// currency symbols.
	Symbols

	// Letters are Unicode letters (category L).
	Letters

	// Digits are Unicode decimal digits (category Nd), matching is.Digit.
	Digits

	// Numbers are all Unicode numbers (category N: decimal digits, letter
	// numbers such as Roman numerals, and other numerics), matching the numeric
	// half of is.Alnum.
	Numbers
)

// contains reports whether r belongs to any group selected by c.
func (c Class) contains(r rune) bool {
	switch {
	case c&Spaces != 0 && unicode.IsSpace(r):
		return true
	case c&Controls != 0 && unicode.IsControl(r):
		return true
	case c&Invisible != 0 && isInvisible(r):
		return true
	case c&Punct != 0 && unicode.IsPunct(r):
		return true
	case c&Symbols != 0 && unicode.IsSymbol(r):
		return true
	case c&Letters != 0 && unicode.IsLetter(r):
		return true
	case c&Digits != 0 && unicode.IsDigit(r):
		return true
	case c&Numbers != 0 && unicode.IsNumber(r):
		return true
	}
	return false
}

// isInvisible reports whether r is a zero-width or formatting character.
func isInvisible(r rune) bool {
	if unicode.Is(unicode.Cf, r) {
		return true
	}
	switch {
	case r == 0x034F: // combining grapheme joiner
		return true
	case r >= 0x180B && r <= 0x180F: // Mongolian free variation selectors
		return true
	case r >= 0xFE00 && r <= 0xFE0F: // variation selectors
		return true
	case r >= 0xE0100 && r <= 0xE01EF: // variation selectors supplement
		return true
	}
	return false
}

// Trim removes characters of the given class from both ends of s. A zero class
// leaves s unchanged.
func Trim(s string, class Class) string {
	if class == 0 {
		return s
	}
	return strings.TrimFunc(s, class.contains)
}

// TrimLeft removes characters of the given class from the start of s.
func TrimLeft(s string, class Class) string {
	if class == 0 {
		return s
	}
	return strings.TrimLeftFunc(s, class.contains)
}

// TrimRight removes characters of the given class from the end of s.
func TrimRight(s string, class Class) string {
	if class == 0 {
		return s
	}
	return strings.TrimRightFunc(s, class.contains)
}

// Remove deletes every character of the given class from s. A zero class
// leaves s unchanged.
func Remove(s string, class Class) string {
	if class == 0 {
		return s
	}
	return strings.Map(func(r rune) rune {
		if class.contains(r) {
			return -1
		}
		return r
	}, s)
}

// Keep deletes every character NOT in the given class from s. It is the inverse
// of Remove and is the honest way to filter to, say, letters: Keep(s, Letters).
func Keep(s string, class Class) string {
	return strings.Map(func(r rune) rune {
		if class.contains(r) {
			return r
		}
		return -1
	}, s)
}

// CollapseSpace trims surrounding whitespace and replaces every internal run of
// whitespace with a single ASCII space.
func CollapseSpace(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	pendingSpace := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			pendingSpace = b.Len() > 0
			continue
		}
		if pendingSpace {
			b.WriteByte(' ')
			pendingSpace = false
		}
		b.WriteRune(r)
	}
	return b.String()
}

// scrub is the shared entry cleaning for the typed normalizers: it removes
// control and invisible characters (never valid inside these formats) and trims
// surrounding whitespace.
func scrub(s string) string {
	return strings.TrimSpace(Remove(s, Controls|Invisible))
}
