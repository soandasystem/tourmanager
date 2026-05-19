package services

import (
	"context"
	"io"

	"tourmanager/config"
	"tourmanager/core/ports"
)

type uploadService struct {
	config  config.Config
	storage ports.UploadStorage
}

func NewUploadService(cfg config.Config, storage ports.UploadStorage) ports.UploadService {
	return &uploadService{
		config:  cfg,
		storage: storage,
	}
}

func (s *uploadService) UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error) {
	objectKey := "uploads/" + filename

	return s.storage.Upload(ctx, file, objectKey, contentType)
}
