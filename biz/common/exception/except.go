package E

import (
	"fmt"
	"runtime/debug"
	"strings"

	"code.byted.org/lang/gg/gptr"
	"code.byted.org/oec/overpass_common_struct/kitex_gen/base"
	"code.byted.org/oec/status_code/Status"
	"github.com/pkg/errors"
)

type Exception struct {
	code       int32
	message    string
	details    []string
	cause      error
	stacktrace string
}

func New(status *Status.Status, opts ...ExceptionOption) *Exception {
	e := &Exception{
		code:    status.Code(),
		message: status.Msg(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *Exception) Error() string {
	if !e.IsException() {
		return "<nil>"
	}
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("\t[Code] %d\n", e.code))
	sb.WriteString(fmt.Sprintf("\t[Message] %s\n", e.message))
	if len(e.details) > 0 {
		sb.WriteString(fmt.Sprintf("\t[Detail] %v\n", strings.Join(e.details, " | ")))
	}
	if e.cause != nil {
		sb.WriteString(fmt.Sprintf("\t[Cause] %s\n", e.cause))
	}
	if len(e.stacktrace) > 0 {
		sb.WriteString(fmt.Sprintf("\t[Stacktrace]\n%s", string(e.stacktrace)))
	}
	return sb.String()
}

func (e *Exception) IsException() bool {
	return e != nil
}

func (e *Exception) AppendDetail(detailFormat string, args ...any) {
	e.details = append(e.details, fmt.Sprintf(detailFormat, args...))
}

func (e *Exception) BindStack(skip int) {
	if len(e.stacktrace) > 0 { // 不覆盖已有 stacktrace
		return
	}
	e.stacktrace = trimStacktrace(debug.Stack(), skip)
}

func (e *Exception) Sprintf(format string, args ...any) string {
	format += ", err: %+v"
	args = append(args, e)
	return fmt.Sprintf(format, args...)
}

func (e *Exception) Code() int32 {
	if e == nil {
		return Status.Success.Code()
	}
	return e.code
}

func (e *Exception) Message() string {
	if e == nil {
		return Status.Success.Msg()
	}
	return e.message
}

func (e *Exception) CodePtr() *int32 {
	return gptr.Of(e.Code())
}

func (e *Exception) MessagePtr() *string {
	return gptr.Of(e.Message())
}

func (e *Exception) ToBaseResp() *base.BaseResp {
	return &base.BaseResp{
		StatusMessage: e.Message(),
		StatusCode:    e.Code(),
	}
}

type ExceptionOption func(*Exception)

func WithCause(cause error) ExceptionOption {
	return func(e *Exception) {
		e.cause = cause
	}
}

func WithDetail(detailFormat string, args ...any) ExceptionOption {
	return func(e *Exception) {
		e.details = append(e.details, fmt.Sprintf(detailFormat, args...))
	}
}

func WithStack(skip int) ExceptionOption {
	return func(e *Exception) {
		e.stacktrace = trimStacktrace(debug.Stack(), skip)
	}
}

// trimStacktrace removes the top 'skip' frames from the stacktrace trace.
func trimStacktrace(stackTrace []byte, skip int) string {
	lines := strings.Split(string(stackTrace), "\n")
	if len(lines) > 2*skip {
		lines = lines[2*skip:] // 每个栈帧由两行表示
	}
	return strings.Join(lines, "\n")
}

func Throw(e *Exception) {
	if e == nil {
		return
	}
	panic(e)
}

func Try(f func()) (b *Box) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		if e, ok := r.(*Exception); ok {
			if e == nil {
				return
			}
			e.BindStack(6)
			b.e = e
		} else {
			b.e = New(Status.SystemError, WithCause(errors.Errorf("panicked: %v", r)), WithStack(6))
		}
	}()
	b = new(Box)
	f()
	return b
}

type Box struct {
	e *Exception
}

func (b *Box) Catch(f func(e *Exception)) {
	if b.e.IsException() {
		f(b.e)
	}
}
