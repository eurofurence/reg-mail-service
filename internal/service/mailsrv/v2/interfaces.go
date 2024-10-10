package mailsrv

import (
	"context"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/entity"
)

type MailService interface {
	SendMail(ctx context.Context, dto mail.MailSendDto, template entity.TemplateV2, textBody string, htmlBody) error
}
