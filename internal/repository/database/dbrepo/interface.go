package dbrepo

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
)

type Repository interface {
	Open()
	Close()
	//Migrate()

	//AddAttendee(ctx context.Context, a *entity.Attendee) (uint, error)
	//UpdateAttendee(ctx context.Context, a *entity.Attendee) error
	GetTemplates(ctx context.Context) (*entity.Template, error)
	//GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error)
	GetTemplateById(ctx context.Context, id string) (*entity.Template, error)

	//RecordHistory(ctx context.Context, h *entity.History) error
}
