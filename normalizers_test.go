package norm

import "testing"

// check is a small helper for (value, ok) assertions.
func check(t *testing.T, name, gotV string, gotOK bool, wantV string, wantOK bool) {
	t.Helper()
	if gotV != wantV || gotOK != wantOK {
		t.Errorf("%s = (%q, %v), want (%q, %v)", name, gotV, gotOK, wantV, wantOK)
	}
}

func TestEmail(t *testing.T) {
	v, ok := Email("  user@example.com ")
	check(t, "trim", v, ok, "user@example.com", true)

	v, ok = Email("User@EXAMPLE.COM")
	check(t, "lowercase domain, keep local", v, ok, "User@example.com", true)

	v, ok = Email("example @ gmail.com  ")
	check(t, "space repair", v, ok, "example@gmail.com", true)

	// A zero-width space inside the address is scrubbed before validation.
	v, ok = Email("user@exa\u200bmple.com")
	check(t, "invisible scrub", v, ok, "user@example.com", true)

	v, ok = Email("not-an-email")
	if ok {
		t.Errorf("invalid email reported valid: %q", v)
	}
}

func TestURL(t *testing.T) {
	v, ok := URL("HTTPS://Example.COM/Path?Q=1")
	check(t, "lowercase scheme+host, keep path/query", v, ok, "https://example.com/Path?Q=1", true)

	// No scheme is not invented.
	v, ok = URL("example.com")
	if ok {
		t.Errorf("schemeless URL reported valid: %q", v)
	}
}

func TestUUID(t *testing.T) {
	const canon = "550e8400-e29b-41d4-a716-446655440000"
	v, ok := UUID("  {550E8400-E29B-41D4-A716-446655440000}  ")
	check(t, "braces+upper", v, ok, canon, true)

	v, ok = UUID("urn:uuid:550e8400-e29b-41d4-a716-446655440000")
	check(t, "urn prefix", v, ok, canon, true)
}

func TestMAC(t *testing.T) {
	const canon = "00:11:22:33:44:55"
	for _, in := range []string{
		"00:11:22:33:44:55",
		"00-11-22-33-44-55",
		"0011.2233.4455",
		"001122334455",
		"00:11:22:33:44:55 ",
	} {
		if v, ok := MAC(in); !ok || v != canon {
			t.Errorf("MAC(%q) = (%q, %v), want (%q, true)", in, v, ok, canon)
		}
	}
	if _, ok := MAC("nonsense"); ok {
		t.Error("invalid MAC reported valid")
	}
}

func TestIBAN(t *testing.T) {
	v, ok := IBAN("de89 3704 0044 0532 0130 00")
	check(t, "iban", v, ok, "DE89370400440532013000", true)

	if _, ok := IBAN("DE00 0000 0000 0000 0000 00"); ok {
		t.Error("bad-checksum IBAN reported valid")
	}
}

func TestBankCard(t *testing.T) {
	v, ok := BankCard("4111 1111 1111 1111")
	check(t, "card spaces", v, ok, "4111111111111111", true)

	v, ok = BankCard("4111-1111-1111-1111")
	check(t, "card dashes", v, ok, "4111111111111111", true)

	if _, ok := BankCard("1234 5678 9012 3456"); ok {
		t.Error("non-Luhn card reported valid")
	}
}

func TestE164(t *testing.T) {
	v, ok := E164("+1 (415) 555-2671")
	check(t, "e164 formatting", v, ok, "+14155552671", true)

	// Non-magical: a national number without "+" stays invalid.
	if _, ok := E164("415 555 2671"); ok {
		t.Error("national number reported valid E164")
	}
}

func TestHashes(t *testing.T) {
	const sha256 = "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855"
	v, ok := SHA256("  " + sha256 + "  ")
	check(t, "sha256 lower+trim", v, ok, Lower(sha256), true)

	v, ok = Hex("DEADBEEF")
	check(t, "hex lower", v, ok, "deadbeef", true)

	if _, ok := SHA256("short"); ok {
		t.Error("short sha256 reported valid")
	}
}

func TestURLPathNotReencoded(t *testing.T) {
	// The path must be preserved byte-for-byte (no percent-encoding), and the
	// scheme/host lowercased. This value is invalid (space in path) but the
	// returned best-effort must not alter the path.
	v, _ := URL("HTTP://EXAMPLE.com/a b")
	if v != "http://example.com/a b" {
		t.Fatalf("URL reencoded path: %q", v)
	}
	// Userinfo is preserved; host is lowercased.
	v, _ = URL("https://User:Pass@API.example.com/Path")
	if v != "https://User:Pass@api.example.com/Path" {
		t.Fatalf("URL userinfo/host: %q", v)
	}
}

func TestUUIDUnmatchedBrace(t *testing.T) {
	// A single unmatched brace must not be stripped into a valid UUID.
	if _, ok := UUID("{550e8400-e29b-41d4-a716-446655440000"); ok {
		t.Fatal("unmatched leading brace accepted")
	}
	if _, ok := UUID("550e8400-e29b-41d4-a716-446655440000}"); ok {
		t.Fatal("unmatched trailing brace accepted")
	}
}
