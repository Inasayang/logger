package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

func zapFileWriter(file string) io.Writer {
	hook, err := rotatelogs.New(
		file+".%Y%m%d",
		rotatelogs.WithLinkName(file),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
