package json

import (
	"context"
)

func UnmarshalStrErrorWriteLog[T any](ctx context.Context, data string) T {
	t, err := Unmarshal[T]([]byte(data))
	if err != nil {
		//log.w(ctx, "UnmarshalString error, data=%s, err=%s", data, err)
	}
	return t
}

func MarshalStrErrorWriteLog(ctx context.Context, data any) string {
	t, err := MarshalString(data)
	if err != nil {
		//logger.Warnf(ctx, "MarshalString error, data=%v, err=%s", data, err)
	}
	return t
}

func UnmarshalStrErrorPanic[T any](ctx context.Context, data string) T {
	t, err := Unmarshal[T]([]byte(data))
	if err != nil {
		//logger.Warnf(ctx, "Unmarshal error, data=%v, err=%s", data, err)
		panic("json unmarshal error")
	}
	return t
}

func MarshalStrErrorPanic(ctx context.Context, data any) string {
	t, err := MarshalString(data)
	if err != nil {
		//logger.Warnf(ctx, "MarshalString error, data=%v, err=%s", data, err)
		panic("json marshal error")
	}
	return t
}
