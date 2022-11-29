package historizeddb

import (
	"context"
	"errors"

	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/dbrepo"
)

type HistorizingRepository struct {
	wrappedRepository dbrepo.Repository
}

func Create(wrappedRepository dbrepo.Repository) dbrepo.Repository {
	return &HistorizingRepository{wrappedRepository: wrappedRepository}
}

func (r *HistorizingRepository) Open() {
	r.wrappedRepository.Open()
}

func (r *HistorizingRepository) Close() {
	r.wrappedRepository.Close()
}

func (r *HistorizingRepository) GetTemplates(ctx context.Context) (*entity.Template, error) {
	return r.wrappedRepository.GetTemplates(ctx)
}

func (r *HistorizingRepository) CreateTemplate(ctx context.Context, cid string, lang string, title string, data string) error {
	panic("implement me")
}

func (r *HistorizingRepository) DeleteTemplate(ctx context.Context, uuid string, permanent bool) error {
	panic("implement me")
}

func (r *HistorizingRepository) UpdateTemplate(ctx context.Context, uuid string, data string) error {
	panic("implement me")
}

func (r *HistorizingRepository) GetTemplateById(ctx context.Context, id string) (*entity.Template, error) {
	return r.wrappedRepository.GetTemplateById(ctx, id)
}

func (r *HistorizingRepository) GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error) {
	panic("implement me")
}

func (r *HistorizingRepository) RecordHistory(ctx context.Context, h *entity.History) error {
	return errors.New("not allowed to directly manipulate history")
}
