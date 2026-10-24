package title

type TitleConfig struct {
	MinLength         int      `json:"min_length"`          // Minimum length of the title
	MaxLength         int      `json:"max_length"`          // Maximum length of the title
	NotAllowedSymbols []string `json:"not_allowed_symbols"` // Symbols not allowed in titles
	CacheSize         int      `json:"cache_size"`          // Size of the title cache
}
