package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"log/slog"
	"os"
)

var Log *Extend

type Extend struct {
	*slog.Logger
	handler *PrettyHandler
}
type PrettyHandlerOptions struct {
	SlogOpts   slog.HandlerOptions
	TimeFormat string
	UserColor  bool
	OutPutJson bool
} //日志配置结构体

type PrettyHandler struct {
	slog.Handler
	l   *log.Logger
	opt *PrettyHandlerOptions
} //这里写的是具体的调用

func (l *Extend) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}
func (l *Extend) DebugMsgf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Debug(msg)
}

func (l *Extend) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Extend) InfoMsgf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Info(msg)
}

func (l *Extend) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *Extend) WarnMsgf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Warn(msg)
}
func (l *Extend) Error(msg string, args ...any) error {
	l.Logger.Error(msg, args...)
	return fmt.Errorf(msg)
}

func (l *Extend) ErrorMsgf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Error(msg)
	return fmt.Errorf(msg)
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {

	TimeStr := r.Time.Format(h.opt.TimeFormat)
	level := "[" + r.Level.String() + "]:"

	if h.opt.UserColor {
		switch r.Level {
		case slog.LevelDebug:
			level = color.MagentaString(level)
		case slog.LevelInfo:
			level = color.BlueString(level)
		case slog.LevelWarn:
			level = color.YellowString(level)
		case slog.LevelError:
			level = color.RedString(level)
		}
	}
	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})
	if len(fields) > 0 && h.opt.OutPutJson {
		b, err := json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
		h.l.Printf("%s %s %s %s\n", TimeStr, level, r.Message, string(b))
	} else {
		h.l.Printf("%s %s %s \n", TimeStr, level, r.Message)
	}
	return nil
}
func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
		opt:     &opts,
	}
	return h
}

func Init(option ...Option) {
	opts := PrettyHandlerOptions{
		SlogOpts:   slog.HandlerOptions{Level: slog.LevelDebug},
		TimeFormat: "2006-01-02 15:04:05",
		UserColor:  true,
		OutPutJson: true,
	}
	for _, opt := range option {
		opt(&opts)
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	Log = &Extend{
		Logger:  slog.New(handler),
		handler: handler,
	}
}
