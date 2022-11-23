package utility

import (
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Sendgrid struct {
	Client *sendgrid.Client
}

func NewSendGrid(apiKey string) *Sendgrid {
	client := sendgrid.NewSendClient(apiKey)
	return &Sendgrid{
		Client: client,
	}
}

func (s *Sendgrid) SendMail(from, to *mail.Email, subject, plainTextContent, htmlContent string) error {
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	_, err := s.Client.Send(message)
	if err != nil {
		return err
	}

	return nil
}
