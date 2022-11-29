package templatesrv

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
)

type TemplateService interface {
	GetTemplates(ctx context.Context) ([]*entity.Template, error)

	CreateTemplate(ctx context.Context, cid string, lang string, title string, data string) (string, error)
	UpdateTemplate(ctx context.Context, uuid string, data *entity.Template) error
	DeleteTemplate(ctx context.Context, uuid string, permanent bool) error
	GetTemplate(ctx context.Context, id string) (*entity.Template, error)
	GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error)
}
