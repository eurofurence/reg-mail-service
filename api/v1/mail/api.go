package mail

type TemplateRequestDto struct {
	Name      string `json:"name"`
	Variables string `json:"variables"`
	Nick      string `json:"nick"`
}
