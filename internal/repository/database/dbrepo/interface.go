package dbrepo

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
)

type Repository interface {
	Open()
	Close()
	//Migrate()

	GetTemplates(ctx context.Context) (*entity.Template, error)

	DeleteTemplate(ctx context.Context, uuid string) error
	UpdateTemplate(ctx context.Context, uuid string, data string) error
	GetTemplateById(ctx context.Context, id string) (*entity.Template, error)
	GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error)

	//RecordHistory(ctx context.Context, h *entity.History) error
}
