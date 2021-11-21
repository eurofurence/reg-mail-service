package templatectl

import (
	dto "github.com/eurofurence/reg-mail-service/api/v1/template"
	"github.com/eurofurence/reg-mail-service/internal/entity"
)

func mapTemplateToDto(a *entity.Template, dto *dto.TemplateDto) {
	// this cannot fail
	//dto.id = fmt.Sprint(a.id)
	dto.Lang = a.Language
	dto.Title = a.Title
	dto.Data = a.Data
	dto.CreatedAt = a.CreatedAt
	dto.UpdatedAt = a.UpdatedAt
}
