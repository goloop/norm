# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.2] - 2026-10-18

Normalize IP addresses with a single parse.

## [0.4.0] - 2026-07-12

### Added
- `EmailFold` returns `Email` with the whole address lower-cased, for using an
  email as a case-insensitive identity (login key, uniqueness or dedup key), so
  `Boss@Acme.io` and `boss@acme.io` fold to the same value. `Email` still
  preserves the local part's case for exact deliverability.
- `Clean` returns safe, single-line display text in one call: it removes
  invisible and non-space control characters, collapses every whitespace run to
  a single space, and trims the ends - the common cleanup for a free-text form
  field before it is stored or shown.
