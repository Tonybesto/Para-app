package ports

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"mime/multipart"
	"paramount_school/internal/model"
)

type Repository interface {
	CreateStudent(student *model.Student) error
	FindStudentByEmail(email string) (*model.Student, error)
	FindStudentByPhone(phone string) (*model.Student, error)
	UpdateStudentProfile(email string, student model.Student) error
	TokenInBlacklist(token *string) bool
	AddTokenToBlacklist(email string, token string) error
	FindHighSchools() ([]model.HighSchool, error)
	CreateHighSchool(highSchool *model.HighSchool) error
}

// AWSRepository interface to implement AWS
type AWSRepository interface {
	UploadFileToS3(h *session.Session, file multipart.File, fileName string, size int64) (string, error)
}

// SmsServiceRepository interface to implement AWS
type SmsServiceRepository interface {
	SendSms(msg, to string) (*string, error)
}
