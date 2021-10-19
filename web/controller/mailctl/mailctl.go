package mailctl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eurofurence/reg-mail-service/api/v1/health"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/eurofurence/reg-mail-service/web/util/media"
	"github.com/go-chi/chi"
	"github.com/go-http-utils/headers"
)

func Create(server chi.Router) {
	server.Get("/api/v1/mail/check", mailCheck)
}

func mailCheck(w http.ResponseWriter, r *http.Request) {
	logging.Ctx(r.Context()).Info("mail health")

	dto := health.HealthResultDto{Status: "up"}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, dto)
}

func writeJson(ctx context.Context, w http.ResponseWriter, v interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		logging.Ctx(ctx).Warn(fmt.Sprintf("error while encoding json response: %v", err))
	}
}
