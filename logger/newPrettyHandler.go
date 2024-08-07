package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/1755616537/utils"
	"github.com/gogf/gf/encoding/gjson"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"
)

// ANSI转义码定义
const (
	//洋红色
	Magenta = "\u001B[35m"
	//蓝色
	Blue = "\u001B[34m"
	//黄色
	Yellow = "\u001B[33m"
	//红色
	Red = "\u001B[31m"
	//青色
	Cyan = "\u001B[36m"
	//白色
	White = "\u001B[37m"

	//结束颜色码
	ResetAll = "\u001B[0m"
)

// 字符串转颜色
func StringToColor(str, color string) string {
	return fmt.Sprint(color, str, ResetAll)
}

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = StringToColor(level, Magenta)
	case slog.LevelInfo:
		level = StringToColor(level, Blue)
	case slog.LevelWarn:
		level = StringToColor(level, Yellow)
	case slog.LevelError:
		level = StringToColor(level, Red)
	}

	var errfields *fmtErrtype
	var errfieldsErrError string
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		switch a.Value.Kind() {
		case slog.KindAny:
			switch v := a.Value.Any().(type) {
			case error:
				if printErrorStack {
					a.Value = fmtErr(v)
					errfields = fmtErr2(v)
				}
				errfieldsErrError = v.Error()
			}
		}

		fields[a.Key] = a.Value.Any()

		return true
	})

	if printErrorStack && errfields != nil {
		//type value struct {
		//	Error fmtErrtype `json:"error"`
		//}
		//var kk value
		//err := json.Unmarshal(gjson.New(errfields).MustToJson(), &kk)
		//if err != nil {
		//	return err
		//}

		if errfields.Trace != nil {
			fields["error"] = errfields
		} else {
			fields["error"] = errfieldsErrError
		}
	} else if errfieldsErrError != "" {
		fields["error"] = errfieldsErrError
	}

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[2006/01/02 15:04:05.000000]")
	msg := StringToColor(r.Message, Cyan)
	message := make(map[string]interface{})
	message["data"] = r.Message
	message["fields"] = fields

	if onLogFile {
		_ = h.Handler.Handle(ctx, slog.NewRecord(r.Time, r.Level, gjson.New(message).MustToJsonString(), r.PC))
		//h.l.Println(timeStr, r.Level.String()+":", r.Message, string(b))
	} else {
		//_ = h.Handler.Handle(ctx, slog.Record{
		//	Level: r.Level,
		//	PC:    r.PC,
		//})
	}

	var errorStackB string
	if printStack {
		stack := make(map[string]interface{})
		stack[StringToColor(slog.SourceKey, Yellow)] = utils.GetErrorStack(r.PC)
		StackByte, err := json.MarshalIndent(stack, "", "  ")
		if err != nil {
			return err
		}
		errorStackB = strings.Replace(
			string(StackByte),
			"\"\\u001b[33msource\\u001b[0m\": {",
			fmt.Sprint(
				StringToColor(slog.SourceKey, Yellow),
				": {",
			),
			1,
		)
	}
	fmt.Println(timeStr, level, msg, string(b))
	if printStack {
		fmt.Println(errorStackB)
	}

	return nil
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

// CustomHandler 是一个自定义的日志处理器，实现了 slog.Handler 接口。
type CustomHandler struct {
	slog.Handler
}

// HandleLog 实现了 slog.Handler 接口的 Handle 方法。
func (h *CustomHandler) HandleLog(r *slog.Record) error {
	// 获取当前时间
	timestamp := time.Now().Format(time.RFC3339)

	// 获取日志级别
	level := _recordToLevel(r)

	// 获取日志消息
	msg := _msg(r)

	// 构建自定义格式的日志输出
	logOutput := fmt.Sprintf("%s | %s | %s\n", timestamp, level, msg)

	// 输出到标准输出
	_, err := os.Stdout.Write([]byte(logOutput))
	if err != nil {
		return err
	}

	return nil
}

// _recordToLevel 是一个辅助函数，用于从 Record 中提取日志级别。
func _recordToLevel(r *slog.Record) string {
	switch r.Level {
	case slog.LevelDebug:
		return "DEBUG"
	case slog.LevelInfo:
		return "INFO"
	case slog.LevelWarn:
		return "WARN"
	case slog.LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// _msg 是一个辅助函数，用于从 Record 中提取日志消息。
func _msg(r *slog.Record) string {
	var msg string
	r.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "msg" {
			msg = attr.Value.Any().(string)
			return true
		}
		return false
	})
	return msg
}
