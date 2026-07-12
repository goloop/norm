# norm - reference

`norm` normalizes strings toward the canonical form of a known type and
validates the result with `goloop/is`. Full English reference; Ukrainian in
[DOC.UK.md](DOC.UK.md).

## Contents

- [Model](#model)
- [Typed normalizers](#typed-normalizers)
- [Per-type rules](#per-type-rules)
- [Character toolkit](#character-toolkit)
- [Symmetry with is](#symmetry-with-is)
- [Scope](#scope)

## Model

Each typed normalizer returns `(value string, ok bool)`: the best-effort cleaned
value and whether it validates. The value is returned even when `ok` is false,
so callers can offer a suggestion.

The pipeline for every typed normalizer is:

1. **Scrub** control and invisible characters (never valid inside these
   formats) and trim surrounding whitespace.
2. Apply **safe canonical** steps (for example lowercasing a case-insensitive
   domain or scheme). If the result validates, return it.
3. Apply a bounded **repair** (for example removing internal whitespace) only
   when still invalid, and validate again.
4. Return the cleaned value with the `is` verdict.

Because the result is always confirmed with an `is` predicate, a value is never
reported valid unless `is` agrees. A valid input, including a quoted email local
part that legitimately contains spaces, is preserved because the repair runs
only when the milder result is invalid.

## Typed normalizers

```go
func Email(s string) (string, bool)
func URL(s string) (string, bool)
func UUID(s string) (string, bool)
func MAC(s string) (string, bool)
func IBAN(s string) (string, bool)
func BankCard(s string, kinds ...is.CardKind) (string, bool)
func E164(s string) (string, bool)
func Hex(s string) (string, bool)
func MD5(s string) (string, bool)
func SHA1(s string) (string, bool)
func SHA256(s string) (string, bool)
func SHA512(s string) (string, bool)
```

## Per-type rules

- **Email**: trim; lowercase the domain (after the last `@`) while preserving
  the local part (domains are case-insensitive, local parts technically are
  not); if still invalid, remove internal whitespace. No IDN/punycode
  conversion is performed.
- **URL**: trim; lowercase the scheme and host only. The path, query and
  fragment are case-sensitive and untouched. A missing scheme is never invented,
  so `example.com` stays invalid.
- **UUID**: trim; strip a surrounding `{...}` and a leading `urn:uuid:`;
  lowercase.
- **MAC**: regroup to lowercase colon pairs. Accepts `:`, `-`, `.` and the
  separatorless form; handles EUI-48 (12 hex) and EUI-64 (16 hex).
- **IBAN**: remove spaces, uppercase, validate strictly (including checksum).
- **BankCard**: remove spaces and dashes, validate with Luhn and any given
  `is.CardKind`. **Security:** the returned value is a full card number (PAN);
  do not log or store it unnecessarily, and a Luhn-valid number is not proof of
  a real or intended card.
- **E164**: remove spaces, dashes, parentheses and dots. Non-magical: a national
  number without a leading `+` stays invalid.
- **Hex / MD5 / SHA1 / SHA256 / SHA512**: trim and lowercase (the canonical form
  of a hex digest).

## Character toolkit

```go
type Class uint
const (
	Spaces Class = 1 << iota // Unicode White_Space (incl. NBSP)
	Controls                 // Unicode Cc
	Invisible                // Cf + variation selectors + CGJ (zero-width/format)
	Punct                    // Unicode P
	Symbols                  // Unicode S
	Letters                  // Unicode L
	Digits                   // Unicode Nd
)

func Trim(s string, class Class) string      // remove class from both ends
func TrimLeft(s string, class Class) string
func TrimRight(s string, class Class) string
func Remove(s string, class Class) string    // remove class everywhere
func Keep(s string, class Class) string      // keep only class (inverse of Remove)
func CollapseSpace(s string) string          // internal whitespace runs -> one space
func Lower(s string) string                  // Unicode lower case
func Upper(s string) string                  // Unicode upper case
```

Combine groups with `|`, for example `norm.Remove(s, norm.Controls|norm.Invisible)`.
A zero `Class` is a no-op for `Trim` and `Remove`.

## Symmetry with is

`norm` mirrors the **format** predicates of `is` (Email, URL, UUID, ...), where
"normalize" has a single, safe meaning: produce the canonical form and validate.

It deliberately does **not** mirror the **property** predicates (`is.Alpha`,
`is.Digit`, `is.Space`, ...). A `norm.Alpha` would suggest a canonical form but
would actually remove non-letters, which is lossy filtering. Use the toolkit
instead: `norm.Keep(s, norm.Letters)` for the "letters only" case, or
`norm.Remove` / `norm.Trim` for the rest. The case predicates map to the plain
transforms `norm.Lower` and `norm.Upper`.

## Scope

`norm` cleans and validates known types and provides character-group tools. It
does not parse structured documents, transliterate (see `goloop/t13n`), slugify
(see `goloop/slug`) or change letter case beyond `Lower`/`Upper`.
