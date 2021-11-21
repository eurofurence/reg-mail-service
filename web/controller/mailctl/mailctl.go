package mailctl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/eurofurence/reg-mail-service/api/v1/health"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/eurofurence/reg-mail-service/web/util/media"
	"github.com/go-chi/chi"
	"github.com/go-http-utils/headers"
)

func Create(server chi.Router) {
	server.Get("/api/v1/mail/check", mailCheck)

	server.Post("/api/v1/mail/send", sendTemplate)
}

func mailCheck(w http.ResponseWriter, r *http.Request) {
	logging.Ctx(r.Context()).Info("mail health")

	dto := health.HealthResultDto{Status: "up"}

	w.Header().Add(headers.ContentType, media.ContentTypeApplicationJson)
	w.WriteHeader(http.StatusOK)
	writeJson(r.Context(), w, dto)
}

func sendTemplate(w http.ResponseWriter, r *http.Request) {
	// Template
	cid := r.Header.Get("cid")
	lang := r.Header.Get("lang")

	// TODO: Fetch Template (either from DB or Cache, if implemented) and store inside variable
	// to be able to get also the Title and other Information from the template.
	// In this case, CID lookup has to be used, since the language code is passed to this service.

	// TODO: Should cache be implemented?
	//t, _ := template.ParseFiles("assets/cache/de_DE/guest.txt")

	// Recipients
	recipients := []string{
		r.Header.Get("recipient"),
	}

	// Sender
	from := config.EmailFrom()
	password := config.EmailFromPassword()
	smtpHost := config.SmtpHost()
	smtpPort := config.SmtpPort()

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	var body bytes.Buffer

	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", media.ContentMimeHeaders))) // TODO: Replace with actual Subject from Template

	// TODO: Read Template JSON => Generate Struct => Fill Variables?
	//t.Execute(&body, struct {
	//	Var1 string
	//		Var2 string
	//	}{
	//		Var1: "Foo Bar",
	//		Var2: "This is a test message in a Plain-Text template",
	//	})

	// Send
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, body.Bytes())
	if err != nil {
		logging.Ctx(r.Context()).Error(err)
		return
	}
	logging.Ctx(r.Context()).Info("Mail with template (", cid, "/", lang, ") sent to: ", recipients)

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
