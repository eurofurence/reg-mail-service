package config

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/url"
	"regexp"
	"strconv"
)

func setConfigurationDefaults(c *conf) {
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
	if c.Security.CorsAllowOrigin == "" {
		c.Security.CorsAllowOrigin = "*"
	}
}

const portPattern = "^[1-9][0-9]{0,4}$"

func addError(errs validationErrors, key string, value interface{}, message string) {
	errs[key] = append(errs[key], fmt.Sprintf("value '%v' %s", value, message))
}

func validateServerConfiguration(errs validationErrors, sc serverConfig) {
	if sc.Port == "" {
		addError(errs, "server.port", sc.Port, "cannot be empty")
	} else {
		port, err := strconv.ParseUint(sc.Port, 10, 16)
		if err != nil {
			addError(errs, "server.port", sc.Port, "is not a valid port number")
		} else if port <= 1024 {
			addError(errs, "server.port", sc.Port, "must be a nonprivileged port")
		}
	}
}

func validateMailConfiguration(errs validationErrors, m mailConfig) {
	re := regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")

	if m.From == "" {
		addError(errs, "mail.from", m.From, "cannot be empty")
	} else {
		from := re.Find([]byte(m.From))
		if from == nil {
			addError(errs, "mail.from", m.From, "is not a valid email address")
		}
	}

	if m.Host == "" {
		addError(errs, "mail.smtp-host", m.Host, "cannot be empty")
	}

	if m.Port == "" {
		addError(errs, "mail.smtp-port", m.Port, "cannot be empty")
	}
}

func validateDatabaseConfiguration(errs url.Values, c databaseConfig) {

}

func validateSecurityConfiguration(errs validationErrors, sc securityConfig) {
	parsedKeySet = make([]*rsa.PublicKey, 0)
	for i, keyStr := range sc.Oidc.TokenPublicKeysPEM {
		publicKeyPtr, err := jwt.ParseRSAPublicKeyFromPEM([]byte(keyStr))
		if err != nil {
			addError(errs, fmt.Sprintf("security.oidc.token_public_keys_PEM[%d]", i), "(redacted)", fmt.Sprintf("failed to parse RSA public key in PEM format: %s", err.Error()))
		} else {
			parsedKeySet = append(parsedKeySet, publicKeyPtr)
		}
	}
}
