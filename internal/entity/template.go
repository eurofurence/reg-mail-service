package entity

import "github.com/jinzhu/gorm"

type Template struct {
	gorm.Model
	CommonID string `gorm:"type:varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Language string `gorm:"type:enum('en-US','de-DE') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Title    string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Data     string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
}
