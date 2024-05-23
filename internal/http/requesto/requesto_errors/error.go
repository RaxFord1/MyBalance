package requesto_errors

import (
	"context"
	"fmt"
	"log"
	"runtime"
)

const DefaultErrorText = `internal error`

type IError interface {
	HTTPCode() int
	GetErrCode() string
	Text() string
	StackTrace() string
	Equal(error) bool
	Error() string
}

type Error struct {
	HttpCode int    `json:"-"`
	ErrCode  string `json:"errCode"`
	ErrText  string `json:"errText"`
	Stack    string `json:"-"`
}

func (g *Error) HTTPCode() int {
	return g.HttpCode
}

func (g *Error) GetErrCode() string {
	return g.ErrCode
}

func (g *Error) Text() string {
	return g.ErrText
}

func (g *Error) StackTrace() string {
	return g.Stack
}

func (g *Error) Equal(err error) bool {
	if iError, ok := err.(IError); ok {
		return iError.GetErrCode() == g.ErrCode
	}

	return false
}

func (g *Error) Error() string {
	return g.ErrText
}

func (g *Error) New(ctx context.Context) *Error {
	return New(ctx, g)
}

func (g *Error) NewWithMsg(ctx context.Context, text string) *Error {
	return NewWithMsg(ctx, text, g)
}

func New(ctx context.Context, err IError) *Error {
	return causeWithMsg(ctx, "", err, true)
}

func NewWithMsg(ctx context.Context, text string, err IError) *Error {
	return causeWithMsg(ctx, text, err, true)
}

func causeWithMsg(ctx context.Context, text string, err IError, needLog bool) *Error {
	if len(text) == 0 {
		if len(err.Text()) == 0 {
			text = DefaultErrorText
		} else {
			text = err.Text()
		}
	}

	stack := getStackTrace()

	if needLog {
		log.Println(stack)
	}

	return &Error{
		HttpCode: err.HTTPCode(),
		ErrCode:  err.GetErrCode(),
		ErrText:  text,
		Stack:    stack,
	}
}

func getStackTrace() string {
	outputStr := ""
	_, fileName, fileLine, done := runtime.Caller(2)
	for i := 3; done; i++ {
		outputStr += fmt.Sprintf("%s:%d\n", fileName, fileLine)
		_, fileName, fileLine, done = runtime.Caller(i)
	}
	return outputStr
}
