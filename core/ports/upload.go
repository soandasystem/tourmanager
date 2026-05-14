package ports

import (
	"context"
	"io"
)

type UploadService interface {
	UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error)
}

type UploadStorage interface {
	Upload(ctx context.Context, file io.Reader, filename string, contentType string) (string, error)
}
