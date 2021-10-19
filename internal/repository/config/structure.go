package config

type conf struct {
	Server   serverConfig   `yaml:"server"`
	Security securityConfig `yaml:"security"`
}

type serverConfig struct {
	Port string `yaml:"port"`
}

type securityConfig struct {
	DisableCors bool `yaml:"disable_cors"`
}

type validationErrors map[string][]string
