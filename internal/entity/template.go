package entity

type Template struct {
	Base
	CommonID string `gorm:"type:varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Language string `gorm:"type:enum('en-US','de-DE') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Title    string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Data     string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
}
