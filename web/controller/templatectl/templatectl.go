package templatectl

import (
	"context"
	"encoding/json"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/api/v1/template"
	"github.com/eurofurence/reg-mail-service/web/util/ctlutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/eurofurence/reg-mail-service/api/v1/health"
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
	server.Route("/api/v1/templates", func(r chi.Router) {
		r.Get("/", getTemplates)
		r.Post("/", createTemplate)
	})

	server.Route("/api/v1/templates/{uuid}", func(r chi.Router) {
		r.Get("/", getTemplate)
		r.Put("/", updateTemplate)
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

func createTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, err := parseBodyToTemplateDto(ctx, w, r)
	if err != nil {
		templateParseErrorHandler(r.Context(), w, r, err)
		return
	}

	err = templateService.CreateTemplate(r.Context(), dto.CommonID, dto.Lang, dto.Subject, dto.Data)
	if err != nil {
		templateInvalidErrorHandler(ctx, w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get Template by UUID
func getTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, err := parseBodyToTemplateDto(ctx, w, r)
	if err != nil {
		templateParseErrorHandler(r.Context(), w, r, err)
		return
	}

	temp, err := templateService.GetTemplate(r.Context(), dto.UUID)
	if err != nil {
		templateInvalidErrorHandler(ctx, w, r, err)
		return
	}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, temp)
}

// Update Template by UUID
func updateTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := chi.URLParam(r, "uuid")
	data := r.Header.Get("data")

	err := templateService.UpdateTemplate(r.Context(), uuid, data)
	if err != nil {
		templateInvalidErrorHandler(ctx, w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete Template by UUID
func deleteTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, err := parseBodyToTemplateDto(ctx, w, r)
	if err != nil {
		templateParseErrorHandler(r.Context(), w, r, err)
		return
	}

	permanent, err := strconv.ParseBool(r.Header.Get("permanent"))
	if err != nil {
		templateParseErrorHandler(ctx, w, r, err)
		return
	}

	err = templateService.DeleteTemplate(r.Context(), dto.UUID, permanent)
	if err != nil {
		templateInvalidErrorHandler(ctx, w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get Template by Common ID
func getTemplates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, err := parseBodyToTemplateDto(ctx, w, r)
	if err != nil {
		templateParseErrorHandler(r.Context(), w, r, err)
		return
	}

	temp, err := templateService.GetTemplateByCid(r.Context(), dto.CommonID, dto.Lang)
	if err != nil {
		templateInvalidErrorHandler(ctx, w, r, err)
		return
	}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, temp)
}

func writeJson(ctx context.Context, w http.ResponseWriter, v interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		logging.Ctx(ctx).Warn(fmt.Sprintf("error while encoding json response: %v", err))
	}
}

func parseBodyToTemplateDto(ctx context.Context, w http.ResponseWriter, r *http.Request) (*template.TemplateDto, error) {
	decoder := json.NewDecoder(r.Body)
	dto := &template.TemplateDto{}
	err := decoder.Decode(dto)
	if err != nil {
		templateParseErrorHandler(ctx, w, r, err)
	}
	return dto, err
}

// --- error handlers ---

func templateServerError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("template database error: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "template.server.error", http.StatusBadGateway, url.Values{"error": {err.Error()}})
}

func templateInvalidErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("template could not be found: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "template.invalid.error", http.StatusNotFound, url.Values{"error": {err.Error()}})
}

func templateParseErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("template body could not be parsed: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "template.parse.error", http.StatusBadRequest, url.Values{"error": {err.Error()}})
}
