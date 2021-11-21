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
	password := "pass"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Recipients
	to := []string{
		"example@example.com",
	}

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Load Template and fill with Variables, load from File Cache
	t, _ := template.ParseFiles("tmp/9b7585fd-4a2d-11ec-b88e-3431c4db8789.txt")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders))) // TODO: Replace with actual Subject from Template

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
}

func writeJson(ctx context.Context, w http.ResponseWriter, v interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		logging.Ctx(ctx).Warn(fmt.Sprintf("error while encoding json response: %v", err))
	}
}
