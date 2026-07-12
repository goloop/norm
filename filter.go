package norm

// This file provides the honest, lossy counterparts of the character-class
// predicates of is (Alpha, Alnum, Digit). They are intentionally named "*Only"
// and return a plain string, not (string, bool): there is deliberately no
// canonicalizing norm.Alpha/norm.Digit, because "normalizing to letters" is not
// canonicalization but filtering, and a misleading name is the real hazard. For
// arbitrary groups use Keep and Remove directly.

// AlphaOnly returns s with everything but Unicode letters removed. It is the
// filtering counterpart of is.Alpha.
func AlphaOnly(s string) string { return Keep(s, Letters) }

// AlnumOnly returns s with everything but Unicode letters and numbers removed.
// It is the filtering counterpart of is.Alnum, which accepts any Unicode letter
// or number (so, like is.Alnum, it keeps Roman numerals and other numerics, not
// only decimal digits).
func AlnumOnly(s string) string { return Keep(s, Letters|Numbers) }

// DigitsOnly returns s with everything but decimal digits removed. It is the
// filtering counterpart of is.Digit.
func DigitsOnly(s string) string { return Keep(s, Digits) }
