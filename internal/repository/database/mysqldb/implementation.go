package mysqldb

import (
	"context"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"gorm.io/gorm/schema"
	"net/url"
	"time"

	"github.com/eurofurence/reg-mail-service/internal/entity"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/repository/database/dbrepo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func Create() dbrepo.Repository {
	return &MysqlRepository{}
}

func (r *MysqlRepository) Open() error {
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mail_",
		},
	}
	connectString := config.DatabaseMysqlConnectString()

	db, err := gorm.Open(mysql.Open(connectString), &gormConfig)
	if err != nil {
		aulogging.Logger.NoCtx().Error().WithErr(err).Printf("failed to open mysql connection: %s", err.Error())
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		aulogging.Logger.NoCtx().Error().WithErr(err).Printf("failed to configure mysql connection: %s", err.Error())
		return err
	}

	// see https://making.pusher.com/production-ready-connection-pooling-in-go/
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(50)
	sqlDb.SetConnMaxLifetime(time.Minute * 10)

	r.db = db
	return nil
}

func (r *MysqlRepository) Close() {
	// no more db close in gorm v2
}

func (r *MysqlRepository) Migrate() error {
	err := r.db.AutoMigrate(
		&entity.Template{},
		&entity.History{},
	)
	if err != nil {
		aulogging.Logger.NoCtx().Error().WithErr(err).Printf("failed to migrate mysql db: %s", err.Error())
		return err
	}
	return nil
}

func (r *MysqlRepository) GetTemplates(ctx context.Context) ([]*entity.Template, error) {
	result := make([]*entity.Template, 0)
	buffer := entity.Template{}

	rows, err := r.db.Model(&buffer).Order("cid, lang").Rows()
	if err != nil {
		aulogging.Logger.Ctx(ctx).Error().WithErr(err).Printf("mysql error during selection of templates: %s", err.Error())
		return result, err
	}
	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			aulogging.Logger.Ctx(ctx).Warn().WithErr(err2).Printf("secondary error closing recordset during find: %s", err2.Error())
		}
	}()

	for rows.Next() {
		err = r.db.ScanRows(rows, &buffer)
		if err != nil {
			aulogging.Logger.Ctx(ctx).Error().WithErr(err).Printf("error reading template during find: %s", err.Error())
			return result, err
		}
		copied := buffer
		result = append(result, &copied)
	}

	return result, nil
}

func (r *MysqlRepository) CreateTemplate(ctx context.Context, tpl *entity.Template) error {
	err := r.db.Create(&tpl).Error
	if err != nil {
		aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("mysql error during template creation: %s", err.Error())
	}
	return err
}

// Note: The DeleteTemplate function does not really delete a database entry.
// Rather it sets the 'deleted_at' timestamp, which results it not being found on the
// get/update queries.
// This could be useful for some sort of "Archive" in the Dashboard for the Admins to restore
// deleted templates or as some sort of Backup.
func (r *MysqlRepository) DeleteTemplate(ctx context.Context, uuid string, permanent bool) error {
	var a entity.Template
	err := r.db.Error

	if permanent {
		err = r.db.Unscoped().Delete(&a, "id = ?", uuid).Error
	} else {
		err = r.db.Delete(&a, "id = ?", uuid).Error
	}

	if err != nil {
		aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("mysql error during template deletion: %s", err.Error())
	}
	return err
}

func (r *MysqlRepository) UpdateTemplate(ctx context.Context, uuid string, data *entity.Template) error {
	var temp *entity.Template

	temp, err := r.GetTemplateById(ctx, uuid)
	if err != nil {
		aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("mysql error during template update: %s", err.Error())
		return err
	}

	temp.Data = data.Data
	temp.Subject = data.Subject
	temp.CommonID = data.CommonID
	temp.Language = data.Language

	err = r.db.Save(temp).Error
	if err != nil {
		aulogging.Logger.Ctx(ctx).Warn().WithErr(err).Printf("mysql error during template update: %s", err.Error())
	}
	return err
}

func (r *MysqlRepository) GetTemplateById(ctx context.Context, id string) (*entity.Template, error) {
	var a entity.Template
	err := r.db.First(&a, "id = ?", id).Error
	if err != nil {
		aulogging.Logger.Ctx(ctx).Info().WithErr(err).Printf("mysql error during template select by id '%s': %s", url.QueryEscape(id), err.Error())
	}
	return &a, err
}

func (r *MysqlRepository) GetTemplateByCid(ctx context.Context, cid string, lang string) (*entity.Template, error) {
	var a entity.Template
	err := r.db.First(&a, "cid = ? AND lang = ?", cid, lang).Error
	if err != nil {
		aulogging.Logger.Ctx(ctx).Info().WithErr(err).Printf("mysql error during template select [%s/%s]: %s. Trying en-US fallback...", url.QueryEscape(cid), url.QueryEscape(lang), err.Error())

		err := r.db.First(&a, "cid = ? AND lang = ?", cid, "en-US").Error
		if err != nil {
			aulogging.Logger.Ctx(ctx).Info().WithErr(err).Printf("mysql error during template select: %s", err.Error())
		}
		return &a, err
	}
	return &a, err
}
