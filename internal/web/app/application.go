package app

import (
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/database"
)

type Application interface {
	Run() int
}

type Impl struct{}

func New() Application {
	return &Impl{}
}

func (i *Impl) Run() int {
	config.ParseCommandLineFlags()
	setupLogging("mail-service", config.UseEcsLogging())

	if err := config.StartupLoadConfiguration(); err != nil {
		return 1
	}
	setLoglevel(config.LoggingSeverity())

	if err := database.Open(); err != nil {
		return 1
	}
	defer database.Close()
	if err := database.MigrateIfSwitchedOn(); err != nil {
		return 1
	}

	if err := runServerWithGracefulShutdown(); err != nil {
		return 2
	}

	return 0
}
