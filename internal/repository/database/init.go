package database

import (
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/dbrepo"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/inmemorydb"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/mysqldb"
)

var (
	ActiveRepository dbrepo.Repository
)

func SetRepository(repository dbrepo.Repository) {
	ActiveRepository = repository
}

func Open() error {
	var r dbrepo.Repository
	if config.DatabaseUse() == "mysql" {
		aulogging.Logger.NoCtx().Info().Print("Opening mysql database...")
		//r = historizeddb.Create(mysqldb.Create())
		r = mysqldb.Create()
	} else {
		aulogging.Logger.NoCtx().Info().Print("Opening inmemory database...")
		//r = historizeddb.Create(inmemorydb.Create())
		r = inmemorydb.Create()
	}
	err := r.Open()
	SetRepository(r)
	return err
}

func Close() {
	aulogging.Logger.NoCtx().Info().Print("Closing database...")
	GetRepository().Close()
	SetRepository(nil)
}

func MigrateIfSwitchedOn() (err error) {
	if config.MigrateDatabase() {
		aulogging.Logger.NoCtx().Info().Print("Migrating database...")
		err = GetRepository().Migrate()
	} else {
		aulogging.Logger.NoCtx().Info().Print("Not migrating database. Provide -migrate-database command line switch to enable.")
	}
	return
}

func GetRepository() dbrepo.Repository {
	if ActiveRepository == nil {
		aulogging.Logger.NoCtx().Fatal().Print("You must Open() the database before using it. This is an error in your implementation.")
	}
	return ActiveRepository
}
