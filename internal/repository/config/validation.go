package config

import (
	"crypto/rsa"
	"fmt"
	"github.com/eurofurence/reg-mail-service/internal/web/util/validation"
	"github.com/golang-jwt/jwt/v4"
	"net/url"
	"regexp"
)

func setConfigurationDefaults(c *Application) {
	if c.Server.Port == "" {
		c.Server.Port = "8080"
	}
	if c.Server.ReadTimeout <= 0 {
		c.Server.ReadTimeout = 5
	}
	if c.Server.WriteTimeout <= 0 {
		c.Server.WriteTimeout = 5
	}
	if c.Server.IdleTimeout <= 0 {
		c.Server.IdleTimeout = 5
	}
	if c.Logging.Severity == "" {
		c.Logging.Severity = "INFO"
	}
	if c.Database.Use == "" {
		c.Database.Use = "inmemory"
	}
	if c.Security.Cors.AllowOrigin == "" {
		c.Security.Cors.AllowOrigin = "*"
	}
}

const portPattern = "^[1-9][0-9]{0,4}$"

func validateServerConfiguration(errs url.Values, c ServerConfig) {
	if validation.ViolatesPattern(portPattern, c.Port) {
		errs.Add("server.port", "must be a number between 1 and 65535")
	}
	validation.CheckIntValueRange(&errs, 1, 300, "server.read_timeout_seconds", c.ReadTimeout)
	validation.CheckIntValueRange(&errs, 1, 300, "server.write_timeout_seconds", c.WriteTimeout)
	validation.CheckIntValueRange(&errs, 1, 300, "server.idle_timeout_seconds", c.IdleTimeout)
}

var allowedSeverities = []string{"DEBUG", "INFO", "WARN", "ERROR"}

func validateLoggingConfiguration(errs url.Values, c LoggingConfig) {
	if validation.NotInAllowedValues(allowedSeverities[:], c.Severity) {
		errs.Add("logging.severity", "must be one of DEBUG, INFO, WARN, ERROR")
	}
}

const mailRegexp = `^[^@\s]+@[^@\s]+$`

func validateMailConfiguration(errs url.Values, m MailConfig) {
	re := regexp.MustCompile(mailRegexp)

	if m.From == "" {
		errs.Add("mail.from", m.From+" cannot be empty")
	} else {
		from := re.Find([]byte(m.From))
		if from == nil {
			errs.Add("mail.from", m.From+" is not a valid email address")
		}
	}

	if m.Host == "" {
		errs.Add("mail.smtp_host", m.Host+" cannot be empty")
	}

	if m.Port == "" {
		errs.Add("mail.smtp_port", m.Port+" cannot be empty")
	}
}

var allowedDatabases = []DatabaseType{Mysql, Inmemory}

func validateDatabaseConfiguration(errs url.Values, c DatabaseConfig) {
	if validation.NotInAllowedValues(allowedDatabases[:], c.Use) {
		errs.Add("database.use", "must be one of mysql, inmemory")
	}
	if c.Use == Mysql {
		validation.CheckLength(&errs, 1, 256, "database.username", c.Username)
		validation.CheckLength(&errs, 1, 256, "database.password", c.Password)
		validation.CheckLength(&errs, 1, 256, "database.database", c.Database)
	}
}

func validateSecurityConfiguration(errs url.Values, c SecurityConfig) {
	validation.CheckLength(&errs, 16, 256, "security.fixed.api", c.Fixed.Api)
	validation.CheckLength(&errs, 1, 256, "security.oidc.admin_group", c.Oidc.AdminGroup)

	parsedKeySet = make([]*rsa.PublicKey, 0)
	for i, keyStr := range c.Oidc.TokenPublicKeysPEM {
		publicKeyPtr, err := jwt.ParseRSAPublicKeyFromPEM([]byte(keyStr))
		if err != nil {
			errs.Add(fmt.Sprintf("security.oidc.token_public_keys_PEM[%d]", i), fmt.Sprintf("failed to parse RSA public key in PEM format: %s", err.Error()))
		} else {
			parsedKeySet = append(parsedKeySet, publicKeyPtr)
		}
	}
}
