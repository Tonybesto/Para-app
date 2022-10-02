package repository

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"net/http"
	"os"
	"paramount_school/internal/ports"
)

type AWS struct {
}

func NewAWS() ports.AWSRepository {
	return &AWS{}
}

func (a *AWS) UploadFileToS3(h *session.Session, file multipart.File, fileName string, size int64) (string, error) {
	// get the file size and read the file content into a buffer
	buffer := make([]byte, size)
	_, err2 := file.Read(buffer)
	if err2 != nil {
		return "", err2
	}
	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file you're uploading
	url := "https://s3-us-east-2.amazonaws.com/lunch-wallet/" + fileName
	_, err := s3.New(h).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:                  aws.String(fileName),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	return url, err
}
