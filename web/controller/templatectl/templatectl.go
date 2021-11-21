package templatectl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eurofurence/reg-mail-service/api/v1/health"
	dto "github.com/eurofurence/reg-mail-service/api/v1/template"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/eurofurence/reg-mail-service/internal/service/templatesrv"
	"github.com/eurofurence/reg-mail-service/web/util/media"
	"github.com/go-chi/chi"
	"github.com/go-http-utils/headers"
)

var templateService templatesrv.TemplateService

func Create(server chi.Router) {
	server.Get("/api/v1/template/check", templateCheck)

	server.Route("/api/v1/template/{uuid}", func(r chi.Router) {
		r.Get("/", getTemplate)
		r.Post("/", updateTemplate)
		r.Delete("/", deleteTemplate)
	})
}

func templateCheck(w http.ResponseWriter, r *http.Request) {
	logging.Ctx(r.Context()).Info("templatectl health")

	dto := health.HealthResultDto{Status: "up"}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, dto)
}

func getTemplate(w http.ResponseWriter, r *http.Request) {
	// Get Template by UUID
	uuid := chi.URLParam(r, "uuid")

	template, err := templateService.GetTemplate(r.Context(), uuid)
	if err != nil {
		//templateNotFoundErrorHandler(ctx, w, r, uuid)
		return
	}

	dto := dto.TemplateDto{}
	mapTemplateToDto(template, &dto)
	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, dto)
}

func updateTemplate(w http.ResponseWriter, r *http.Request) {
	// Update Template and Create if it does not exist yet
}

func deleteTemplate(w http.ResponseWriter, r *http.Request) {
	// Remove Template by UUID
}

func writeJson(ctx context.Context, w http.ResponseWriter, v interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		logging.Ctx(ctx).Warn(fmt.Sprintf("error while encoding json response: %v", err))
	}
}
