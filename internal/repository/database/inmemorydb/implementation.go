package inmemorydb

import (
	"context"
	"errors"
	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/dbrepo"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type InmemoryRepository struct {
	internalStore map[string]*entity.Template
}

func Create() dbrepo.Repository {
	return &InmemoryRepository{}
}

func (r *InmemoryRepository) Open() {
	r.internalStore = make(map[string]*entity.Template)
}

func (r *InmemoryRepository) Close() {
	r.internalStore = nil
}

func (r *InmemoryRepository) GetTemplates(ctx context.Context) (*entity.Template, error) {
	for _, value := range r.internalStore {
		copiedTemplate := *value
		return &copiedTemplate, nil
	}
	return &entity.Template{}, errors.New("no templates found in db")

	// should return a slice with all templates
}

func (r *InmemoryRepository) CreateTemplate(ctx context.Context, cid string, lang string, subject string, data string) error {
	// should take a filled entity and just add the id
	// also note, GORM ids are normally uint, not uuids

	newId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	newTemplate := &entity.Template{
		Base: entity.Base{
			ID:        newId.String(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		},
		CommonID: cid,
		Language: lang,
		Subject:  subject,
		Data:     data,
	}

	r.internalStore[newId.String()] = newTemplate

	// need to return id
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

func (r *InmemoryRepository) UpdateTemplate(ctx context.Context, uuid string, data string) error {
	// should take an updated entity, or needs subject
	return nil
}

func (r *InmemoryRepository) GetTemplateById(ctx context.Context, id string) (*entity.Template, error) {
	template, ok := r.internalStore[id]
	if !ok {
		return &entity.Template{}, errors.New("template not found")
	}
	copy := *template
	return &copy, nil
}

func (r *InmemoryRepository) GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error) {
	for id, tpl := range r.internalStore {
		if tpl.CommonID == cid && tpl.Language == lang {
			return r.GetTemplateById(ctx, id)
		}
	}
	return &entity.Template{}, errors.New("no template found")
}
