package s3

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func CreateSession() *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(viper.GetString("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(
				viper.GetString("AWS_ACCESS"),
				viper.GetString("AWS_SECRET"),
				"",
			),
		},
	))
	return sess
}

func UploadImageToS3(image []byte, sess *session.Session) (string, error) {
	uploader := s3manager.NewUploader(sess)
	key := uuid.New().String() // Generating a unique key for the image
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("gigforge/profile-images/"),
		Key:    aws.String(key),
		Body:   bytes.NewReader(image),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Println("error", err)
		return "", err
	}
	return upload.Location, nil
}
