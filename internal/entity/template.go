package entity

type Template struct {
	Base
	CommonID string `gorm:"column:cid; CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;index:unique"`
	Language string `gorm:"column:lang; type:varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;index:unique"`
	Subject  string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Data     string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
}
