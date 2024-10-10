package templatectl

import (
	"context"
	"encoding/json"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v2/template"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/service/templatesrv/v2"
	"github.com/eurofurence/reg-mail-service/internal/web/filter"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctlutil"
	"github.com/eurofurence/reg-mail-service/internal/web/util/media"
	"github.com/go-chi/chi/v5"
	"github.com/go-http-utils/headers"
	"net/http"
	"net/url"
	"slices"
)

var templateService templatesrv.TemplateService

func init() {
	templateService = &templatesrv.TemplateServiceImplData{}
}

func Create(server chi.Router) {
	server.Get("/api/v2/tenants/{tenant}/templates", filter.HasGroupOrApiToken(config.OidcAdminGroup(), getTemplates))
	server.Post("/api/v2/tenants/{tenant}/templates", filter.HasGroupOrApiToken(config.OidcAdminGroup(), createTemplate))
	server.Get("/api/v2/tenants/{tenant}/templates/{uuid}", filter.HasGroupOrApiToken(config.OidcAdminGroup(), getTemplate))
	server.Put("/api/v2/tenants/{tenant}/templates/{uuid}", filter.HasGroupOrApiToken(config.OidcAdminGroup(), updateTemplate))
	server.Delete("/api/v2/tenants/{tenant}/templates/{uuid}", filter.HasGroupOrApiToken(config.OidcAdminGroup(), deleteTemplate))
}

func getTemplates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tenantId := chi.URLParam(r, "tenant")

	if !slices.Contains(config.ServerTenants(), tenantId) {
		tenantNotFoundErrorHandler(ctx, w, r, fmt.Errorf("tenant '%s' not found in configuration", tenantId))
		return
	}

	cid := r.URL.Query().Get("cid")
	lang := r.URL.Query().Get("lang")

	templates, err := templateService.GetTemplates(ctx)
	if err != nil {
		templateDatabaseError(ctx, w, r, err)
		return
	}

	result := template.TemplateListDtoV2{Templates: make([]template.TemplateDtoV2, 0)}
	for _, tpl := range templates {
		if cid == "" || tpl.CommonID == cid {
			if lang == "" || tpl.Language == lang {
				dto := template.TemplateDtoV2{}
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
	tenantId := chi.URLParam(r, "tenant")

	if !slices.Contains(config.ServerTenants(), tenantId) {
		tenantNotFoundErrorHandler(ctx, w, r, fmt.Errorf("tenant '%s' not found in configuration", tenantId))
		return
	}

	dto, err := parseBodyToTemplateDto(ctx, w, r)
	if err != nil {
		templateParseErrorHandler(r.Context(), w, r, err)
		return
	}

	validationErrs := validate(ctx, dto)
	if len(validationErrs) != 0 {
		templateValidationErrorHandler(ctx, w, r, validationErrs)
		return
	}

	uuid, err := templateService.CreateTemplate(r.Context(), dto.CommonID, dto.Lang, dto.Subject, dto.Text, dto.HTML, dto.Attachments, dto.Embedded)
	if err != nil {
		templateParseErrorHandler(ctx, w, r, err)
		return
	}

	w.Header().Add(headers.Location, fmt.Sprintf("/api/v2/tenants/%s/templates/%s", tenantId, uuid))
	w.WriteHeader(http.StatusOK)
}

// Get Template by UUID
func getTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tenantId := chi.URLParam(r, "tenant")

	if !slices.Contains(config.ServerTenants(), tenantId) {
		tenantNotFoundErrorHandler(ctx, w, r, fmt.Errorf("tenant '%s' not found in configuration", tenantId))
		return
	}

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
	tenantId := chi.URLParam(r, "tenant")

	if !slices.Contains(config.ServerTenants(), tenantId) {
		tenantNotFoundErrorHandler(ctx, w, r, fmt.Errorf("tenant '%s' not found in configuration", tenantId))
		return
	}

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

	validationErrs := validate(ctx, dto)
	if len(validationErrs) != 0 {
		templateValidationErrorHandler(ctx, w, r, validationErrs)
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
	tenantId := chi.URLParam(r, "tenant")

	if !slices.Contains(config.ServerTenants(), tenantId) {
		tenantNotFoundErrorHandler(ctx, w, r, fmt.Errorf("tenant '%s' not found in configuration", tenantId))
		return
	}

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

func parseBodyToTemplateDto(ctx context.Context, w http.ResponseWriter, r *http.Request) (*template.TemplateDtoV2, error) {
	decoder := json.NewDecoder(r.Body)
	dto := &template.TemplateDtoV2{}
	err := decoder.Decode(dto)
	if err != nil {
		templateParseErrorHandler(ctx, w, r, err)
	}
	return dto, err
}

// --- error handlers ---

func tenantNotFoundErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("tenant could not be found: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "tenant.notfound.error", http.StatusNotFound, url.Values{"error": {err.Error()}})
}

func templateValidationErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, errs url.Values) {
	aulogging.Logger.Ctx(ctx).Warn().Printf("received mail data with validation errors: %v", errs)
	ctlutil.ErrorHandler(ctx, w, r, "template.invalid.error", http.StatusBadRequest, errs)
}

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
