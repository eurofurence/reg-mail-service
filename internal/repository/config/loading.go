package config

import (
	"crypto/rsa"
	"errors"
	"flag"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/eurofurence/reg-mail-service/internal/repository/logging/consolelogging/logformat"
	"gopkg.in/yaml.v2"
)

var (
	configurationData     *conf
	configurationLock     *sync.RWMutex
	configurationFilename string
	dbMigrate             bool
	ecsLogging            bool

	parsedKeySet []*rsa.PublicKey
)

var (
	ErrorConfigArgumentMissing = errors.New("configuration file argument missing. Please specify using -config argument. Aborting")
	ErrorConfigFile            = errors.New("failed to read or parse configuration file. Aborting")
)

func init() {
	configurationData = &conf{}
	configurationLock = &sync.RWMutex{}

	flag.StringVar(&configurationFilename, "config", "config.yaml", "config file path")
	flag.BoolVar(&dbMigrate, "migrate-database", false, "migrate database on startup")
	flag.BoolVar(&ecsLogging, "ecs-json-logging", false, "switch to structured json logging")
}

// ParseCommandLineFlags is exposed separately so you can skip it for tests
func ParseCommandLineFlags() {
	flag.Parse()
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

func loadConfiguration() error {
	yamlFile, err := os.ReadFile(configurationFilename)
	if err != nil {
		// cannot use logging package here as this would create a circular dependency (logging needs config)
		aulogging.Logger.NoCtx().Error().Printf("failed to load configuration file '%s': %v", configurationFilename, err)
		return err
	}
	err = parseAndOverwriteConfig(yamlFile)
	return err
}

func StartupLoadConfiguration() error {
	aulogging.Logger.NoCtx().Info().Print("Reading configuration...")
	if configurationFilename == "" {
		// cannot use logging package here as this would create a circular dependency (logging needs config)
		aulogging.Logger.NoCtx().Error().Print("Configuration file argument missing. Please specify using -config argument. Aborting.")
		return ErrorConfigArgumentMissing
	}
	err := loadConfiguration()
	if err != nil {
		// cannot use logging package here as this would create a circular dependency (logging needs config)
		aulogging.Logger.NoCtx().Error().Print("Error reading or parsing configuration file. Aborting.")
		return ErrorConfigFile
	}
	return nil
}

func Configuration() *conf {
	configurationLock.RLock()
	defer configurationLock.RUnlock()
	return configurationData
}
