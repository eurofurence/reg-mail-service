package templatesrv

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
)

type TemplateService interface {
	GetTemplate(ctx context.Context, id string) (*entity.Template, error)
}
