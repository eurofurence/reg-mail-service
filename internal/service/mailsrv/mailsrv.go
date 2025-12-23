package mailsrv

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	gomail "github.com/wneessen/go-mail"
)

type MailServiceImplData struct {
}

func (s *MailServiceImplData) SendMail(ctx context.Context, dto mail.MailSendDto, template entity.Template, body string) error {
	// Create a new message and set sender
	m := gomail.NewMsg()
	aulogging.Logger.Ctx(ctx).Info().Printf("Preparing to send mail with template (%s/%s)...", dto.CommonID, dto.Lang)

	aulogging.Logger.Ctx(ctx).Info().Printf("From: %s", config.EmailFrom())
	if err := m.From(config.EmailFrom()); err != nil {
		aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("failed to set From: address '%s': %s", config.EmailFrom(), err.Error())
		return err
	}
	if config.MessageIdDomain() != "" {
		msgId := generateMessageId(ctx, config.MessageIdDomain())
		aulogging.Logger.Ctx(ctx).Info().Printf("Message-ID: %s", msgId)
		m.SetMessageIDWithValue(msgId)
	}
	// Avoid getting auto-response emails and triggering spam filters
	m.SetGenHeader("Precedence", "bulk")
	aulogging.Logger.Ctx(ctx).Info().Print("Precedence: bulk")

	// add extra Bcc if configured
	if config.AddAutoBcc() != "" {
		dto.Bcc = append(dto.Bcc, config.AddAutoBcc())
	}

	// Set recipients & Subject
	if config.MailDevMode() {
		aulogging.Logger.Ctx(ctx).Info().Printf("To: %s", strings.Join(config.MailDevMails(), ", "))
		// only log original To:
		aulogging.Logger.Ctx(ctx).Info().Printf("Original-To: %s", strings.Join(dto.To, ", "))
		if err := m.To(config.MailDevMails()...); err != nil {
			aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("failed to set To: addresses '%v': %s", config.MailDevMails(), err.Error())
			return err
		}
		aulogging.Logger.Ctx(ctx).Info().Printf("Subject: [regtest] %s", template.Subject)
		m.Subject(fmt.Sprintf("[regtest] %s", template.Subject))
		if len(dto.Cc) > 0 {
			// only log
			aulogging.Logger.Ctx(ctx).Info().Printf("Original-Cc: %s", strings.Join(dto.Cc, ", "))
		}
		if len(dto.Bcc) > 0 {
			// only log
			aulogging.Logger.Ctx(ctx).Info().Printf("Original-Bcc: %s", strings.Join(dto.Bcc, ", "))
		}
	} else {
		if len(dto.To) > 0 {
			aulogging.Logger.Ctx(ctx).Info().Printf("To: %s", strings.Join(dto.To, ", "))
			if err := m.To(dto.To...); err != nil {
				aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("failed to set To: addresses: %s", err.Error())
				return err
			}
		}
		aulogging.Logger.Ctx(ctx).Info().Printf("Subject: %s", template.Subject)
		m.Subject(template.Subject)

		if len(dto.Cc) > 0 {
			aulogging.Logger.Ctx(ctx).Info().Printf("Cc: %s", strings.Join(dto.Cc, ", "))
			if err := m.Cc(dto.Cc...); err != nil {
				aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("failed to set Cc: addresses '%v': %s", dto.Cc, err.Error())
				return err
			}
		}

		if len(dto.Bcc) > 0 {
			aulogging.Logger.Ctx(ctx).Info().Printf("Bcc: %s", strings.Join(dto.Bcc, ", "))
			if err := m.Bcc(dto.Bcc...); err != nil {
				aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("failed to set Bcc: addresses '%v': %s", dto.Bcc, err.Error())
				return err
			}
		}
	}

	// Set Body
	m.SetBodyString(gomail.TypeTextPlain, body)

	// Send E-Mail
	if dto.Async {
		asyncSend(ctx, dto, m, body)
		return nil
	} else {
		err := syncSend(ctx, dto, m, body)
		return err
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
