package config

import "strings"

func ServerAddr() string {
	return ":" + configuration().Server.Port
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

func SmtpPort() string {
	return configurationData.Mail.Port
}

func DatabaseUse() string {
	return Configuration().Database.Use
}

func DatabaseMysqlConnectString() string {
	c := Configuration().Database.Mysql
	return c.Username + ":" + c.Password + "@" +
		c.Database + "?" + strings.Join(c.Parameters, "&")
}

func LoggingSeverity() string {
	//return Configuration().Logging.Severity
	return ""
}

func FixedApiToken() string {
	return Configuration().Security.Fixed.Api
}

func IsCorsDisabled() bool {
	return configuration().Security.DisableCors
}
