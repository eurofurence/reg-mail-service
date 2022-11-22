package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"sync"

	"github.com/eurofurence/reg-mail-service/internal/repository/logging/consolelogging/logformat"
	"gopkg.in/yaml.v2"
)

var (
	configurationData *conf
	configurationLock *sync.RWMutex
)

func init() {
	configurationData = &conf{}
	configurationLock = &sync.RWMutex{}
}

func logValidationErrors(errs validationErrors) error {
	if len(errs) != 0 {
		var keys []string
		for key := range errs {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, k := range keys {
			key := k
			val := errs[k]
			for _, errorvalue := range val {
				// cannot use logging package here as this would create a circular dependency (logging needs config)
				log.Print(logformat.Logformat("ERROR", "00000000", fmt.Sprintf("configuration error: %s: %v", key, errorvalue)))
			}
		}
		return errors.New("configuration validation error, see log output for details")
	}

	return nil
}

func configuration() *conf {
	return configurationData
}

func setConfigurationDefaults(c *conf) {
	if c.Server.Port == "" {
		c.Server.Port = "8181"
	}
}

func validateConfiguration(newConfigurationData *conf) error {
	errs := validationErrors{}

	validateServerConfiguration(errs, newConfigurationData.Server)
	//validateLoggingConfiguration(errs, newConfigurationData.Logging)
	validateMailConfiguration(errs, newConfigurationData.Mail)
	validateSecurityConfiguration(errs, newConfigurationData.Security)
	//validateDatabaseConfiguration(errs, newConfigurationData.Database)

	return logValidationErrors(errs)
}

func parseAndOverwriteConfig(yamlFile []byte) error {
	newConfigurationData := &conf{}
	err := yaml.UnmarshalStrict(yamlFile, newConfigurationData)
	if err != nil {
		return err
	}

	setConfigurationDefaults(newConfigurationData)

	err = validateConfiguration(newConfigurationData)
	if err != nil {
		return err
	}

	configurationData = newConfigurationData
	return nil
}

func LoadConfiguration(configurationFilename string) error {
	if configurationFilename == "" {
		return errors.New("no configuration filename provided")
	}

	log.Print(logformat.Logformat("INFO", "00000000", fmt.Sprintf("Reading configuration at %s ...", configurationFilename)))
	yamlFile, err := ioutil.ReadFile(configurationFilename)
	if err != nil {
		return err
	}

	err = parseAndOverwriteConfig(yamlFile)
	return err
}

func Configuration() *conf {
	configurationLock.RLock()
	defer configurationLock.RUnlock()
	return configurationData
}
