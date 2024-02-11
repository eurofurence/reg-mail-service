package mailsrv

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/Shopify/gomail"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"strings"
	"time"
)

type MailServiceImplData struct {
}

func (s *MailServiceImplData) SendMail(ctx context.Context, dto mail.MailSendDto, template entity.Template, body string) error {
	// Create a new message and set sender
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailFrom())
	if config.MessageIdDomain() != "" {
		m.SetHeader("Message-ID", generateMessageId(ctx, config.MessageIdDomain()))
	}

	// Set recipients & Subject
	if config.MailDevMode() {
		m.SetHeader("To", config.MailDevMails()...)
		m.SetHeader("Subject", fmt.Sprintf("[regtest] %s", template.Subject))
	} else {
		m.SetHeader("To", dto.To...)
		m.SetHeader("Subject", template.Subject)

		if len(dto.Cc) > 0 {
			m.SetHeader("Cc", dto.Cc...)
		}

		// add extra Bcc if configured
		if config.AddAutoBcc() != "" {
			dto.Bcc = append(dto.Bcc, config.AddAutoBcc())
		}

		if len(dto.Bcc) > 0 {
			m.SetHeader("Bcc", dto.Bcc...)
		}
	}

	// Set Body
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

const messageIdTimestampFormat = "20060102150405.000"

var fallbackToken uint8 = 0

func generateMessageId(ctx context.Context, domain string) string {
	timestamp := time.Now().Format(messageIdTimestampFormat)

	token := make([]byte, 4)
	_, err := rand.Read(token)
	if err != nil {
		aulogging.Logger.Ctx(ctx).Warn().Printf("failed to generate random token for message id - using counter")
		f := fallbackToken
		token = []byte{0, 0, 0, f}
		if f == 255 {
			fallbackToken = 0
		} else {
			fallbackToken = f + 1
		}
	}
	hexToken := hex.EncodeToString(token)

	return fmt.Sprintf("<%s.%s@%s>", timestamp, hexToken, domain)
}
