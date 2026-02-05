package services

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type LocalStorage struct {
	baseDir string
}

func NewLocalStorage(baseDir string) *LocalStorage {
	if baseDir == "" {
		baseDir = "uploads"
	}
	return &LocalStorage{baseDir: baseDir}
}

func (s *LocalStorage) Save(file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll(s.baseDir, 0o755); err != nil {
		return "", err
	}

	randomSuffix := randomString(8)
	ext := filepath.Ext(file.Filename)
	name := strings.TrimSuffix(file.Filename, ext)
	filename := fmt.Sprintf("%s_%s%s", name, randomSuffix, ext)
	path := filepath.Join(s.baseDir, filename)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return path, nil
}

func randomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "rand"
	}
	return fmt.Sprintf("%x", b)[:n]
}
