package dto

import "time"

type TemplateDto struct {
	ID        string    `json:"id"`
	CommonID  string    `json:"cid"`
	Lang      string    `json:"lang"`
	Title     string    `json:"title"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
