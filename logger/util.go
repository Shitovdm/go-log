package logger

import (
	"github.com/gofrs/uuid"
	"io"
)

type noopCloser struct {
	io.Writer
}

func (n noopCloser) Close() error {
	return nil
}

func GenerateNewSessionUuid() string {
	u4, _ := uuid.NewV4()
	return u4.String()
}
