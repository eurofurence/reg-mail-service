package config

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
)

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
}

func validateDatabaseConfiguration(errs url.Values, c databaseConfig) {

}

func validateSecurityConfiguration(errs validationErrors, sc securityConfig) {
}
