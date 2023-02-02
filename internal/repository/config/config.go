package config

import (
	"crypto/rsa"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func UseEcsLogging() bool {
	return ecsLogging
}

func ServerAddr() string {
	c := Configuration()
	return fmt.Sprintf("%s:%s", c.Server.Address, c.Server.Port)
}

func ServerReadTimeout() time.Duration {
	return time.Second * time.Duration(Configuration().Server.ReadTimeout)
}

func ServerWriteTimeout() time.Duration {
	return time.Second * time.Duration(Configuration().Server.WriteTimeout)
}

func ServerIdleTimeout() time.Duration {
	return time.Second * time.Duration(Configuration().Server.IdleTimeout)
}

func MailLogOnly() bool {
	return configurationData.Mail.LogOnly
}

func MailDevMode() bool {
	return configurationData.Mail.DevMode
}

func MailDevMails() []string {
	return configurationData.Mail.DevMails
}

func EmailFrom() string {
	return configurationData.Mail.From
}

func EmailFromPassword() string {
	return configurationData.Mail.FromPass
}

func SmtpHost() string {
	return configurationData.Mail.Host
}

func SmtpPort() int {
	port, _ := strconv.Atoi(configurationData.Mail.Port)
	return port
}

func DatabaseUse() DatabaseType {
	return Configuration().Database.Use
}

func DatabaseMysqlConnectString() string {
	c := Configuration().Database
	return c.Username + ":" + c.Password + "@" +
		c.Database + "?" + strings.Join(c.Parameters, "&")
}

func MigrateDatabase() bool {
	return dbMigrate
}

func LoggingSeverity() string {
	return Configuration().Logging.Severity
}

func FixedApiToken() string {
	return Configuration().Security.Fixed.Api
}

func OidcIdTokenCookieName() string {
	return Configuration().Security.Oidc.IdTokenCookieName
}

func OidcAccessTokenCookieName() string {
	return Configuration().Security.Oidc.AccessTokenCookieName
}

func OidcKeySet() []*rsa.PublicKey {
	return parsedKeySet
}

func OidcAdminGroup() string {
	return Configuration().Security.Oidc.AdminGroup
}

func OidcAllowedAudience() string {
	return Configuration().Security.Oidc.Audience
}

func OidcAllowedIssuer() string {
	return Configuration().Security.Oidc.Issuer
}

func AuthServiceBaseUrl() string {
	return Configuration().Security.Oidc.AuthService
}

func IsCorsDisabled() bool {
	return Configuration().Security.Cors.DisableCors
}

func CorsAllowOrigin() string {
	return Configuration().Security.Cors.AllowOrigin
}
