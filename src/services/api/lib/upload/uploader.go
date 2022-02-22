package upload

import (
	"context"
	"io"

	"github.com/zikwall/fsclient/impl"
)

type Uploader interface {
	UploadFile(ctx context.Context, name string, file io.Reader) error
}

type FileUploader struct {
	uploader impl.Client
}

func NewFileUploader(client impl.Client) Uploader {
	return &FileUploader{
		uploader: client,
	}
}

func (f *FileUploader) UploadFile(ctx context.Context, name string, file io.Reader) error {
	return f.uploader.SendFile(ctx, impl.FileDest{
		Name: name,
		File: file,
	})
}
