package norm

import "testing"

func TestPhone(t *testing.T) {
	v, ok := Phone("+1 (415) 555-2671")
	check(t, "phone formatting", v, ok, "+14155552671", true)
	if _, ok := Phone("415 555 2671"); ok {
		t.Error("national number reported valid phone")
	}
}

func TestIMSI(t *testing.T) {
	v, ok := IMSI("310 150 123456789")
	check(t, "imsi", v, ok, "310150123456789", true)
	if _, ok := IMSI("12345"); ok {
		t.Error("short IMSI reported valid")
	}
}

func TestBin(t *testing.T) {
	v, ok := Bin("0b1010 1010")
	check(t, "bin unwrap", v, ok, "0b10101010", true)

	v, ok = Bin("0B1010")
	check(t, "bin lowercase prefix", v, ok, "0b1010", true)

	if _, ok := Bin("1012"); ok {
		t.Error("non-binary reported valid")
	}
}

func TestDecimal(t *testing.T) {
	v, ok := Decimal("1 234 567")
	check(t, "decimal ungroup", v, ok, "1234567", true)
	if _, ok := Decimal("12a"); ok {
		t.Error("non-decimal reported valid")
	}
}

func TestFloat(t *testing.T) {
	v, ok := Float("  -3.14 ")
	check(t, "float trim", v, ok, "-3.14", true)
	// Grammar is not repaired: an internal space is not removed.
	if _, ok := Float("3. 14"); ok {
		t.Error("spaced float reported valid")
	}
}

func TestRGBColor(t *testing.T) {
	v, ok := RGBColor("RGB(255, 0, 0)")
	check(t, "rgb lower keyword", v, ok, "rgb(255, 0, 0)", true)

	v, ok = RGBColor("Rgb(0,0,0)")
	check(t, "rgb mixed case", v, ok, "rgb(0,0,0)", true)

	if _, ok := RGBColor("rgb(300,0,0)"); ok {
		t.Error("out-of-range RGB reported valid")
	}
}

func TestScrubWhitespaceConsistency(t *testing.T) {
	// A space and a newline inside a float must be treated the same (both make
	// it invalid): scrub no longer deletes interior control whitespace.
	if _, ok := Float("3.\n14"); ok {
		t.Error("Float repaired an interior newline")
	}
	if _, ok := Float("3. 14"); ok {
		t.Error("Float accepted an interior space")
	}
	// Decimal removes all whitespace grouping (space and newline alike).
	if v, ok := Decimal("1\n2"); !ok || v != "12" {
		t.Errorf("Decimal newline group = (%q, %v)", v, ok)
	}
	// A non-whitespace control byte is still scrubbed everywhere.
	if v, ok := Email("user\x00@example.com"); !ok || v != "user@example.com" {
		t.Errorf("Email control scrub = (%q, %v)", v, ok)
	}
}
