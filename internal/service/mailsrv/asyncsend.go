package mailsrv

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"strings"
	"time"

	aulogging "github.com/StephanHCB/go-autumn-logging"
	auzerolog "github.com/StephanHCB/go-autumn-logging-zerolog"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/database"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctxvalues"
	"github.com/rs/zerolog/log"
	gomail "github.com/wneessen/go-mail"
)

func asyncContext(sourceCtx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx := context.Background()

	// child context so we do not cancel global context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	// provide values map in context
	ctx = ctxvalues.CreateContextWithValueMap(ctx)

	// carry over request id and logger
	requestId := ctxvalues.RequestId(sourceCtx)
	if requestId != "" {
		ctx = ctxvalues.SetRequestId(ctx, requestId)
		ctx = log.Logger.With().Str(auzerolog.RequestIdFieldName, requestId).Logger().WithContext(ctx)
	} else {
		ctx = log.Logger.With().Str(auzerolog.RequestIdFieldName, aulogging.DefaultRequestIdValue).Logger().WithContext(ctx)
	}

	// add timeout
	var expire context.CancelFunc
	ctx, expire = context.WithTimeout(ctx, timeout)

	return ctx, func() {
		expire()
		cancel()
	}
}

func asyncSend(sourceCtx context.Context, dto mail.MailSendDto, msg *gomail.Msg, body string) {
	sendCtx, sendCancel := asyncContext(sourceCtx, time.Second*20)

	failCtx, logCancel := asyncContext(sourceCtx, time.Second*60)

	aulogging.Logger.Ctx(failCtx).Info().Printf("asynchronously sending mail to %s ...", strings.Join(dto.To, ", "))
	go func() {
		defer sendCancel()
		defer logCancel()

		err := syncSend(sendCtx, dto, msg, body)
		if err != nil {
			aulogging.Logger.Ctx(failCtx).Warn().Printf("asynchronously sending mail failed: %s", err.Error())

			renderedRequest, err2 := json.Marshal(dto)
			if err2 != nil {
				aulogging.Logger.Ctx(failCtx).Error().Printf("rendering failed mail request to json also failed! Mail will be lost! %s", err2.Error())
				return
			}

			failureRecord := entity.Failure{
				CommonID: dto.CommonID,
				Language: dto.Lang,
				Request:  string(renderedRequest),
				Error:    err.Error(),
			}
			id, err3 := database.GetRepository().AddFailedMailRequest(failCtx, &failureRecord)
			if err3 != nil {
				aulogging.Logger.Ctx(failCtx).Error().Printf("adding failed mail request to db also failed! Mail will be lost! %s", err3.Error())
				return
			}
			aulogging.Logger.Ctx(failCtx).Info().Printf("failure successfully recorded with id %d", id)
			return
		}
		aulogging.Logger.Ctx(failCtx).Info().Printf("success sending mail to %s", strings.Join(dto.To, ", "))
		aulogging.Logger.Ctx(failCtx).Info().Print("async processing done")
	}()
}

func syncSend(ctx context.Context, dto mail.MailSendDto, msg *gomail.Msg, body string) (err error) {
	if !config.MailLogOnly() {
		var opts []gomail.Option
		opts = append(opts, gomail.WithPort(config.SmtpPort()))
		authIs := "disabled"
		if len(config.EmailFromPassword()) > 0 {
			opts = append(opts, gomail.WithSMTPAuth(gomail.SMTPAuthPlain), gomail.WithUsername(config.EmailFrom()), gomail.WithPassword(config.EmailFromPassword()))
			authIs = "enabled"
		}
		if config.SmtpInsecureSkipVerifyTLS() {
			opts = append(opts, gomail.WithTLSConfig(&tls.Config{
				InsecureSkipVerify: true,
			}))
		}

		client, err := gomail.NewClient(config.SmtpHost(), opts...)
		if err != nil {
			aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("failed to create email client: smtp host '%s', auth %s: %s", config.SmtpHost(), authIs, err.Error())
			_ = client.Close()
			return err
		}
		if err := client.DialAndSend(msg); err != nil {
			aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("failed to send email: %s", err.Error())
			_ = client.Close()
			return err
		}

		_ = client.Close()
		aulogging.Logger.Ctx(ctx).Info().Printf("Mail with template (%s/%s) sent.", dto.CommonID, dto.Lang)
	} else {
		aulogging.Logger.Ctx(ctx).Info().Printf("Mail body with template (%s/%s) logged below (**not** sent).", dto.CommonID, dto.Lang)
		aulogging.Logger.Ctx(ctx).Info().Printf(body)
	}
	return err
}
