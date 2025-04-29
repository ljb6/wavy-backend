package email

type EmailHandler struct {
	emailService *EmailService
}

func NewHandler(service *EmailService) *EmailHandler {
	return &EmailHandler{emailService: service}
}