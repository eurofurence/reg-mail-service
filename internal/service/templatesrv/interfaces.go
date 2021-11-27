package templatesrv

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
)

type TemplateService interface {
	GetTemplates(ctx context.Context) (*entity.Template, error)

	UpdateTemplate(ctx context.Context, uuid string, data string) error
	DeleteTemplate(ctx context.Context, uuid string) error
	GetTemplate(ctx context.Context, id string) (*entity.Template, error)
	GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error)
}
