package app

import (
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/StephanHCB/go-autumn-logging-zerolog/loggermiddleware"
	"github.com/eurofurence/reg-mail-service/internal/web/controller/healthctl"
	"github.com/eurofurence/reg-mail-service/internal/web/controller/mailctl"
	"github.com/eurofurence/reg-mail-service/internal/web/controller/templatectl"
	middleware2 "github.com/eurofurence/reg-mail-service/internal/web/middleware"
	"net/http"

	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/go-chi/chi/v5"
)

func Create() chi.Router {
	aulogging.Logger.NoCtx().Debug().Print("Setting up router")
	server := chi.NewRouter()

	server.Use(middleware2.AddRequestIdToContextAndResponse)
	server.Use(loggermiddleware.AddZerologLoggerToContext)
	server.Use(middleware2.RequestLogger)
	server.Use(middleware2.PanicRecoverer)
	server.Use(middleware2.CorsHandling)
	server.Use(middleware2.TokenValidator)

	healthctl.Create(server)
	mailctl.Create(server)
	templatectl.Create(server)
	return server
}

func Serve(server chi.Router) {
	setupLogging("mail-service", config.UseEcsLogging())
	setLoglevel(config.LoggingSeverity())

	address := config.ServerAddr()
	logging.NoCtx().Info("Listening on " + address)
	err := http.ListenAndServe(address, server)
	if err != nil {
		logging.NoCtx().Error(err)
	}
}
