package entity

import "gorm.io/gorm"

type Failure struct {
	gorm.Model
	CommonID string `gorm:"column:cid; type:varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Language string `gorm:"column:lang; type:varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Request  string `gorm:"column:request; type:longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci; NOT NULL"`
	Error    string `gorm:"column:error; type:longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci; NOT NULL"`
}
