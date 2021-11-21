package mailctl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"text/template"

	"github.com/eurofurence/reg-mail-service/api/v1/health"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/eurofurence/reg-mail-service/web/util/media"
	"github.com/go-chi/chi"
	"github.com/go-http-utils/headers"
)

func Create(server chi.Router) {
	server.Get("/api/v1/mail/check", mailCheck)

	server.Post("/api/v1/mail/sendtemplate", sendTemplate)
}

func mailCheck(w http.ResponseWriter, r *http.Request) {
	logging.Ctx(r.Context()).Info("mail health")

	dto := health.HealthResultDto{Status: "up"}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, dto)
}

func sendTemplate(w http.ResponseWriter, r *http.Request) {
	// Sender
	from := config.EmailFrom()
	password := config.EmailFromPassword()
	smtpHost := config.SmtpHost()
	smtpPort := config.SmtpPort()

	// Recipients
	to := []string{
		"example@example.com",
	}

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Load Template and fill with Variables, load from File Cache
	// TODO: Check if File is in Cache and check Database if Cache is Up to Date
	t, _ := template.ParseFiles("assets/cache/de_DE/guest.txt")

	var body bytes.Buffer

	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", media.ContentMimeHeaders))) // TODO: Replace with actual Subject from Template

	// TODO: Read Template JSON => Generate Struct => Fill Variables?
	t.Execute(&body, struct {
		Var1 string
		Var2 string
	}{
		Var1: "Foo Bar",
		Var2: "This is a test message in a Plain-Text template",
	})

	// Send
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		logging.Ctx(r.Context()).Error(err)
		return
	}
	logging.Ctx(r.Context()).Info("Template Sent Successfully!")

	w.WriteHeader(http.StatusOK)
}

func writeJson(ctx context.Context, w http.ResponseWriter, v interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		logging.Ctx(ctx).Warn(fmt.Sprintf("error while encoding json response: %v", err))
	}
}
