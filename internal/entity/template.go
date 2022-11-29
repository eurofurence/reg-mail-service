package entity

type Template struct {
	Base
	CommonID string `gorm:"column:cid; CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`                                                                                             // unique index on column pair cid + lang for quick search
	Language string `gorm:"column:lang; type:enum('en-US','en-GB','de-DE','de-CH','es-ES','fr-FR','it-IT','nl-NL','pl-PL','ru-RU') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"` // just use string, or adding extra languages becomes a hassle
	Subject  string `gorm:"type:varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
	Data     string `gorm:"type:text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;NOT NULL"`
}
