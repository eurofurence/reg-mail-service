package template

import "time"

type TemplateDtoV2 struct {
	UUID        string    `json:"uuid"`
	CommonID    string    `json:"cid"`
	Lang        string    `json:"lang"`
	Subject     string    `json:"subject"`
	Text        string    `json:"text"`
	HTML        string    `json:"html"`
	Attachments []string  `json:"attachments"`
	Embedded    []string  `json:"embedded"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TemplateListDtoV2 struct {
	Templates []TemplateDtoV2 `json:"templates"`
}
