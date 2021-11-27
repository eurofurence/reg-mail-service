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

func init() {
	templateService = &templatesrv.TemplateServiceImplData{}
}

func Create(server chi.Router) {
	server.Get("/api/v1/template/check", templateCheck)

	server.Get("/api/v1/template", getTemplateByCid)

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

func getTemplates(w http.ResponseWriter, r *http.Request) {

}

// Get Template by UUID
func getTemplate(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	template, err := templateService.GetTemplate(r.Context(), uuid)
	if err != nil {
		templateNotFoundErrorHandler(r.Context(), uuid)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dto := dto.TemplateDto{}
	mapTemplateToDto(template, &dto)

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, dto)
}

// Update Template by UUID
func updateTemplate(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	data := r.Header.Get("data")

	err := templateService.UpdateTemplate(r.Context(), uuid, data)
	if err != nil {
		templateNotFoundErrorHandler(r.Context(), uuid)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete Template by UUID
func deleteTemplate(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	err := templateService.DeleteTemplate(r.Context(), uuid)
	if err != nil {
		templateNotFoundErrorHandler(r.Context(), uuid)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get Template by Common ID
func getTemplateByCid(w http.ResponseWriter, r *http.Request) {
	cid := r.Header.Get("cid")
	lang := r.Header.Get("lang")

	template, err := templateService.GetTemplateByCid(r.Context(), cid, lang)
	if err != nil {
		//templateNotFoundErrorHandler(r.Context(), uuid)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dto := dto.TemplateDto{}
	mapTemplateToDto(template, &dto)

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, dto)
}

func templateNotFoundErrorHandler(ctx context.Context, uuid string) {
	logging.Ctx(ctx).Warn("template uuid ", uuid, " not found")
}

func writeJson(ctx context.Context, w http.ResponseWriter, v interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		logging.Ctx(ctx).Warn(fmt.Sprintf("error while encoding json response: %v", err))
	}
}
