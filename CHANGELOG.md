# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.0] - 2026-04-16

### Added

- **New v2 Alphanumeric CNPJ Module** (`pkg/br_documents/v2`)
  - Full support for Brazil's new alphanumeric CNPJ format (effective July 2026)
  - Módulo 11 algorithm with dual check digits (DV1, DV2)
  - ASCII - 48 character-to-numeric conversion
  - 14-character format: 8 alphanumeric root + 4 alphanumeric branch + 2 numeric check digits
  - Automatic check digit calculation and validation
  - Blacklist support for ambiguous letters (I, O, Q, U, F) with `AllowAmbiguousLetters` flag
  - Masked format support (XX.XXX.XXX/YYYY-ZZ)
  - Full backward compatibility with numeric-only CNPJs
  - Input sanitization for mixed case, masks, and whitespace
  - Comprehensive test coverage (98.5%)
  - 20 unit tests covering all major functions and edge cases

### Documentation

- Added "Alphanumeric CNPJ (v2) — July 2026 Format" section to README.md
- Documented usage examples for raw and masked CNPJ generation
- Configuration parameters documentation
- Validation usage examples

## [1.1.2] - 2026-04-09

### Fixed

- Special character handling in template processing

## [1.1.1] - 2026-04-08

### Fixed

- Removed unnecessary package dependencies

## [1.1.0] - 2026-04-07

### Fixed

- Go version compatibility

## [1.0.10] - 2026-03-15

### Fixed

- Various bug fixes and improvements

## [1.0.9] - 2026-03-01

### Initial Release

- Support for multiple invoice types (NF-e, NFC-e, CFe, NFeDevolucao)
- Customizable CPF and CNPJ numbers
- XML tag blocking capability
- Cross-platform support
- Brazilian fiscal identifier generation (numeric CNPJ, CPF, Access Key)
