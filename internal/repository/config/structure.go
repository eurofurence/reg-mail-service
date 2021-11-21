package config

type conf struct {
	Server   serverConfig   `yaml:"server"`
	Mail     mailConfig     `yaml:"mail"`
	Database databaseConfig `yaml:"database"`
	Security securityConfig `yaml:"security"`
}

type serverConfig struct {
	Port string `yaml:"port"`
}

type mailConfig struct {
	From string `yaml:"from"`
}

type databaseConfig struct {
	Use   string      `yaml:"use"`
	Mysql mysqlConfig `yaml:"mysql"`
}

type securityConfig struct {
	DisableCors bool `yaml:"disable_cors"`
}

type mysqlConfig struct {
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
	Database   string   `yaml:"database"`
	Parameters []string `yaml:"parameters"`
}

type validationErrors map[string][]string
