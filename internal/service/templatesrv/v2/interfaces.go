package templatesrv

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
)

type TemplateService interface {
	GetTemplates(ctx context.Context) ([]*entity.TemplateV2, error)

	CreateTemplate(ctx context.Context, cid string, lang string, title string, text string, html string, attachments []string, embedded []string) (string, error)
	UpdateTemplate(ctx context.Context, uuid string, data *entity.TemplateV2) error
	DeleteTemplate(ctx context.Context, uuid string) error
	GetTemplate(ctx context.Context, id string) (*entity.TemplateV2, error)
	GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.TemplateV2, error)
}
