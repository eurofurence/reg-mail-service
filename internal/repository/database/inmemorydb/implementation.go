package inmemorydb

import (
	"context"
	"errors"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/dbrepo"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/templates"
	"sort"
	"time"
)

type InmemoryRepository struct {
	internalStore map[string]*entity.Template
}

func Create() dbrepo.Repository {
	return &InmemoryRepository{}
}

func (r *InmemoryRepository) Open() error {
	r.internalStore = make(map[string]*entity.Template)
	return nil
}

func (r *InmemoryRepository) Close() {
	r.internalStore = nil
}

func (r *InmemoryRepository) Migrate() error {
	_ = templates.SeedDefaultTemplates(context.Background(), r)
	return nil
}

func (r *InmemoryRepository) GetTemplates(ctx context.Context) ([]*entity.Template, error) {
	result := make([]*entity.Template, 0)
	for _, value := range r.internalStore {
		copiedTemplate := *value
		result = append(result, &copiedTemplate)
	}
	sort.Slice(result, func(i, j int) bool {
		return templateLess(result[i], result[j])
	})
	return result, nil
}

func templateLess(a *entity.Template, b *entity.Template) bool {
	if a == nil || b == nil {
		return b != nil
	}
	if a.CommonID != b.CommonID {
		return a.CommonID < b.CommonID
	}
	return a.Language < b.Language
}

func (r *InmemoryRepository) CreateTemplate(ctx context.Context, tpl *entity.Template) error {
	// should take a filled entity and just add the id
	// also note, GORM ids are normally uint, not uuids

	tpl.CreatedAt = time.Time{}
	tpl.UpdatedAt = time.Time{}

	tplCopy := *tpl
	r.internalStore[tpl.ID] = &tplCopy

	return nil
}

// Note: The DeleteTemplate function does not really delete a database entry.
// Rather it sets the 'deleted_at' timestamp, which results it not being found on the
// get/update queries.
// This could be useful for some sort of "Archive" in the Dashboard for the Admins to restore
// deleted templates or as some sort of Backup.
func (r *InmemoryRepository) DeleteTemplate(ctx context.Context, uuid string, permanent bool) error {
	_, err := r.GetTemplateById(ctx, uuid)
	if err != nil {
		return err
	}
	if permanent {
		delete(r.internalStore, uuid)
	} else {
		t := time.Now()
		r.internalStore[uuid].DeletedAt = &t
	}
	return nil
}

func (r *InmemoryRepository) UpdateTemplate(ctx context.Context, uuid string, data *entity.Template) error {
	tpl, err := r.GetTemplateById(ctx, uuid)
	if err != nil {
		return err
	}

	tpl.Data = data.Data
	tpl.Subject = data.Subject
	tpl.CommonID = data.CommonID
	tpl.Language = data.Language

	// TODO ensure uniqueness?

	r.internalStore[tpl.ID] = tpl

	return nil
}

func (r *InmemoryRepository) GetTemplateById(ctx context.Context, id string) (*entity.Template, error) {
	template, ok := r.internalStore[id]
	if !ok {
		return &entity.Template{}, errors.New("template not found")
	}
	templateCopy := *template
	return &templateCopy, nil
}

func (r *InmemoryRepository) GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error) {
	for id, tpl := range r.internalStore {
		if tpl.CommonID == cid && tpl.Language == lang {
			return r.GetTemplateById(ctx, id)
		}
	}
	return &entity.Template{}, errors.New("no template found")
}
