package norm

import "github.com/goloop/is/v2"

// Base64 normalizes a standard base64 string by removing whitespace (base64 is
// often line-wrapped) and reports whether the result is valid. Padding and the
// base64 alphabet are preserved.
func Base64(s string) (string, bool) {
	v := Remove(scrub(s), Spaces)
	return v, is.Base64(v)
}

// Base64URL normalizes a URL-safe base64 string by removing whitespace and
// reports whether the result is valid.
func Base64URL(s string) (string, bool) {
	v := Remove(scrub(s), Spaces)
	return v, is.Base64URL(v)
}

// JWT normalizes a JSON Web Token (scrub and trim; a JWT contains no
// whitespace) and reports whether the result is well-formed.
func JWT(s string) (string, bool) {
	v := scrub(s)
	return v, is.JWT(v)
}

// HexColor normalizes a hexadecimal color (scrub, trim and lowercase) and
// reports whether the result is valid. A missing "#" is not added.
func HexColor(s string) (string, bool) {
	v := Lower(scrub(s))
	return v, is.HexColor(v)
}

// RGBColor normalizes a CSS rgb(...) color by lowercasing the "rgb" keyword (the
// numbers and separators are unaffected) and reports whether the result is
// valid.
func RGBColor(s string) (string, bool) {
	v := Lower(scrub(s))
	return v, is.RGBColor(v)
}
