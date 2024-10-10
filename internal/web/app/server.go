package app

import (
	"context"
	"errors"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/StephanHCB/go-autumn-logging-zerolog/loggermiddleware"
	"github.com/eurofurence/reg-mail-service/internal/web/controller/healthctl"
	"github.com/eurofurence/reg-mail-service/internal/web/controller/mailctl"
	mailctlv2 "github.com/eurofurence/reg-mail-service/internal/web/controller/mailctl/v2"
	"github.com/eurofurence/reg-mail-service/internal/web/controller/templatectl"
	templatectlv2 "github.com/eurofurence/reg-mail-service/internal/web/controller/templatectl/v2"
	middleware2 "github.com/eurofurence/reg-mail-service/internal/web/middleware"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/go-chi/chi/v5"
)

func CreateRouter(ctx context.Context) chi.Router {
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
	mailctlv2.Create(server)
	templatectl.Create(server)
	templatectlv2.Create(server)
	return server
}

func newServer(ctx context.Context, router chi.Router) *http.Server {
	aulogging.Logger.NoCtx().Debug().Print("setting up server")
	return &http.Server{
		Addr:         config.ServerAddr(),
		Handler:      router,
		ReadTimeout:  config.ServerReadTimeout(),
		WriteTimeout: config.ServerWriteTimeout(),
		IdleTimeout:  config.ServerIdleTimeout(),
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}
}

func runServerWithGracefulShutdown() error {
	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	handler := CreateRouter(ctx)
	srv := newServer(ctx, handler)

	go func() {
		<-sig
		defer cancel()
		aulogging.Logger.NoCtx().Debug().Print("Stopping services now")

		tCtx, tcancel := context.WithTimeout(ctx, time.Second*5)
		defer tcancel()

		if err := srv.Shutdown(tCtx); err != nil {
			aulogging.Logger.NoCtx().Error().WithErr(err).Printf("Couldn't shutdown server gracefully: %s", err.Error())
			os.Exit(3)
		}
	}()

	aulogging.Logger.NoCtx().Info().Print("Running service on ", config.ServerAddr())
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		aulogging.Logger.NoCtx().Error().WithErr(err).Printf("Server closed unexpectedly: %s", err.Error())
		return err
	}

	return nil
}
