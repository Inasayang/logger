package logger

import (
	"go.uber.org/zap/zapcore"
	"time"
)

func jsonEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:    "m",
		LevelKey:      "l",
		TimeKey:       "t",
		NameKey:       "n",
		CallerKey:     "c",
		FunctionKey:   "f",
		StacktraceKey: "st",
		//EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02 15:04:05.999999999Z07:00"))
		},
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	})
}
