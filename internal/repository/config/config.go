package config

import "strings"

func ServerAddr() string {
	return ":" + configuration().Server.Port
}

func EmailFrom() string {
	return configurationData.Mail.From
}

func DatabaseUse() string {
	return Configuration().Database.Use
}

func DatabaseMysqlConnectString() string {
	c := Configuration().Database.Mysql
	return c.Username + ":" + c.Password + "@" +
		c.Database + "?" + strings.Join(c.Parameters, "&")
}

func IsCorsDisabled() bool {
	return configuration().Security.DisableCors
}
