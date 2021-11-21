package templatesrv

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/database"
)

type TemplateServiceImplData struct {
}

func (s *TemplateServiceImplData) GetTemplates(ctx context.Context) (*entity.Template, error) {
	templates, err := database.GetRepository().GetTemplates(ctx)
	return templates, err
}

func (s *TemplateServiceImplData) GetTemplate(ctx context.Context, id string) (*entity.Template, error) {
	template, err := database.GetRepository().GetTemplateById(ctx, id)
	return template, err
}

func (s *TemplateServiceImplData) GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error) {
	template, err := database.GetRepository().GetTemplateByCid(ctx, cid, lang)
	return template, err
}