package template

import "time"

type TemplateDto struct {
	UUID      string    `json:"uuid"`
	CommonID  string    `json:"cid"`
	Lang      string    `json:"lang"`
	Subject   string    `json:"subject"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
