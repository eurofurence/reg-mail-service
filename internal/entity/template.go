package entity

type Template struct {
	Base
	CommonID string `gorm:"column:cid; type:varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;uniqueIndex:template"`
	Language string `gorm:"column:lang; type:varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;uniqueIndex:template"`
	Subject  string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Data     string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
}

type Attachment struct {
	Base
	Name    string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Content string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
}

type TemplateV2 struct {
	Base
	CommonID    string `gorm:"column:cid; type:varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;uniqueIndex:template"`
	Language    string `gorm:"column:lang; type:varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL;uniqueIndex:template"`
	Subject     string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Text        string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	HTML        string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Attachments []Attachment
	Embedded    []Attachment
}
