package author

type AuthorConfig struct {
	MinNameLength int `json:"min_name_length"` // Minimum length of the author's name
	MaxNameLength int `json:"max_name_length"` // Maximum length of the author's name
}
