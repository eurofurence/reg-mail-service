package templatectl

import (
	"context"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/api/v1/template"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/web/util/validation"
	"net/url"
)

const langPattern = "^[a-z]{2}-[A-Z]{2}$"

func validate(ctx context.Context, a *template.TemplateDto) url.Values {
	errs := url.Values{}

	validation.CheckLength(&errs, 1, 5, "lang", a.Lang)
	if validation.ViolatesPattern(langPattern, a.Lang) {
		errs.Add("lang", "language field is not plausible, must match "+langPattern+" (e.g.: de-DE)")
	}

	if len(a.CommonID) < 1 {
		errs.Add("template", "CommonID cannot be empty")
	}

	if len(a.Subject) < 1 {
		errs.Add("template", "Subject cannot be empty")
	}

	if len(a.Data) < 1 {
		errs.Add("template", "Data cannot be empty")
	}

	if len(errs) != 0 {
		if config.LoggingSeverity() == "DEBUG" {
			logger := aulogging.Logger.Ctx(ctx).Debug()
			for key, val := range errs {
				logger.Printf("template dto validation error for key %s: %s", key, val)
			}
		}
	}
	return errs
}
