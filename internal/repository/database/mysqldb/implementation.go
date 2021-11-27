package mysqldb

import (
	"context"
	"time"

	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/dbrepo"
	"github.com/eurofurence/reg-mail-service/internal/repository/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlRepository struct {
	db *gorm.DB
}

func Create() dbrepo.Repository {
	return &MysqlRepository{}
}

func (r *MysqlRepository) Open() {
	db, err := gorm.Open("mysql", config.DatabaseMysqlConnectString())
	if err != nil {
		logging.NoCtx().Fatal("failed to open mysql connection: %v", err)
	}

	// see https://making.pusher.com/production-ready-connection-pooling-in-go/
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetConnMaxLifetime(time.Minute * 10)

	r.db = db
}

func (r *MysqlRepository) Close() {
	err := r.db.Close()
	if err != nil {
		logging.NoCtx().Fatal("failed to close mysql connection: ", err)
	}
}

func (r *MysqlRepository) GetTemplates(ctx context.Context) (*entity.Template, error) {
	var a entity.Template
	err := r.db.First(&a).Error
	if err != nil {
		logging.Ctx(ctx).Info("mysql error during selection of templates: ", err)
	}
	return &a, err
}

// Note: The DeleteTemplate function does not really delete a database entry.
// Rather it sets the 'deleted_at' timestamp, which results it not being found on the
// get/update queries.
// This could be useful for some sort of "Archive" in the Dashboard for the Admins to restore
// deleted templates or as some sort of Backup.
func (r *MysqlRepository) DeleteTemplate(ctx context.Context, uuid string) error {
	var a entity.Template
	err := r.db.Delete(&a, "id = ?", uuid).Error
	if err != nil {
		logging.Ctx(ctx).Info("mysql error during template deletion: ", err)
	}
	return err
}

func (r *MysqlRepository) UpdateTemplate(ctx context.Context, uuid string, data string) error {
	var temp *entity.Template

	temp, err := r.GetTemplateById(ctx, uuid)
	if err != nil {
		logging.Ctx(ctx).Info("mysql error during template update: ", err)
		return err
	}

	temp.Data = data
	err = r.db.Save(temp).Error
	if err != nil {
		logging.Ctx(ctx).Info("mysql error during template update: ", err)
	}
	return err
}

func (r *MysqlRepository) GetTemplateById(ctx context.Context, id string) (*entity.Template, error) {
	var a entity.Template
	err := r.db.First(&a, "id = ?", id).Error
	if err != nil {
		logging.Ctx(ctx).Info("mysql error during template select: ", err)
	}
	return &a, err
}

func (r *MysqlRepository) GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error) {
	var a entity.Template
	err := r.db.First(&a, "cid = ? AND lang = ?", cid, lang).Error
	if err != nil {
		logging.Ctx(ctx).Info("mysql error during template select [", cid, "/", lang, "]: ", err, ". Trying en-US fallback...")

		err := r.db.First(&a, "cid = ? AND lang = ?", cid, "en-US").Error
		if err != nil {
			logging.Ctx(ctx).Info("mysql error during template select: ", err)
		}
		return &a, err
	}
	return &a, err
}
