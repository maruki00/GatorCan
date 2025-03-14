package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const (
	// use .env for security (avoid Sensitive Data Exposure)
	// please visit OWASP Top 10 - Sensitive Data Exposure for more information.
	AWS_ACCESS_KEY_ID     = "12345"
	AWS_SECRET_ACCESS_KEY = "12345"
	AWS_REGION            = "us-east-1"
	BUCKET                = "user-assignment"
)

type S3 struct {
	Client *s3.Client
	Bucket string
}

func NewS3() (*S3, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(AWS_REGION),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     "12345",
				SecretAccessKey: "12345",
			}, nil
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &S3{
		Client: s3.NewFromConfig(cfg),
		Bucket: BUCKET,
	}, nil
}

func (u *S3) UploadFile(ctx context.Context, key, filePath, contentType string, public bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", filePath, err)
	}
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket:      aws.String(u.Bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	}

	if public {
		input.ACL = types.ObjectCannedACLPublicRead
	}

	_, err = u.Client.PutObject(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return nil
}
