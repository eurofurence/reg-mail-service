package mailctl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/service/mailsrv/v2"
	"github.com/eurofurence/reg-mail-service/internal/service/templatesrv/v2"
	"github.com/eurofurence/reg-mail-service/internal/web/filter"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctlutil"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctxvalues"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

var mailService mailsrv.MailService
var templateService templatesrv.TemplateService

func init() {
	mailService = &mailsrv.MailServiceImplData{}
	templateService = &templatesrv.TemplateServiceImplData{}
}

func Create(server chi.Router) {
	server.Post("/api/v2/tenants/{tenant}/mail", filter.HasGroupOrApiToken(config.OidcAdminGroup(), sendMail))
	server.Post("/api/v2/tenants/{tenant}/mail/preview", filter.HasGroupOrApiToken(config.OidcAdminGroup(), sendPreviewMail))
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tenantId := chi.URLParam(r, "tenant")

	if !slices.Contains(config.ServerTenants(), tenantId) {
		tenantNotFoundErrorHandler(ctx, w, r, fmt.Errorf("tenant '%s' not found in configuration", tenantId))
		return
	}

	// Parse the received body data to a "MailSendDto"
	dto, err := parseBodyToMailSendDto(ctx, w, r)
	if err != nil {
		mailParseErrorHandler(ctx, w, r, err)
		return
	}

	validationErrs := validate(ctx, dto)
	if len(validationErrs) != 0 {
		mailValidationErrorHandler(ctx, w, r, validationErrs)
		return
	}

	// Look up template by Common ID and Language
	// Falls back to en-US if language not found
	template, err := templateService.GetTemplateByCid(r.Context(), dto.CommonID, dto.Lang)
	if err != nil {
		mailParseErrorHandler(ctx, w, r, err)
		return
	}

	// Prepare E-Mail Text Content & Generate Body
	tempTextResult := template.Text

	for k, v := range dto.Variables {
		if strings.HasPrefix(k, "embed:") {
			continue
		} // Skip embeds in text-only mails
		tempTextResult = strings.ReplaceAll(tempTextResult, "{{ "+k+" }}", v)
	}

	// Prepare E-Mail HTML Content & Generate Body
	tempHTMLResult := template.HTML

	for k, v := range dto.Variables {
		if strings.HasPrefix(k, "embed:") {
			slice := strings.Split(k, ":")
			tempHTMLResult = strings.ReplaceAll(tempHTMLResult, "{{ "+k+" }}", fmt.Sprintf("<img src=\"data:%s;base64, %s\" />", slice[1], v))
		}
		tempHTMLResult = strings.ReplaceAll(tempHTMLResult, "{{ "+k+" }}", v)
	}

	// Set Developer Text
	if config.MailDevMode() {
		// Has to be in reverse order in order to have it at the top => Old body gets put under the new one
		tempTextResult = "\n ///// END OF DEVELOPMENT MESSAGE \\\\\\\\\\ \n\n" + tempTextResult
		tempTextResult = " To: " + strings.Join(dto.To, ";") + "\n\n CC: " + strings.Join(dto.Cc, ";") + "\n\n BCC: " + strings.Join(dto.Bcc, ";") + tempTextResult
		tempTextResult = "\n ///// MAIL DEVELOPMENT MODE ACTIVE \\\\\\\\\\ \n Original receivers:\n\n" + tempTextResult

		tempHTMLResult = "\n ///// END OF DEVELOPMENT MESSAGE \\\\\\\\\\ \n\n" + tempHTMLResult
		tempHTMLResult = " To: " + strings.Join(dto.To, ";") + "\n\n CC: " + strings.Join(dto.Cc, ";") + "\n\n BCC: " + strings.Join(dto.Bcc, ";") + tempHTMLResult
		tempHTMLResult = "\n ///// MAIL DEVELOPMENT MODE ACTIVE \\\\\\\\\\ \n Original receivers:\n\n" + tempHTMLResult
	}

	// Send the E-Mail
	err = mailService.SendMail(ctx, *dto, *template, tempTextResult, tempHTMLResult)

	if err != nil {
		mailServerErrorHandler(ctx, w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendPreviewMail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tenantId := chi.URLParam(r, "tenant")

	if !slices.Contains(config.ServerTenants(), tenantId) {
		tenantNotFoundErrorHandler(ctx, w, r, fmt.Errorf("tenant '%s' not found in configuration", tenantId))
		return
	}

	// Parse the received body data to a "MailSendDto"
	dto, err := parseBodyToMailSendDto(ctx, w, r)
	if err != nil {
		mailParseErrorHandler(ctx, w, r, err)
		return
	}

	// Ensure To, Cc, Bcc unset
	if len(dto.To) > 0 || len(dto.Cc) > 0 || len(dto.Bcc) > 0 {
		mailForbiddenErrorHandler(ctx, w, r, errors.New("mail preview not allowed to set target addresses!"))
		return
	}
	email := ctxvalues.Email(ctx)
	if email == "" {
		mailForbiddenErrorHandler(ctx, w, r, errors.New("you did not provide an email address in your token!"))
		return
	}
	dto.To = []string{email}

	// Look up template by Common ID and Language
	// Falls back to en-US if language not found
	template, err := templateService.GetTemplateByCid(r.Context(), dto.CommonID, dto.Lang)
	if err != nil {
		mailParseErrorHandler(ctx, w, r, err)
		return
	}

	// Prepare E-Mail Content & Generate Body
	tempResult := template.Data

	for k, v := range dto.Variables {
		tempResult = strings.ReplaceAll(tempResult, "{{ "+k+" }}", v)
	}

	// Send the E-Mail
	err = mailService.SendMail(ctx, *dto, *template, tempResult)

	if err != nil {
		mailServerErrorHandler(ctx, w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func parseBodyToMailSendDto(ctx context.Context, w http.ResponseWriter, r *http.Request) (*mail.MailSendDto, error) {
	decoder := json.NewDecoder(r.Body)
	dto := &mail.MailSendDto{}
	err := decoder.Decode(dto)
	if err != nil {
		mailParseErrorHandler(ctx, w, r, err)
	}
	return dto, err
}

// --- error handlers ---

func tenantNotFoundErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("tenant could not be found: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "tenant.notfound.error", http.StatusNotFound, url.Values{"error": {err.Error()}})
}

func mailValidationErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, errs url.Values) {
	aulogging.Logger.Ctx(ctx).Warn().Printf("received mail data with validation errors: %v", errs)
	ctlutil.ErrorHandler(ctx, w, r, "mail.data.invalid", http.StatusBadRequest, errs)
}

func mailServerErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("mail could not be sent: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "mail.server.error", http.StatusBadGateway, url.Values{"error": {err.Error()}})
}

func mailParseErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("mail send json body parse error: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "mail.parse.error", http.StatusBadRequest, url.Values{"error": {err.Error()}})
}

func mailForbiddenErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("mail send forbidden error: %s", err.Error())
	ctlutil.ErrorHandler(ctx, w, r, "mail.forbidden.error", http.StatusForbidden, url.Values{"error": {err.Error()}})
}
