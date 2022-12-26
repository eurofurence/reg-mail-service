package mailsrv

import (
	"context"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"gopkg.in/gomail.v2"
)

type MailServiceImplData struct {
}

func (s *MailServiceImplData) SendMail(ctx context.Context, dto mail.MailSendDto, template entity.Template, body string) error {
	// Create a new message and set sender
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailFrom())

	// Set recipients
	if config.MailDevMode() {
		m.SetHeader("To", config.MailDevMails()...)
	} else {
		m.SetHeader("To", dto.To...)
	}

	if !config.MailDevMode() {
		m.SetHeader("Cc", dto.Cc...)
		m.SetHeader("Bcc", dto.Bcc...)
	}

	// Set subject & Body
	m.SetHeader("Subject", template.Subject)
	m.SetBody("text/plain", body)

	// Send E-Mail
	err := error(nil)
	if !config.MailLogOnly() {
		d := gomail.NewDialer(config.SmtpHost(), config.SmtpPort(), config.EmailFrom(), config.EmailFromPassword())
		err = d.DialAndSend(m)

		aulogging.Logger.Ctx(ctx).Info().Printf("Mail with template (%s/%s) sent. TO: %s. CC: %s. BCC: %s",
			dto.CommonID, dto.Lang, dto.To, dto.Cc, dto.Bcc)
	} else {
		aulogging.Logger.Ctx(ctx).Info().Printf("Mail body with template (%s/%s) logged below (**not** sent).", dto.CommonID, dto.Lang)
		aulogging.Logger.Ctx(ctx).Info().Printf(body)
	}

	return err
}
