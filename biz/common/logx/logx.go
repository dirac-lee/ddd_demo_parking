// Package logx works as both log and alert.
//
// As a log printer, logx defines 3 log levels: Info, Warn and Error for
// logger collection in Argos.
//
// As an alert trigger, logx will alert when log level is Error. It defines
// 2 alert levels: Notice and Fatal.
package logx

import (
	"fmt"
	"strings"

	"code.byted.org/gopkg/context"
	"code.byted.org/gopkg/logs/v2/log"
	"github.com/dirac-lee/ddd_demo_parking/biz/common/runtimex"
	"github.com/luci/go-render/render"
)

const (
	Skip = 4
)

var (
	logInfo  = log.V2.Info
	logWarn  = log.V2.Warn
	logError = log.V2.Error
)

func Info(ctx context.Context, format string, v ...any) {
	logInfo().With(ctx).CallDepth(1).Str(fmt.Sprintf(format, v...)).Emit()
}

func Warn(ctx context.Context, format string, v ...any) {
	logWarn().With(ctx).CallDepth(1).Str(fmt.Sprintf(format, v...)).Emit()
}

func Error(ctx context.Context, format string, v ...any) {
	logError().With(ctx).CallDepth(1).Str(fmt.Sprintf(format, v...)).Emit()
}

func AlertWarn(ctx context.Context, format string, v ...any) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ALERT WARN: "+format, v...))
	sb.WriteString("\nstacktrace:\n")
	sb.WriteString(runtimex.Stacktrace())
	logError().With(ctx).CallDepth(1).Str(sb.String()).Emit()
}

func AlertError(ctx context.Context, format string, v ...any) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ALERT ERROR: "+format, v...))
	sb.WriteString("\nstacktrace:\n")
	sb.WriteString(runtimex.Stacktrace())
	logError().With(ctx).CallDepth(1).Str(sb.String()).Emit()
}

func InfoInput(ctx context.Context, input map[string]any) {
	logInfo().With(ctx).CallDepth(1).Str(formatCallAndArgs("receive input", input)).Emit()
}

func InfoOutput(ctx context.Context, output map[string]any) {
	logInfo().With(ctx).CallDepth(1).Str(formatCallAndArgs("return output", output)).Emit()
}

func InfoTopic(ctx context.Context, topic string, args map[string]any) {
	logInfo().With(ctx).CallDepth(1).Str(formatCallAndArgs(topic, args)).Emit()
}

func WarnTopic(ctx context.Context, topic string, args map[string]any) {
	logWarn().With(ctx).CallDepth(1).Str(formatCallAndArgs(topic, args)).Emit()
}

func WarnCall(ctx context.Context, callee string, calleeArgs map[string]any, err error) {
	topic := fmt.Sprintf("-> %v failed, \nCall Warn", callee)
	if len(calleeArgs) > 0 {
		topic += fmt.Sprintf("\nCall Args: %v", formatArgs(calleeArgs))
	}
	logWarn().With(ctx).CallDepth(1).Str(formatCallAndArgs(topic, map[string]any{
		"err": err,
	})).Emit()
}

func ErrorCall(ctx context.Context, callee string, calleeArgs map[string]any, err error) {
	topic := fmt.Sprintf("-> %v failed, ", callee)
	if len(calleeArgs) > 0 {
		topic += fmt.Sprintf("\nCall Args: %v", formatArgs(calleeArgs))
	}
	topic += "\nCall Error"
	logError().With(ctx).CallDepth(1).Str(formatCallAndArgs(topic, map[string]any{
		"err": err,
	})).Emit()
}

func formatCallAndArgs(topic string, args map[string]any) string {
	var sb strings.Builder
	sb.WriteString(formatCaller(Skip, topic))
	sb.WriteString(":")
	sb.WriteString(formatArgs(args))
	sb.WriteString("\n")
	return sb.String()
}

func formatCaller(skip int, topic string) string {
	var sb strings.Builder
	method := runtimex.GetCaller(skip)
	sb.WriteString(method)
	sb.WriteString(" ")
	sb.WriteString(topic)
	return sb.String()
}

func formatArgs(args map[string]any) string {
	var sb strings.Builder
	for argName, argValue := range args {
		sb.WriteByte('\n')
		sb.WriteString("\t- ")
		sb.WriteString(argName)
		sb.WriteString(": ")
		sb.WriteString(render.Render(argValue))
	}
	return sb.String()
}
