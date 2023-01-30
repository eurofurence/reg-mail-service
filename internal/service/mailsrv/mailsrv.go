package mailsrv

import (
	"context"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"gopkg.in/gomail.v2"
	"strings"
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

		if len(dto.Cc) > 0 {
			m.SetHeader("Cc", dto.Cc...)
		}

		if len(dto.Bcc) > 0 {
			m.SetHeader("Bcc", dto.Bcc...)
		}
	}

	// Set subject & Body
	m.SetHeader("Subject", template.Subject)
	m.SetBody("text/plain", body)

	// Send E-Mail
	err := error(nil)
	if !config.MailLogOnly() {
		d := &gomail.Dialer{Host: config.SmtpHost(), Port: config.SmtpPort()}

		if len(config.EmailFromPassword()) > 0 {
			d.Username = config.EmailFrom()
			d.Password = config.EmailFromPassword()
		}

		err = d.DialAndSend(m)

		aulogging.Logger.Ctx(ctx).Info().Printf("Mail with template (%s/%s) sent.", dto.CommonID, dto.Lang)
		logTargets(ctx, dto)
		aulogging.Logger.Ctx(ctx).Info().Printf("Subject: %s", template.Subject)
	} else {
		aulogging.Logger.Ctx(ctx).Info().Printf("Mail body with template (%s/%s) logged below (**not** sent).", dto.CommonID, dto.Lang)
		logTargets(ctx, dto)
		aulogging.Logger.Ctx(ctx).Info().Printf("Subject: %s", template.Subject)
		aulogging.Logger.Ctx(ctx).Info().Printf(body)
	}

	return err
}

func logTargets(ctx context.Context, dto mail.MailSendDto) {
	if len(dto.To) > 0 {
		aulogging.Logger.Ctx(ctx).Info().Printf("To: %s", strings.Join(dto.To, ", "))
	}
	if len(dto.Cc) > 0 {
		aulogging.Logger.Ctx(ctx).Info().Printf("Cc: %s", strings.Join(dto.Cc, ", "))
	}
	if len(dto.Bcc) > 0 {
		aulogging.Logger.Ctx(ctx).Info().Printf("Bcc: %s", strings.Join(dto.Bcc, ", "))
	}
}