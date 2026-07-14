package b2

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"tourmanager/core/ports"
)


type b2Storage struct {
	client    *s3.Client
	bucket    string
	endpoint  string
	publicURL string // URL pública de descarga, ej: https://f005.backblazeb2.com/file/<bucket>
}

func NewB2Storage(ctx context.Context, keyID, applicationKey, bucket, region, endpoint, publicURL string) ports.UploadStorage {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(keyID, applicationKey, "")),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("no se pudo cargar la config de b2: %v", err)
	}

	// La forma correcta y moderna de forzar el endpoint en el SDK v2 para S3 (Backblaze B2)
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})

	return &b2Storage{
		client:    client,
		bucket:    bucket,
		endpoint:  endpoint,
		publicURL: strings.TrimRight(publicURL, "/"),
	}
}

func (s *b2Storage) Upload(ctx context.Context, file io.Reader, objectKey string, contentType string) (string, error) {
	// B2 no soporta chunked transfer encoding: se debe enviar Content-Length explícito.
	// Leemos todo en memoria para conocer el tamaño antes de llamar a PutObject.
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error leyendo contenido para subir: %w", err)
	}

	contentLength := int64(len(data))

	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(objectKey),
		Body:          bytes.NewReader(data),
		ContentType:   aws.String(contentType),
		ContentLength: &contentLength,
	})
	if err != nil {
		return "", fmt.Errorf("error al subir a B2: %w", err)
	}

	publicURL := fmt.Sprintf("%s/%s", s.publicURL, objectKey)
	return publicURL, nil
}

