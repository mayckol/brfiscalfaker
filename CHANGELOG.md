# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.0] - 2026-11-24

### Added

- `WithIssuedAt(time.Time)`: sets the issue date and time (`dhEmi`, `dhSaiEnt`/`dSaiEnt`, `dhRecbto`, CF-e `dEmi`/`hEmi`) and the access key's AAMM, written in Brasília time (UTC-03:00). The default issue date is now a random moment in the last 30 days instead of any date since 1900.
- `WithSeed(int64)`: reproducible output.
- `WithDestCNPJ(string)`: recipient CNPJ for NF-e.
- `WithAlphanumericCNPJ()`: random alphanumeric issuer CNPJ (July 2026 format).
- Alphanumeric CNPJ support end to end: `WithCNPJ`/`WithDestCNPJ` accept alphanumeric CNPJs (raw or masked, normalized to 14 upper-case characters), and the access key embeds them with the check digit computed as modulo 11 over ASCII - 48 values. Numeric keys are unchanged.
- `br_documents.AccessKeyDV`, `br_documents.ValidateAccessKey`, `br_documents.CompleteCPF`; `AccessKeyConfig` gains optional `Model`, `UF` and `IssuedAt`.
- `v2.Normalize`, `v2.CompleteCNPJ` and `v2.HasValidFormat` in `pkg/br_documents/v2`.
- CLI flags `--dest-cnpj`, `--alnum-cnpj`, `--issued-at` and `--seed`.

### Changed

- **`WithCNPJ` now sets only the issuer (emit) CNPJ.** Before, it set every CNPJ in the document (recipient, carrier, pickup, delivery, card). Use `WithDestCNPJ` for the NF-e recipient.
- `WithCNPJ` and `WithDestCNPJ` make `Generate` return an error when the value does not have the CNPJ layout (12 alphanumeric characters and 2 digits); check digits are not enforced.
- `WithCPF` in NF-e makes the recipient a `<CPF>` instead of a `<CNPJ>`; masks are stripped.
- `tpEmis`, `tpAmb` and the `protNFe` `tpAmb` are fixed to `1`; `cStat` is `100`.
- Item values and totals are consistent (`vProd = qCom × vUnCom`, totals and payments equal the item value, taxes computed from it) and bounded (at most R$ 2,500.00).
- In the CF-e template, the issuer CNPJ is in `emit` and `ide/CNPJ` holds a separate software-house CNPJ.

### Fixed

- Access key: model digits follow the document (55 NF-e, 65 NFC-e, 59 CF-e); UF, AAMM, series, number, `tpEmis`, `cNF` and `cDV` match the document fields; `protNFe/infProt/chNFe` and the signature reference carry the key instead of a UUID.
- CF-e: `dEmi`, `hEmi`, `vItem`, `nCFe`, `nserieSAT`, `vCFe` and the other empty fields are filled; the key follows the CF-e layout (cUF AAMM CNPJ 59 nserieSAT nCFe cNF DV) and matches `Id="CFe..."`.
- NFC-e and NF-e devolução: `infCpl`, `infRespTec`, `dhSaiEnt`, IPI and other empty fields are filled.
- `CFOP` in NFC-e and CF-e is a valid sale CFOP instead of random letters; empty `cEAN` is `SEM GTIN`.

## [1.2.2] - 2026-06-25

### Fixed

- `detProdCFOP` generates only valid sale CFOPs (no 59xx).

## [1.2.1] - 2026-06-09

### Added

- Shipped the `pkg/br_documents/v2` alphanumeric CNPJ module source, documented in 1.2.0 but missing from that tag.

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
