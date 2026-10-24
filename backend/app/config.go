package app

import (
	"remedia/modules/author"
	"remedia/modules/bookmarks"
	"remedia/modules/comments"
	"remedia/modules/news"
	"remedia/modules/payment"
	"remedia/modules/subscription"
	"remedia/modules/title"
	"remedia/modules/user"
)

type Config struct {
	Server       ServerConfig                    `json:"server"`
	Author       author.AuthorConfig             `json:"author"`
	Bookmarks    bookmarks.BookmarksConfig       `json:"bookmarks"`
	Comments     comments.CommentsConfig         `json:"comments"`
	News         news.NewsConfig                 `json:"news"`
	Payment      payment.PaymentConfig           `json:"payment"`
	Subscription subscription.SubscriptionConfig `json:"subscription"`
	Title        title.TitleConfig               `json:"title"`
	User         user.UserConfig                 `json:"user"`
}

type ServerConfig struct {
	Port               int    `json:"port"`                  // Port on which the server will listen
	MaxSizeRequestBody int    `json:"max_size_request_body"` // Maximum size of the request body in bytes
	RevokeKeyTTL       int    `json:"revoke_key_ttl"`        // Time-to-live for revoked keys in seconds
	LoggingLevel       string `json:"logging_level"`         // Logging level (e.g., "debug", "info", "warn", "error")
	LoggingToFilePath  string `json:"logging_to_file_path"`  // Path to the log file
	AllowCORS          bool   `json:"allow_cors"`            // Whether to allow Cross-Origin Resource Sharing (CORS)
	RateLimitEnabled   bool   `json:"rate_limit_enabled"`    // Whether rate limiting is enabled
	SecretKey          string `json:"-"`                     // Path to the secret key file
}
