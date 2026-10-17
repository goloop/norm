[![Go Reference](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/goloop/norm) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/norm/blob/v1/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# norm

`norm` normalizes user-entered strings. It is the writing companion to
[`goloop/is`](https://github.com/goloop/is): where `is` reads (is this an
email?), `norm` cleans a value toward the canonical form of that type and
validates the result.

```go
value, ok := norm.Email("  Example @ Gmail.com ")
// value == "Example@gmail.com", ok == true
```

Standard library only, plus `goloop/is` for validation.

## Install

```bash
go get github.com/goloop/norm
```

## Typed normalizers

Each mirrors an `is` predicate and returns `(cleaned, valid)`. The cleaned value
is returned even when invalid, so a UI can suggest "did you mean X?".

| Function | Normalization |
|----------|---------------|
| `Email` | trim, lowercase domain (local part preserved), repair internal spaces |
| `EmailFold` | `Email`, then lowercase the whole address (case-insensitive identity key) |
| `URL` | trim, lowercase scheme and host only (no scheme is invented) |
| `UUID` | trim, strip `{}` and `urn:uuid:`, lowercase |
| `MAC` | regroup to lowercase colon pairs (accepts `:`, `-`, `.`, none) |
| `IBAN` | remove spaces, uppercase, strict checksum |
| `BankCard` | remove spaces and dashes, Luhn (and optional kinds) |
| `E164`, `Phone` | remove spaces, parentheses, dashes, dots (no country code invented) |
| `IMSI` | remove spaces and dashes |
| `IP`, `IPv4`, `IPv6` | canonical form (compressed, lowercase IPv6) |
| `Domain`, `Hostname` | trim, lowercase |
| `Base64`, `Base64URL` | remove whitespace (unwrap) |
| `JWT` | trim, scrub |
| `HexColor` | trim, lowercase (no `#` invented) |
| `RGBColor` | lowercase the `rgb` keyword |
| `Bin` | remove whitespace, lowercase `0b` prefix |
| `Decimal` | remove whitespace (ungroup digits) |
| `Float` | trim and scrub only (grammar is not repaired) |
| `Hex`, `MD5`, `SHA1`, `SHA256`, `SHA512` | trim, lowercase |

Every value is first scrubbed of control and invisible characters (a zero-width
space in an email is removed before validation). A value is never reported valid
unless `is` agrees, and `norm` never strips arbitrary punctuation or invents
data.

## Character toolkit

For input that is not a fixed external format, use the character-group tools.
`Class` is a bitmask (combine with `|`):

```go
norm.Trim(s, norm.Spaces|norm.Controls) // remove groups from both ends
norm.Remove(s, norm.Invisible)          // remove a group everywhere
norm.Keep(s, norm.Letters|norm.Digits)  // keep only these groups
norm.CollapseSpace(s)                   // runs of whitespace -> one space
norm.Clean(s)                           // strip invisible/control, collapse space, trim
norm.Lower(s), norm.Upper(s)            // Unicode case
```

`Clean` is the one-call cleanup for a free-text field (a name, title or
comment): it removes invisible and non-space control characters, collapses every
whitespace run to a single space and trims the ends, so pasted text cannot carry
zero-width bytes, hidden line breaks or ragged spacing into storage.

Groups: `Spaces`, `Controls`, `Invisible` (zero-width / format characters, BOM,
soft hyphen, variation selectors), `Punct`, `Symbols`, `Letters`, `Digits`
(decimal), `Numbers` (all Unicode numbers).

For the character-class predicates of `is`, the honest lossy counterparts are
named `*Only` and return a plain string (not `(string, bool)`), so the name
signals filtering rather than canonicalization:

```go
norm.AlphaOnly("a1-b2")   // "ab"
norm.AlnumOnly("a1-b2!")  // "a1b2"
norm.DigitsOnly("+1 415") // "1415"
```

There is deliberately no `norm.Alpha`/`norm.Digit` (a name implying a canonical
form for a lossy filter), and no mirror for grammar or policy predicates
(`Numeric`, `Title`, `Nickname`, `VariableName`, `SelectorName`), where
filtering would change meaning or identity.

## Documentation

- English reference: [DOC.md](DOC.md)
- Ukrainian reference: [DOC.UK.md](DOC.UK.md)

## License

MIT - see [LICENSE](LICENSE).

## Release v1.0.2

Normalize IP addresses with a single parse.
