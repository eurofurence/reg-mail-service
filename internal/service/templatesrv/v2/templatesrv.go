package templatesrv

import (
	"context"
	"github.com/google/uuid"

	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/database"
)

type TemplateServiceImplData struct {
}

func (s *TemplateServiceImplData) CreateTemplate(ctx context.Context, cid string, lang string, title string, text string, html string, attachments []string, embedded []string) (string, error) {
	var a entity.TemplateV2

	if lang == "" {
		lang = "en-US"
	}

	newId, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	a.ID = newId.String()
	a.CommonID = cid
	a.Language = lang
	a.Subject = title
	a.Text = text
	a.HTML = html
	a.Attachments = attachments
	a.Embedded = embedded

	err = database.GetRepository().CreateTemplate(ctx, &a)
	return a.ID, err
}

func (s *TemplateServiceImplData) UpdateTemplate(ctx context.Context, uuid string, data *entity.TemplateV2) error {
	err := database.GetRepository().UpdateTemplate(ctx, uuid, data)
	return err
}

func (s *TemplateServiceImplData) DeleteTemplate(ctx context.Context, uuid string) error {
	err := database.GetRepository().DeleteTemplate(ctx, uuid, true)
	return err
}

func (s *TemplateServiceImplData) GetTemplates(ctx context.Context) ([]*entity.TemplateV2, error) {
	templates, err := database.GetRepository().GetTemplates(ctx)
	return templates, err
}

func (s *TemplateServiceImplData) GetTemplate(ctx context.Context, id string) (*entity.TemplateV2, error) {
	template, err := database.GetRepository().GetTemplateById(ctx, id)
	return template, err
}

func (s *TemplateServiceImplData) GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.TemplateV2, error) {
	template, err := database.GetRepository().GetTemplateByCid(ctx, cid, lang)
	return template, err
}

func (s *TemplateServiceImplData) GetTenant(ctx context.Context, id string) (*entity.Tenant, error) {
	tenant, err := database.GetRepository().GetTenantById(ctx, id)
	return tenant, err
}
