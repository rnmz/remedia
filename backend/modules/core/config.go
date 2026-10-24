package core

type CoreConfig struct {
	Valkey ValkeyConfig `json:"valkey"` // Configuration for the Valkey service

}

type ValkeyConfig struct {
	Host     string
	Port     string
	Password string // Optional; leave empty if auth is disabled
}

type DatabaseConfig struct {
	Host     string `json:"host"`     // Hostname or IP address of the database server
	Port     int    `json:"port"`     // Port on which the database server is listening
	Username string `json:"username"` // Username for database authentication
	Password string `json:"password"` // Password for database authentication
}

const (
	S3  = "s3"
	Ark = "ark"
)

type FilesConfig struct {
	Type string `json:"type"` // Type of file storage (e.g., "s3", "ark")

}

type S3Config struct {
	Endpoint     string `json:"endpoint"`       // URL of the S3-compatible storage service
	Region       string `json:"region"`         // S3 region
	Bucket       string `json:"bucket"`         // S3 bucket
	UsePathStyle bool   `json:"use_path_style"` // MinIO and many compatible: true
	UseSSL       bool   `json:"use_ssl"`        // False only for local MinIO

	PublicBaseURL string `json:"public_base_url"` // CDN or public domain, from which files are served
	KeyPrefix     string `json:"key_prefix"`      // FOr example "titles/", to avoid conflicts with everything in the root bucket

	PresignTTLSeconds int `json:"presign_ttl_seconds"` // If private links and signed
	MaxUploadSize     int `json:"max_upload_size"`     // Maximum file size allowed for upload

	AccessKeyID     string `json:"-"` // Acccess Key ID
	SecretAccessKey string `json:"-"` // Secret Access Key
}

type ArkConfig struct {
	Url           string `json:"url"`            // URL of the Ark storage service
	Region        string `json:"region"`         // Region for Ark storage
	SecretKey     string `json:"-"`              // Path to the secret key file for Ark storage
	CompressFiles bool   `json:"compress_files"` // Should files be compressed in storage
}
