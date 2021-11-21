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
	//GetAttendeeById(ctx context.Context, id uint) (*entity.Attendee, error)
	GetTemplateById(ctx context.Context, id string) (*entity.Template, error)
	//CountAttendeesByNicknameZipEmail(ctx context.Context, nickname string, zip string, email string) (int64, error)
	//MaxAttendeeId(ctx context.Context) (uint, error)

	//RecordHistory(ctx context.Context, h *entity.History) error
}
