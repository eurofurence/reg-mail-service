package config

type conf struct {
	Server   serverConfig   `yaml:"server"`
	Mail     mailConfig     `yaml:"mail"`
	Database databaseConfig `yaml:"database"`
	Security securityConfig `yaml:"security"`
	Logging  loggingConfig  `yaml:"logging"`
}

type serverConfig struct {
	Port string `yaml:"port"`
}

type loggingConfig struct {
	Severity string `yaml:"severity"`
}

type fixedTokenConfig struct {
	Api string `yaml:"api"` // shared-secret for server-to-server backend authentication
}

type openIdConnectConfig struct {
	TokenCookieName    string   `yaml:"token_cookie_name"`     // optional, if set, the jwt token is also read from this cookie (useful for mixed web application setups, see reg-auth-service)
	TokenPublicKeysPEM []string `yaml:"token_public_keys_PEM"` // a list of public RSA keys in PEM format, see https://github.com/Jumpy-Squirrel/jwks2pem for obtaining PEM from openid keyset endpoint
	AdminRole          string   `yaml:"admin_role"`            // the role/group claim that supplies admin rights
}

type mailConfig struct {
	From     string `yaml:"from"`
	FromPass string `yaml:"from-password"`
	Host     string `yaml:"smtp-host"`
	Port     string `yaml:"smtp-port"`
}

type databaseConfig struct {
	Use   string      `yaml:"use"`
	Mysql mysqlConfig `yaml:"mysql"`
}

type securityConfig struct {
	Fixed           fixedTokenConfig    `yaml:"fixed_token"`
	Oidc            openIdConnectConfig `yaml:"oidc"`
	DisableCors     bool                `yaml:"disable_cors"`
	CorsAllowOrigin string              `yaml:"cors_allow_origin"`
}

type mysqlConfig struct {
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
	Database   string   `yaml:"database"`
	Parameters []string `yaml:"parameters"`
}

type validationErrors map[string][]string
