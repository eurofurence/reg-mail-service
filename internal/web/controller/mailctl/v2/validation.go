package mailctl

import (
	"context"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/mail"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/web/util/validation"
	"net/url"
)

const emailPattern = `^[^@\s]+@[^@\s]+$`
const langPattern = "^[a-z]{2}-[A-Z]{2}$"

func validate(ctx context.Context, a *mail.MailSendDto) url.Values {
	errs := url.Values{}

	validation.CheckLength(&errs, 1, 5, "lang", a.Lang)
	if validation.ViolatesPattern(langPattern, a.Lang) {
		errs.Add("lang", "language field is not plausible, must match "+langPattern+" (e.g.: de-DE)")
	}

	if len(a.To) > 0 {
		for _, email := range a.To {
			validation.CheckLength(&errs, 1, 200, "email", email)
			if validation.ViolatesPattern(emailPattern, email) {
				errs.Add("email", "to (recipient) field is not plausible, must match "+emailPattern)
			}
		}
	} else {
		errs.Add("email", "to (recipient) field cannot be empty")
	}

	if len(a.Cc) > 0 {
		for _, email := range a.Cc {
			validation.CheckLength(&errs, 1, 200, "email", email)
			if validation.ViolatesPattern(emailPattern, email) {
				errs.Add("email", "to (recipient) field is not plausible, must match "+emailPattern)
			}
		}
	}

	if len(a.Bcc) > 0 {
		for _, email := range a.Bcc {
			validation.CheckLength(&errs, 1, 200, "email", email)
			if validation.ViolatesPattern(emailPattern, email) {
				errs.Add("email", "to (recipient) field is not plausible, must match "+emailPattern)
			}
		}
	}

	if len(errs) != 0 {
		if config.LoggingSeverity() == "DEBUG" {
			logger := aulogging.Logger.Ctx(ctx).Debug()
			for key, val := range errs {
				logger.Printf("mail dto validation error for key %s: %s", key, val)
			}
		}
	}
	return errs
}
