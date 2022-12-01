package templatectl

import (
	dto "github.com/eurofurence/reg-mail-service/internal/api/v1/template"
	"github.com/eurofurence/reg-mail-service/internal/entity"
)

func mapTemplateToDto(a *entity.Template, dto *dto.TemplateDto) {
	// this cannot fail
	dto.UUID = a.ID
	dto.CommonID = a.CommonID
	dto.Lang = a.Language
	dto.Subject = a.Subject
	dto.Data = a.Data
	dto.CreatedAt = a.CreatedAt
	dto.UpdatedAt = a.UpdatedAt
}

func mapDtoToTemplate(dto *dto.TemplateDto, a *entity.Template) {
	a.ID = dto.UUID
	a.CommonID = dto.CommonID
	a.Language = dto.Lang
	a.Subject = dto.Subject
	a.Data = dto.Data
}
