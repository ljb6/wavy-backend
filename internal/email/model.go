package email

type EmailReq struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
