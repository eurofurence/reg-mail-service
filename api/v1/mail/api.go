package mail

type MailSendDto struct {
	CommonID  string            `json:"cid"`
	Lang      string            `json:"lang"`
	To        []string          `json:"to"`
	Cc        []string          `json:"cc"`
	Bcc       []string          `json:"bcc"`
	Variables map[string]string `json:"variables"`
}
