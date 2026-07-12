package norm

import "testing"

func TestIP(t *testing.T) {
	// IPv6 is canonicalized (lowercase, compressed).
	v, ok := IP("2001:0DB8:0000:0000:0000:0000:0000:0001")
	check(t, "ipv6 canonical", v, ok, "2001:db8::1", true)

	v, ok = IP(" 192.168.000.001 ")
	// Leading zeros are not valid in netip; the scrubbed value is returned as-is.
	if ok {
		t.Logf("note: %q accepted as %q", "192.168.000.001", v)
	}

	v, ok = IP("192.168.1.1")
	check(t, "ipv4", v, ok, "192.168.1.1", true)

	if _, ok := IP("not-an-ip"); ok {
		t.Error("invalid IP reported valid")
	}
}

func TestIPv4IPv6(t *testing.T) {
	if v, ok := IPv4("10.0.0.1"); !ok || v != "10.0.0.1" {
		t.Errorf("IPv4 = (%q, %v)", v, ok)
	}
	if v, ok := IPv6("FE80::1"); !ok || v != "fe80::1" {
		t.Errorf("IPv6 = (%q, %v)", v, ok)
	}
	// Cross-type: an IPv4 is not a valid IPv6.
	if _, ok := IPv6("10.0.0.1"); ok {
		t.Error("IPv4 accepted as IPv6")
	}
}

func TestDomainHostname(t *testing.T) {
	v, ok := Domain("  Example.COM ")
	check(t, "domain lower", v, ok, "example.com", true)

	v, ok = Hostname("MyHost.Local")
	check(t, "hostname lower", v, ok, "myhost.local", true)
}

func TestBase64(t *testing.T) {
	// Line-wrapped base64 has its whitespace removed.
	v, ok := Base64("aGVsbG8g\nd29ybGQ=")
	check(t, "base64 unwrap", v, ok, "aGVsbG8gd29ybGQ=", true)

	v, ok = Base64URL("aGVsbG8gd29ybGQ")
	if !ok {
		t.Errorf("Base64URL = (%q, %v)", v, ok)
	}
}

func TestHexColor(t *testing.T) {
	v, ok := HexColor("#AABBCC")
	check(t, "hexcolor lower", v, ok, "#aabbcc", true)

	// is.HexColor accepts a bare hex triplet, so norm mirrors that; a missing
	// "#" is never invented, only lowercased.
	if v, ok := HexColor("AABBCC"); v != "aabbcc" {
		t.Errorf("HexColor bare = (%q, %v)", v, ok)
	}
}

func TestFilterOnly(t *testing.T) {
	if got := AlphaOnly("a1b2-c3!"); got != "abc" {
		t.Errorf("AlphaOnly = %q", got)
	}
	if got := AlnumOnly("a1b2-c3!"); got != "a1b2c3" {
		t.Errorf("AlnumOnly = %q", got)
	}
	if got := DigitsOnly("+1 (415) 555-2671"); got != "14155552671" {
		t.Errorf("DigitsOnly = %q", got)
	}
}

func TestAlnumOnlyKeepsNumbers(t *testing.T) {
	// is.Alnum accepts any Unicode number; AlnumOnly must keep them too.
	if got := AlnumOnly("aⅣb"); got != "aⅣb" {
		t.Fatalf("AlnumOnly dropped a numeric: %q", got)
	}
	// DigitsOnly stays decimal-only (matches is.Digit).
	if got := DigitsOnly("aⅣb1"); got != "1" {
		t.Fatalf("DigitsOnly = %q", got)
	}
}

func TestIPv6ZoneNotCanonicalized(t *testing.T) {
	// is rejects zoned addresses; norm must return the input unchanged, not a
	// canonicalized-but-invalid value.
	v, ok := IPv6("FE80::1%eth0")
	if ok {
		t.Fatalf("zoned IPv6 reported valid: %q", v)
	}
	if v != "FE80::1%eth0" {
		t.Fatalf("zoned IPv6 was altered: %q", v)
	}
}
