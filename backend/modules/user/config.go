package user

type UserConfig struct {
	MinUsernameLength int                   `json:"min_username_length"` // Minimum length of the username
	MaxUsernameLength int                   `json:"max_username_length"` // Maximum length of the username
	NotAllowedSymbols []string              `json:"not_allowed_symbols"` // Symbols not allowed in usernames
	DefaultRole       string                `json:"default_role"`        // Default role assigned to new users
	Roles             map[string]RoleConfig `json:"roles"`               // Map of role names to their configurations
}

type RoleConfig struct {
	Permissions []string `json:"permissions"` // Permissions associated with the role
}
