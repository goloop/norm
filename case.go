package norm

import "strings"

// Lower returns s with all Unicode letters mapped to lower case. It is the
// normalizing counterpart of is.Lower and always succeeds, so it returns only a
// string.
func Lower(s string) string { return strings.ToLower(s) }

// Upper returns s with all Unicode letters mapped to upper case. It is the
// normalizing counterpart of is.Upper and always succeeds, so it returns only a
// string.
func Upper(s string) string { return strings.ToUpper(s) }
