package bookmarks

type BookmarksConfig struct {
	MinNameLength     int      `json:"min_name_length"`     // Minimum length of the bookmark's name
	MaxNameLength     int      `json:"max_name_length"`     // Maximum length of the bookmark's name
	NotAllowedSymbols []string `json:"not_allowed_symbols"` // Symbols not allowed in bookmark names
}
