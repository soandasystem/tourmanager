package b2

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"tourmanager/core/ports"
)

type b2Storage struct {
	client   *s3.Client
	bucket   string
	endpoint string
}

func NewB2Storage(ctx context.Context, keyID, applicationKey, bucket, region, endpoint string) ports.UploadStorage {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: endpoint,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(keyID, applicationKey, "")),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("no se pudo cargar la config de b2: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &b2Storage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (s *b2Storage) Upload(ctx context.Context, file io.Reader, objectKey string, contentType string) (string, error) {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", fmt.Errorf("error al subir a B2: %w", err)
	}

	publicURL := fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucket, objectKey)
	return publicURL, nil
}
