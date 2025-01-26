package awss3

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	service *s3.Client
}

func New() (*s3Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to load configuration: %s", err))
	}
	svc := s3.NewFromConfig(cfg)

	return &s3Client{svc}, nil
}

func (clt *s3Client) GeneratePreSignedUrl(bucket, key string) (string, error) {
	presignClient := s3.NewPresignClient(clt.service)
	presignParams := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	presignDuration := 5 * time.Minute
	presignedURL, err := presignClient.PresignPutObject(context.TODO(), presignParams, s3.WithPresignExpires(presignDuration))
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate pre-signed url: %s", err))
	}

	return presignedURL.URL, nil
}
