package common

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type FileUploader interface {
	Upload(file *multipart.FileHeader) (string, error)
}

type LocalUploader struct {
	Folder string // folder lưu file, ví dụ: "frontend/public"
}

func NewLocalUploader(folder string) *LocalUploader {
	return &LocalUploader{Folder: folder}
}
func (u *LocalUploader) Upload(file *multipart.FileHeader) (string, error) {
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	savePath := u.Folder + "/" + fileName

	// Tạo folder nếu chưa có
	if err := os.MkdirAll(u.Folder, 0755); err != nil {
		return "", err
	}

	// Nếu file đã tồn tại, trả về URL mà không lưu lại
	if _, err := os.Stat(savePath); err == nil {
		// file tồn tại
		return "/" + fileName, nil
	}

	// Lưu file mới
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return "/" + fileName, nil
}
