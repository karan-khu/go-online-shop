package upload

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v5"
)

var allowedExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

func SaveImage(pctx *echo.Context, uploadDir string) (string, error) {
	ct := pctx.Request().Header.Get("Content-Type")
	if !strings.Contains(ct, "multipart/form-data") {
		return "", errors.New("Content-Type must be multipart/form-data")
	}

	file, err := pctx.FormFile("file")
	if err != nil {
		return "", errors.New("no file uploaded")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExt[ext] {
		return "", errors.New("only jpg, png allowed")
	}

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", errors.New("failed to create dir: " + err.Error())
	}

	dst := filepath.Join(uploadDir, file.Filename)

	src, err := file.Open()
	if err != nil {
		return "", errors.New("failed to read file: " + err.Error())
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", errors.New("failed to save: " + err.Error())
	}
	defer out.Close()

	if _, err := io.Copy(out, src); err != nil {
		os.Remove(dst)
		return "", errors.New("failed to save: " + err.Error())
	}

	return "uploads/" + file.Filename, nil
}
