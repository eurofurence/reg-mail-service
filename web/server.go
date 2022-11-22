package web

import (
	"github.com/eurofurence/reg-mail-service/web/middleware"
	"net/http"

	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/eurofurence/reg-mail-service/web/controller/healthctl"
	"github.com/eurofurence/reg-mail-service/web/controller/mailctl"
	"github.com/eurofurence/reg-mail-service/web/controller/templatectl"
	"github.com/eurofurence/reg-mail-service/web/filter/corsfilter"
	"github.com/eurofurence/reg-mail-service/web/filter/logreqid"
	"github.com/go-chi/chi"
)

func Create() chi.Router {
	logging.NoCtx().Info("Building routers...")
	server := chi.NewRouter()

	server.Use(middleware.AddRequestIdToContextAndResponse)
	server.Use(logreqid.LogRequestIdMiddleware())
	server.Use(corsfilter.CorsHeadersMiddleware())
	// server.Use(middleware.TokenValidator) //TODO: Fix Token Validation

	healthctl.Create(server)
	mailctl.Create(server)
	templatectl.Create(server)
	return server
}

func Serve(server chi.Router) {
	address := config.ServerAddr()
	logging.NoCtx().Info("Listening on " + address)
	err := http.ListenAndServe(address, server)
	if err != nil {
		logging.NoCtx().Error(err)
	}
}
