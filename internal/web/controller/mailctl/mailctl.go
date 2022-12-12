package mailctl

import (
	"context"
	"encoding/json"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/health"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/web/filter"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctlutil"
	"github.com/eurofurence/reg-mail-service/internal/web/util/media"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"

	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/service/templatesrv"
	"github.com/go-chi/chi/v5"
	"github.com/go-http-utils/headers"
)

var templateService templatesrv.TemplateService

func init() {
	templateService = &templatesrv.TemplateServiceImplData{}
}

func Create(server chi.Router) {
	server.Post("/api/v1/mail", filter.HasRoleOrApiToken(config.OidcAdminRole(), sendMail))

	server.Get("/api/v1/mail/check", checkHealth)
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	dto := health.HealthResultDto{Status: "up"}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	ctlutil.WriteJson(r.Context(), w, dto)
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse the received body data to a "MailSendDto"
	dto, err := parseBodyToMailSendDto(ctx, w, r)
	if err != nil {
		mailParseErrorHandler(r.Context(), w, r, err)
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
		mailParseErrorHandler(r.Context(), w, r, err)
		return
	}

	// Recipients
	recipients := append(dto.To, dto.Cc...)

	// Sender
	from := config.EmailFrom()
	password := config.EmailFromPassword()
	smtpHost := config.SmtpHost()
	smtpPort := config.SmtpPort()

	// Prepare E-Mail Content
	tempResult := template.Data

	for k, v := range dto.Variables {
		tempResult = strings.ReplaceAll(tempResult, "{{ "+k+" }}", v)
	}

	/// Start of the E-Mail Body
	// Override the recipients to the pre-set developer mails in the config, if development mode is active
	body := []byte("")
	if config.MailDevMode() {
		body = []byte("To: " + strings.Join(config.MailDevMails(), ";") + "\r\n")
	} else {
		body = []byte("To: " + strings.Join(dto.To, ";") + "\r\n")
	}

	// Add Blind CC if any available (and not in Development Mode)
	if len(dto.Cc) > 0 && !config.MailDevMode() {
		body = append(body, []byte("Cc: "+strings.Join(dto.Cc, ";")+"\r\n")...)
	}

	// Add Blind CC if any available (and not in Development Mode)
	if len(dto.Bcc) > 0 && !config.MailDevMode() {
		body = append(body, []byte("Bcc: "+strings.Join(dto.Bcc, ";")+"\r\n")...)
	}

	// Title and Content of E-Mail
	// Add Development Mode content if Mail Development Mode is active in config
	body = append(body, []byte("Subject: "+template.Subject+"\r\n"+"\r\n"+tempResult+"\r\n")...)
	if config.MailDevMode() {
		body = append(body, []byte("\r\n ///// MAIL DEVELOPMENT MODE ACTIVE \\\\\\\\\\ \r\n Original receivers:\n\n")...)
		body = append(body, []byte("To: "+strings.Join(dto.To, ";")+"\n\nCC: "+strings.Join(dto.Cc, ";")+"\n\nBCC: "+strings.Join(dto.Bcc, ";"))...)
	}
	/// End of the E-Mail Body

	// Send the finished E-Mail
	// Ignore Authentication if it should only be logged into the console
	if !config.MailLogOnly() {
		// Authentication
		auth := smtp.PlainAuth("", from, password, smtpHost)

		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, body)
		if err != nil {
			mailServerErrorHandler(r.Context(), w, r, err)
			return
		}
		aulogging.Logger.Ctx(r.Context()).Info().Printf("Mail with template (%s/%s) sent. TO: %s. CC: %s. BCC: %s",
			dto.CommonID, dto.Lang, dto.To, dto.Cc, dto.Bcc)
	} else {
		aulogging.Logger.Ctx(r.Context()).Info().Printf("Mail body with template (%s/%s) logged below (**not** sent).", dto.CommonID, dto.Lang)
		aulogging.Logger.Ctx(r.Context()).Info().Printf(string(body))
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
