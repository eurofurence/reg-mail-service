package entity

type Template struct {
	Base
	CommonID string `gorm:"column:cid; type:varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;index:template"`
	Language string `gorm:"column:lang; type:varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;index:template"`
	Subject  string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Data     string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
}
