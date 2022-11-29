package web

import (
	"github.com/StephanHCB/go-autumn-logging-zerolog/loggermiddleware"
	"github.com/eurofurence/reg-mail-service/web/middleware"
	"net/http"

	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/eurofurence/reg-mail-service/web/controller/healthctl"
	"github.com/eurofurence/reg-mail-service/web/controller/mailctl"
	"github.com/eurofurence/reg-mail-service/web/controller/templatectl"
	"github.com/go-chi/chi/v5"
)

func Create() chi.Router {
	logging.NoCtx().Info("Building routers...")
	server := chi.NewRouter()

	server.Use(middleware.AddRequestIdToContextAndResponse)
	server.Use(loggermiddleware.AddZerologLoggerToContext)
	server.Use(middleware.RequestLogger)
	server.Use(middleware.PanicRecoverer)
	server.Use(middleware.CorsHandling)
	server.Use(middleware.TokenValidator)

	healthctl.Create(server)
	mailctl.Create(server)
	templatectl.Create(server)
	return server
}

func Serve(server chi.Router) {
	setupLogging("attendee-service", config.UseEcsLogging())

	address := config.ServerAddr()
	logging.NoCtx().Info("Listening on " + address)
	err := http.ListenAndServe(address, server)
	if err != nil {
		logging.NoCtx().Error(err)
	}
}
