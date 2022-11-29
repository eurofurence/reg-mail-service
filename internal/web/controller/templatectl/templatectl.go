package templatectl

import (
	"context"
	"encoding/json"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/template"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/service/templatesrv"
	"github.com/eurofurence/reg-mail-service/internal/web/filter"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctlutil"
	"github.com/eurofurence/reg-mail-service/internal/web/util/media"
	"github.com/go-chi/chi/v5"
	"github.com/go-http-utils/headers"
	"net/http"
	"net/url"
)

var templateService templatesrv.TemplateService

func init() {
	templateService = &templatesrv.TemplateServiceImplData{}
}

func Create(server chi.Router) {
	server.Get("/api/v1/templates", filter.HasRoleOrApiToken(config.OidcAdminRole(), getTemplates))
	server.Post("/api/v1/templates", filter.HasRoleOrApiToken(config.OidcAdminRole(), createTemplate))
	server.Get("/api/v1/templates/{uuid}", filter.HasRoleOrApiToken(config.OidcAdminRole(), getTemplate))
	server.Put("/api/v1/templates/{uuid}", filter.HasRoleOrApiToken(config.OidcAdminRole(), updateTemplate))
	server.Delete("/api/v1/templates/{uuid}", filter.HasRoleOrApiToken(config.OidcAdminRole(), deleteTemplate))
}

func getTemplates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cid := r.URL.Query().Get("cid")
	lang := r.URL.Query().Get("lang")

	templates, err := templateService.GetTemplates(ctx)
	if err != nil {
		templateDatabaseError(ctx, w, r, err)
		return
	}

	result := template.TemplateListDto{Templates: make([]template.TemplateDto, 0)}
	for _, tpl := range templates {
		if cid == "" || tpl.CommonID == cid {
			if lang == "" || tpl.Language == lang {
				dto := template.TemplateDto{}
				mapTemplateToDto(tpl, &dto)
				result.Templates = append(result.Templates, dto)
			}
		}
	}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	ctlutil.WriteJson(r.Context(), w, result)
}

func createTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, err := parseBodyToTemplateDto(ctx, w, r)
	if err != nil {
		templateParseErrorHandler(r.Context(), w, r, err)
		return
	}

	uuid, err := templateService.CreateTemplate(r.Context(), dto.CommonID, dto.Lang, dto.Subject, dto.Data)
	if err != nil {
		templateParseErrorHandler(ctx, w, r, err)
		return
	}

	w.Header().Add(headers.Location, fmt.Sprintf("/api/v1/templates/%s", uuid))
	w.WriteHeader(http.StatusOK)
}

// Get Template by UUID
func getTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := chi.URLParam(r, "uuid")

	temp, err := templateService.GetTemplate(r.Context(), uuid)
	if err != nil {
		templateNotFoundErrorHandler(ctx, w, r, err)
		return
	}

	dto := template.TemplateDto{}
	mapTemplateToDto(temp, &dto)

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	ctlutil.WriteJson(r.Context(), w, dto)
}

// Update Template by UUID
func updateTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := chi.URLParam(r, "uuid")

	tpl, err := templateService.GetTemplate(r.Context(), uuid)
	if err != nil {
		templateNotFoundErrorHandler(ctx, w, r, err)
		return
	}

	dto, err := parseBodyToTemplateDto(ctx, w, r)
	if err != nil {
		templateParseErrorHandler(ctx, w, r, err)
		return
	}

	mapDtoToTemplate(dto, tpl)

	err = templateService.UpdateTemplate(r.Context(), uuid, tpl)
	if err != nil {
		templateDatabaseError(ctx, w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete Template by UUID
func deleteTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := chi.URLParam(r, "uuid")

	_, err := templateService.GetTemplate(r.Context(), uuid)
	if err != nil {
		templateNotFoundErrorHandler(ctx, w, r, err)
		return
	}

	err = templateService.DeleteTemplate(r.Context(), uuid)
	if err != nil {
		templateDatabaseError(ctx, w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
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

func templateDatabaseError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("template database error: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "template.database.error", http.StatusBadGateway, url.Values{"error": {err.Error()}})
}

func templateNotFoundErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("template could not be found: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "template.notfound.error", http.StatusNotFound, url.Values{"error": {err.Error()}})
}

func templateParseErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("template body could not be parsed: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "template.parse.error", http.StatusBadRequest, url.Values{"error": {err.Error()}})
}

func templateInvalidErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("template body invalid: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "template.invalid.error", http.StatusBadRequest, url.Values{"error": {err.Error()}})
}
