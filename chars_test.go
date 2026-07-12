package norm

import "testing"

func TestTrim(t *testing.T) {
	if got := Trim("  hello  ", Spaces); got != "hello" {
		t.Fatalf("Trim spaces = %q", got)
	}
	// Zero-width space and BOM at the edges.
	if got := Trim("\u200bhello\ufeff", Invisible); got != "hello" {
		t.Fatalf("Trim invisible = %q", got)
	}
	// Zero class is a no-op.
	if got := Trim("  x  ", 0); got != "  x  " {
		t.Fatalf("Trim zero class = %q", got)
	}
	if got := TrimLeft("--x--", Punct); got != "x--" {
		t.Fatalf("TrimLeft = %q", got)
	}
	if got := TrimRight("--x--", Punct); got != "--x" {
		t.Fatalf("TrimRight = %q", got)
	}
}

func TestRemoveAndKeep(t *testing.T) {
	// Remove zero-width space and BOM embedded in the middle.
	if got := Remove("he\u200bl\ufefflo", Invisible); got != "hello" {
		t.Fatalf("Remove invisible = %q", got)
	}
	if got := Remove("a b\tc\nd", Spaces); got != "abcd" {
		t.Fatalf("Remove spaces = %q", got)
	}
	if got := Remove("x", 0); got != "x" {
		t.Fatalf("Remove zero class = %q", got)
	}
	// Keep only letters and digits.
	if got := Keep("a1-b2 c3!", Letters|Digits); got != "a1b2c3" {
		t.Fatalf("Keep letters/digits = %q", got)
	}
}

func TestCollapseSpace(t *testing.T) {
	if got := CollapseSpace("  a \t b\n\nc  "); got != "a b c" {
		t.Fatalf("CollapseSpace = %q", got)
	}
	if got := CollapseSpace("nospace"); got != "nospace" {
		t.Fatalf("CollapseSpace nochange = %q", got)
	}
}

func TestLowerUpper(t *testing.T) {
	if Lower("ÀBc") != "àbc" {
		t.Fatalf("Lower = %q", Lower("ÀBc"))
	}
	if Upper("àbC") != "ÀBC" {
		t.Fatalf("Upper = %q", Upper("àbC"))
	}
}

func TestInvisibleClassCoverage(t *testing.T) {
	// A representative sample of zero-width / format characters.
	for _, r := range []rune{0x200B, 0xFEFF, 0x200C, 0x200D, 0x00AD, 0x2060, 0x034F, 0xFE0F} {
		if !isInvisible(r) {
			t.Errorf("rune U+%04X should be invisible", r)
		}
	}
	if isInvisible('a') || isInvisible(' ') {
		t.Fatal("false positive in isInvisible")
	}
}

func TestClean(t *testing.T) {
	// zero-width space (U+200B) + tab + collapsing + trim + a newline.
	in := "  Al​ice\tBoss \n(dev) "
	got := Clean(in)
	if got != "Alice Boss (dev)" {
		t.Fatalf("Clean = %q, want %q", got, "Alice Boss (dev)")
	}
	if Clean("   ") != "" {
		t.Errorf("Clean(all spaces) should be empty")
	}
}
