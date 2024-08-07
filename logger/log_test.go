package logger

import (
	"context"
	"errors"
	"github.com/mdobak/go-xerrors"
	"log/slog"
	"net/http"
	"os"
	"testing"
)

func Test_log(t *testing.T) {
	err := IniLog()
	if err != nil {
		return
	}

	slog.Debug(
		"executing database query",
		slog.String("query", "SELECT * FROM users"),
	)
	slog.Info("image upload successful", slog.String("image_id", "39ud88"))
	slog.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 MB"),
	)
	err = errors.New("something happened")
	slog.Error(
		"An error occurred while processing the request",
		slog.Any("error", err),
	)

	ctx := context.Background()
	err = xerrors.New("something happened")
	slog.ErrorContext(ctx, "image uploaded", slog.Any("error", err))

	err = errors.New("something happened2")
	slog.ErrorContext(ctx, "2", slog.Any("error", err))

	err = errors.New("something happened3")
	slog.ErrorContext(ctx, "3", slog.Any("error", xerrors.New(err)))

	err = errors.New("something happened4")
	slog.ErrorContext(ctx, "upload failed", slog.Any("error", err.Error()))

	slog.Error("error2", err)
}

// slog.Logger 转换为 log.Logger
func Test_Tolog(t *testing.T) {
	handler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.NewLogLogger(handler, slog.LevelError)

	_ = http.Server{
		// this API only accepts `log.Logger`
		ErrorLog: logger,
	}
}

// 使用 LogValuer 接口隐藏敏感字段
func Test_logValue(t *testing.T) {
	type LogValuer interface {
		LogValue() slog.Value
	}

	//{
	//	"time": "2023-03-15T14:44:24.223381036+01:00",
	//	"level": "INFO",
	//	"msg": "info",
	//	"user": {
	//	"id": "user-12234",
	//		"name": "Jan Doe"
	//}
	//}
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//	func (u *User) LogValue() slog.Value {
//		return slog.StringValue(u.ID)
//	}
func (u *User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", u.ID),
		slog.String("name", u.FirstName+" "+u.LastName),
	)
}
