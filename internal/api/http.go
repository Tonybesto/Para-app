package api

import (
	"paramount_school/internal/ports"
)

type HTTPHandler struct {
	Repository ports.Repository
	AWS        ports.AWSRepository
	SMS        ports.SmsServiceRepository
}

func NewHTTPHandler(repository ports.Repository, AWS ports.AWSRepository, Sms ports.SmsServiceRepository) *HTTPHandler {
	return &HTTPHandler{
		Repository: repository,
		AWS:        AWS,
		SMS:        Sms,
	}
}
