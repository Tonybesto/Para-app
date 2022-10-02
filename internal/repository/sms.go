package repository

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
	"paramount_school/internal/ports"
)

type SmsService struct {
}

func NewSmsService() ports.SmsServiceRepository {
	return &SmsService{}
}

var client *twilio.RestClient

func (s *SmsService) SendSms(msg, to string) (*string, error) {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	from := os.Getenv("FROM_PHONE")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	params := openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(msg)

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		return nil, err
	}
	return response.Sid, nil
}
