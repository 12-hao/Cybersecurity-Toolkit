package logger

import (
	"log/slog"
)

type Option func(*PrettyHandlerOptions) //实现一个接口 自定义一个这样的函数让下面的函函数返回值是这个函数

func WithLevel(level slog.Level) Option {
	return func(o *PrettyHandlerOptions) {
		o.SlogOpts.Level = level
	}
}

func WithTimeFormat(format string) Option {
	return func(o *PrettyHandlerOptions) {
		o.TimeFormat = format
	}
}

func WithUserColor(useColor bool) Option {
	return func(o *PrettyHandlerOptions) {
		o.UserColor = useColor
	}

}

func WithOutPutJson(outPutJson bool) Option {
	return func(o *PrettyHandlerOptions) {
		o.OutPutJson = outPutJson
	}
}
