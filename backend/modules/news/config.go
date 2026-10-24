package news

type NewsConfig struct {
	Permissions    NewsPermissions `json:"permissions"`      // Permissions for news posts
	MinTitleLength int             `json:"min_title_length"` // Minimum length of the news title
	MaxTitleLength int             `json:"max_title_length"` // Maximum length of the news title
	MinTextLength  int             `json:"min_text_length"`  // Minimum length of the news text
	MaxTextLength  int             `json:"max_text_length"`  // Maximum length of the news text
	AllowImages    bool            `json:"allow_images"`     // Whether images are allowed in news posts
	AllowLinks     bool            `json:"allow_links"`      // Whether links are allowed in news posts
}

type NewsPermissions struct {
	Read   []string `json:"read"`   // Permissions to read news posts
	Write  []string `json:"write"`  // Permissions to write news posts
	Edit   []string `json:"edit"`   // Permissions to edit news posts
	Delete []string `json:"delete"` // Permissions to delete news posts
}
