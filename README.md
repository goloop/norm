[![Go Reference](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/goloop/norm) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/norm/blob/master/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

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
| `URL` | trim, lowercase scheme and host only (no scheme is invented) |
| `UUID` | trim, strip `{}` and `urn:uuid:`, lowercase |
| `MAC` | regroup to lowercase colon pairs (accepts `:`, `-`, `.`, none) |
| `IBAN` | remove spaces, uppercase, strict checksum |
| `BankCard` | remove spaces and dashes, Luhn (and optional kinds) |
| `E164` | remove spaces, dashes, parentheses, dots (no country code invented) |
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
norm.Lower(s), norm.Upper(s)            // Unicode case
```

Groups: `Spaces`, `Controls`, `Invisible` (zero-width / format characters, BOM,
soft hyphen, variation selectors), `Punct`, `Symbols`, `Letters`, `Digits`.

`Keep` and `Remove` are the honest replacement for lossy property mirrors: there
is deliberately no `norm.Alpha` or `norm.Digit`, because "normalizing" to those
is really filtering, so you write `norm.Keep(s, norm.Letters)` instead.

## Documentation

- English reference: [DOC.md](DOC.md)
- Ukrainian reference: [DOC.UK.md](DOC.UK.md)

## License

MIT - see [LICENSE](LICENSE).
