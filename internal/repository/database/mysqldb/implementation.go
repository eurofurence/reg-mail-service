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
		logging.NoCtx().Fatal("failed to close mysql connection: %v", err)
	}
}

func (r *MysqlRepository) GetTemplateById(ctx context.Context, id string) (*entity.Template, error) {
	var a entity.Template
	err := r.db.First(&a, id).Error
	if err != nil {
		logging.Ctx(ctx).Info("mysql error during attendee select - might be ok: %v", err)
	}
	return &a, err
}
