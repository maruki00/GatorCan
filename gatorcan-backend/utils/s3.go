package utils

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/credentials"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/aws/aws-sdk-go-v2/service/s3/types"
// )

// const (
// 	// use .env for security (avoid Sensitive Data Exposure)
// 	// please visit OWASP Top 10 - Sensitive Data Exposure for more information.
// 	AWS_REGION            = "us-east-1"
// 	BUCKET                = "user-assignment"
// 	ENDPOINT              = "http://172.29.0.2:9000" // Include scheme
// 	AWS_ACCESS_KEY_ID     = "12345"
// 	AWS_SECRET_ACCESS_KEY = "12345"
// )

// type S3 struct {
// 	Client *s3.Client
// 	Bucket string
// }

// func NewS3() (*S3, error) {
// 	// Load credentials from environment variables
// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion(AWS_REGION),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
// 			AWS_ACCESS_KEY_ID,
// 			AWS_SECRET_ACCESS_KEY,
// 			"",
// 		)),
// 		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
// 			if service == s3.ServiceID {
// 				return aws.Endpoint{
// 					URL:           ENDPOINT,
// 					SigningRegion: AWS_REGION,
// 				}, nil
// 			}
// 			return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
// 		})),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load AWS config: %w", err)
// 	}

// 	return &S3{
// 		Client: s3.NewFromConfig(cfg),
// 		Bucket: BUCKET,
// 	}, nil
// }

// func (u *S3) UploadFile(ctx context.Context, key, filePath, contentType string, public bool) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return fmt.Errorf("failed to open file %q: %w", filePath, err)
// 	}
// 	defer file.Close()

// 	input := &s3.PutObjectInput{
// 		Bucket:      aws.String(u.Bucket),
// 		Key:         aws.String(key),
// 		Body:        file,
// 		ContentType: aws.String(contentType),
// 	}

// 	if public {
// 		input.ACL = types.ObjectCannedACLPublicRead
// 	}

// 	_, err = u.Client.PutObject(ctx, input)
// 	if err != nil {
// 		return fmt.Errorf("failed to upload file to S3: %w", err)
// 	}

// 	return nil
// }

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const (
	// use .env for security (avoid Sensitive Data Exposure)
	// please visit OWASP Top 10 - Sensitive Data Exposure for more information.
	AWS_REGION = "us-east-1"
	BUCKET     = "user-assignment"
	ENDPOINT   = "http://localhost:9000"
	// ENDPOINT = "http://172.29.0.2:9000"
	S3_USER = "minioadmin"
	S3_PASS = "minioadmin"
)

type S3 struct {
	Client *s3.Client
	Bucket string
}

func NewS3() (*S3, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(AWS_REGION),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			S3_USER,
			S3_PASS,
			"",
		)),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == s3.ServiceID {
				return aws.Endpoint{
					URL:           ENDPOINT,
					SigningRegion: AWS_REGION,
				}, nil
			}
			return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &S3{
		Client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		}),
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
