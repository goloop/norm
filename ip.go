package norm

import (
	"net/netip"

	"github.com/goloop/is/v2"
)

// IP normalizes an IP address to its canonical form (lowercase, compressed
// IPv6) and reports whether the result is valid.
func IP(s string) (string, bool) {
	v := canonicalIP(scrub(s))
	return v, is.IP(v)
}

// IPv4 normalizes an IPv4 address and reports whether the result is valid.
func IPv4(s string) (string, bool) {
	s = scrub(s)
	if a, err := netip.ParseAddr(s); err == nil && a.Is4() {
		v := a.String()
		return v, is.IPv4(v)
	}
	return s, is.IPv4(s)
}

// IPv6 normalizes an IPv6 address to its canonical, compressed lowercase form
// and reports whether the result is valid.
func IPv6(s string) (string, bool) {
	s = scrub(s)
	if a, err := netip.ParseAddr(s); err == nil && a.Is6() && a.Zone() == "" {
		v := a.String()
		return v, is.IPv6(v)
	}
	return s, is.IPv6(s)
}

// canonicalIP returns the canonical text of s when it parses as a zoneless IP
// address, otherwise s unchanged. A zoned address is left as-is because the is
// validators reject zones, so canonicalizing it would only be confusing.
func canonicalIP(s string) string {
	if a, err := netip.ParseAddr(s); err == nil && a.Zone() == "" {
		return a.String()
	}
	return s
}
