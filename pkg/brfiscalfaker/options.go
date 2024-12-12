package brfiscalfaker

// Option defines a function type for generator configuration options.
type Option func(*generationConfig)

// generationConfig holds configuration options for the generator.
type generationConfig struct {
	blockedPlaceholders []string
	CPF                 string
	CNPJ                string
}

// WithBlockedPlaceholders returns an Option that blocks the specified placeholders.
func WithBlockedPlaceholders(placeholders ...string) Option {
	return func(cfg *generationConfig) {
		cfg.blockedPlaceholders = append(cfg.blockedPlaceholders, placeholders...)
	}
}

// WithCPF returns an Option to set a custom CPF.
func WithCPF(cpf string) Option {
	return func(cfg *generationConfig) {
		cfg.CPF = cpf
	}
}

// WithCNPJ returns an Option to set a custom CNPJ.
func WithCNPJ(cnpj string) Option {
	return func(cfg *generationConfig) {
		cfg.CNPJ = cnpj
	}
}
