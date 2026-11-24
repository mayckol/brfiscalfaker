package nfs

import "time"

// Option defines a function type for generator configuration options.
type Option func(*generationConfig)

// generationConfig holds configuration options for the generator.
type generationConfig struct {
	blockedPlaceholders []string
	CPF                 string
	CNPJ                string
	destCNPJ            string
	alphanumericCNPJ    bool
	issuedAt            time.Time
	seed                *int64
}

// WithBlockedPlaceholders returns an Option that blocks the specified placeholders.
func WithBlockedPlaceholders(placeholders ...string) Option {
	return func(cfg *generationConfig) {
		cfg.blockedPlaceholders = append(cfg.blockedPlaceholders, placeholders...)
	}
}

// WithCPF returns an Option to set the recipient's (dest) CPF. A mask is
// accepted and stripped. In NF-e it switches the recipient from CNPJ to CPF,
// unless WithDestCNPJ is also given.
func WithCPF(cpf string) Option {
	return func(cfg *generationConfig) {
		cfg.CPF = cpf
	}
}

// WithCNPJ returns an Option to set the issuer's (emit) CNPJ, which is also
// embedded in the access key. It accepts numeric or alphanumeric (July 2026
// format) CNPJs, raw or masked; the value is normalized to 14 upper-case
// characters, and Generate returns an error when it does not have the CNPJ
// layout (12 alphanumeric characters and 2 digits). Since v1.3.0 it no longer
// sets the other CNPJs of the document (recipient, carrier, pickup/delivery,
// card); use WithDestCNPJ for the NF-e recipient.
func WithCNPJ(cnpj string) Option {
	return func(cfg *generationConfig) {
		cfg.CNPJ = cnpj
	}
}

// WithDestCNPJ returns an Option to set the recipient's (dest) CNPJ in NF-e.
// Numeric and alphanumeric CNPJs are accepted, raw or masked. It takes
// precedence over WithCPF for the NF-e recipient.
func WithDestCNPJ(cnpj string) Option {
	return func(cfg *generationConfig) {
		cfg.destCNPJ = cnpj
	}
}

// WithAlphanumericCNPJ returns an Option that gives the issuer a random, valid
// alphanumeric CNPJ (July 2026 format). WithCNPJ takes precedence over it.
func WithAlphanumericCNPJ() Option {
	return func(cfg *generationConfig) {
		cfg.alphanumericCNPJ = true
	}
}

// WithIssuedAt returns an Option that sets the issue date and time. It drives
// dhEmi, dhSaiEnt/dSaiEnt, dhRecbto (a few seconds later), the CF-e dEmi/hEmi
// and the AAMM part of the access key. Dates are written in Brasília time
// (UTC-03:00), so the instant is kept whatever the location of t.
// Without it, the issue date is a random moment in the last 30 days.
func WithIssuedAt(t time.Time) Option {
	return func(cfg *generationConfig) {
		cfg.issuedAt = t
	}
}

// WithSeed returns an Option that makes the output reproducible: the same seed,
// template and options produce the same XML. Combine it with WithIssuedAt,
// because the default issue date is relative to the current time. A seed of 0
// means a random seed.
func WithSeed(seed int64) Option {
	return func(cfg *generationConfig) {
		cfg.seed = &seed
	}
}
