package comments

type CommentsConfig struct {
	Permissions       CommentPermissions `json:"permissions"`         // Permissions for comments
	MaxCommentLength  int                `json:"max_comment_length"`  // Maximum length of a comment
	AllowImages       bool               `json:"allow_images"`        // Whether images are allowed in comments
	AllowLinks        bool               `json:"allow_links"`         // Whether links are allowed in comments
	NotAllowedSymbols []string           `json:"not_allowed_symbols"` // Symbols not allowed in comments
}

type CommentPermissions struct {
	Read   []string `json:"read"`   // Permissions to read comments
	Write  []string `json:"write"`  // Permissions to write comments
	Edit   []string `json:"edit"`   // Permissions to edit comments
	Delete []string `json:"delete"` // Permissions to delete comments
}
