package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	MAX_FILE_SIZE = 10 * 1024 * 1024 // 10MB
)

var allowedMimeTypes = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"application/pdf": true,
}

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".pdf":  true,
}

func ValidateFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	if header.Size > MAX_FILE_SIZE {
		return "", errors.New("file size exceeds the 5MB limit")
	}

	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "", errors.New("failed to read file for MIME type detection")
	}

	contentType := http.DetectContentType(buffer)
	if !allowedMimeTypes[contentType] {
		return "", fmt.Errorf("invalid file type: %s", contentType)
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", errors.New("failed to reset file pointer")
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("invalid file extension: %s", ext)
	}
	dst := path.Join(os.TempDir(), header.Filename)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return dst, err
}
