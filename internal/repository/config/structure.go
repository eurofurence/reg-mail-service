package config

type (
	DatabaseType string
	LogStyle     string
)

const (
	Inmemory DatabaseType = "inmemory"
	Mysql    DatabaseType = "mysql"

	Plain LogStyle = "plain"
	ECS   LogStyle = "ecs" // default
)

type (
	// Application is the root configuration type
	Application struct {
		Server   ServerConfig   `yaml:"server"`
		Mail     MailConfig     `yaml:"mail"`
		Database DatabaseConfig `yaml:"database"`
		Security SecurityConfig `yaml:"security"`
		Logging  LoggingConfig  `yaml:"logging"`
	}

	// ServerConfig contains all values for http configuration
	ServerConfig struct {
		Address      string `yaml:"address"`
		Port         string `yaml:"port"`
		ReadTimeout  int    `yaml:"read_timeout_seconds"`
		WriteTimeout int    `yaml:"write_timeout_seconds"`
		IdleTimeout  int    `yaml:"idle_timeout_seconds"`
	}

	// DatabaseConfig configures which db to use (mysql, inmemory)
	// and how to connect to it (needed for mysql only)
	DatabaseConfig struct {
		Use        DatabaseType `yaml:"use"`
		Username   string       `yaml:"username"`
		Password   string       `yaml:"password"`
		Database   string       `yaml:"database"`
		Parameters []string     `yaml:"parameters"`
	}

	// SecurityConfig configures everything related to security
	SecurityConfig struct {
		Fixed FixedTokenConfig    `yaml:"fixed_token"`
		Oidc  OpenIdConnectConfig `yaml:"oidc"`
		Cors  CorsConfig          `yaml:"cors"`
	}

	FixedTokenConfig struct {
		Api string `yaml:"api"` // shared-secret for server-to-server backend authentication
	}

	OpenIdConnectConfig struct {
		TokenCookieName    string   `yaml:"token_cookie_name"`     // optional, if set, the jwt token is also read from this cookie (useful for mixed web application setups, see reg-auth-service)
		TokenPublicKeysPEM []string `yaml:"token_public_keys_PEM"` // a list of public RSA keys in PEM format, see https://github.com/Jumpy-Squirrel/jwks2pem for obtaining PEM from openid keyset endpoint
		UserInfoURL        string   `yaml:"user_info_url"`         // validation of admin accesses uses this endpoint to verify the token is still current and access has not been recently revoked
		AdminRole          string   `yaml:"admin_role"`            // the role/group claim that supplies admin rights
	}

	CorsConfig struct {
		DisableCors bool   `yaml:"disable"`
		AllowOrigin string `yaml:"allow_origin"`
	}

	// LoggingConfig configures logging
	LoggingConfig struct {
		Style    LogStyle `yaml:"style"`
		Severity string   `yaml:"severity"`
	}

	// MailConfig contains values for the mail server
	MailConfig struct {
		LogOnly  bool     `yaml:"log_only"`
		DevMode  bool     `yaml:"dev_mode"`
		DevMails []string `yaml:"dev_mails"`
		From     string   `yaml:"from"`
		FromPass string   `yaml:"from_password"`
		Host     string   `yaml:"smtp_host"`
		Port     string   `yaml:"smtp_port"`
	}
)

type validationErrors map[string][]string
