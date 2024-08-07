package logger

import (
	"fmt"
	"github.com/mdobak/go-xerrors"
	"log/slog"
	"path/filepath"
)

type stackFrame struct {
	Func   string `json:"func"`
	Source string `json:"source"`
	Line   int    `json:"line"`
}

type fmtErrtype struct {
	Msg   string       `json:"msg"`
	Trace []stackFrame `json:"trace"`
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	fmt.Println(a)
	switch a.Value.Kind() {
	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			fmt.Println("sadasd")
			a.Value = fmtErr(v)
		}
	}

	return a
}

// marshalStack 从错误中提取堆栈帧
func marshalStack(err error) []stackFrame {
	trace := xerrors.StackTrace(err)

	if len(trace) == 0 {
		return nil
	}

	frames := trace.Frames()

	s := make([]stackFrame, len(frames))

	for i, v := range frames {
		f := stackFrame{
			Source: filepath.Join(
				filepath.Base(filepath.Dir(v.File)),
				filepath.Base(v.File),
			),
			Func: filepath.Base(v.Function),
			Line: v.Line,
		}

		s[i] = f
	}

	return s
}

// fmtErr 返回一个 slog.Value，其中包含键 `msg` 和 `trace`。如果错误没有实现
// interface { StackTrace() errors.StackTrace }，则省略 `trace` 键。
func fmtErr(err error) slog.Value {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("msg", err.Error()))

	frames := marshalStack(err)

	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}

	return slog.GroupValue(groupValues...)
}

func fmtErr2(err error) *fmtErrtype {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("msg", err.Error()))

	frames := marshalStack(err)

	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}

	return &fmtErrtype{
		err.Error(),
		frames,
	}
}
