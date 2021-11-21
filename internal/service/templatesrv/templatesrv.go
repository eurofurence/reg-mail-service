package templatesrv

import (
	"context"

	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/database"
)

type TemplateServiceImplData struct {
}

func (s *TemplateServiceImplData) GetTemplate(ctx context.Context, id string) (*entity.Template, error) {
	template, err := database.GetRepository().GetTemplateById(ctx, id)
	return template, err
}
