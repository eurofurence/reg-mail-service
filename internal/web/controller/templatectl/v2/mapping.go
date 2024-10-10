package templatectl

import (
	dto "github.com/eurofurence/reg-mail-service/internal/api/v2/template"
	"github.com/eurofurence/reg-mail-service/internal/entity"
)

func mapTemplateToDto(a *entity.TemplateV2, dto *dto.TemplateDtoV2) {
	// this cannot fail
	dto.UUID = a.ID
	dto.CommonID = a.CommonID
	dto.Lang = a.Language
	dto.Subject = a.Subject
	dto.Text = a.Text
	dto.HTML = a.HTML
	dto.Attachments = a.Attachments
	dto.Embedded = a.Embedded
	dto.CreatedAt = a.CreatedAt
	dto.UpdatedAt = a.UpdatedAt
}

func mapDtoToTemplate(dto *dto.TemplateDtoV2, a *entity.TemplateV2) {
	a.ID = dto.UUID
	a.CommonID = dto.CommonID
	a.Language = dto.Lang
	a.Subject = dto.Subject
	a.Text = dto.Text
	a.HTML = dto.HTML
	a.Attachments = dto.Attachments
	a.Embedded = dto.Embedded
}
