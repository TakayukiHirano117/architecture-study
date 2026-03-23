package customerr

import (
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
)

type appErr struct {
	trace error
	msg   string
	code  int
}

type AppError interface {
	Code() int
	Msg() string
	Trace() error
	Error() string
}

func (e *appErr) Code() int {
	return e.code
}

func (e *appErr) Msg() string {
	return e.msg
}

func (e *appErr) Error() string {
	return e.msg
}

func (e *appErr) Trace() error {
	return e.trace
}

// TODO: エラーの種類に応じてメソッドとtypeを追加する
// 401, 403はログインを実装したら追加する

type BadRequestErr struct {
	*appErr
}

type UnauthorizedErr struct {
	*appErr
}

type ForbiddenErr struct {
	*appErr
}

type NotFoundErr struct {
	*appErr
}

type ConflictErr struct {
	*appErr
}

type InternalErr struct {
	*appErr
}

// 400番台
// 400
func BadRequest(msg string) *BadRequestErr {
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func BadRequestf(format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)

	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}

func BadRequestWrapf(err2 error, format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)

	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

// 401
func Unauthorized(msg string) *UnauthorizedErr {
	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Unauthorizedf(format string, msg ...any) *UnauthorizedErr {
	message := fmt.Sprintf(format, msg...)

	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}

func UnauthorizedWrapf(err2 error, format string, msg ...any) *UnauthorizedErr {
	message := fmt.Sprintf(format, msg...)

	return &UnauthorizedErr{
		&appErr{
			code:  http.StatusUnauthorized,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

// 403

func Forbidden(msg string) *ForbiddenErr {
	return &ForbiddenErr{
		&appErr{
			code:  http.StatusForbidden,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Forbiddenf(format string, msg ...any) *ForbiddenErr {
	message := fmt.Sprintf(format, msg...)
	return &ForbiddenErr{
		&appErr{
			code:  http.StatusForbidden,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}

func ForbiddenWrapf(err2 error, format string, msg ...any) *ForbiddenErr {
	message := fmt.Sprintf(format, msg...)
	return &ForbiddenErr{
		&appErr{
			code:  http.StatusForbidden,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

// 404
func NotFound(msg string) *NotFoundErr {
	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func NotFoundf(format string, msg ...any) *NotFoundErr {
	message := fmt.Sprintf(format, msg...)

	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}

func NotFoundWrapf(err2 error, format string, msg ...any) *NotFoundErr {
	message := fmt.Sprintf(format, msg...)

	return &NotFoundErr{
		&appErr{
			code:  http.StatusNotFound,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

// 409
func Conflict(msg string) *ConflictErr {
	return &ConflictErr{
		&appErr{
			code:  http.StatusConflict,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Conflictf(format string, msg ...any) *ConflictErr {
	message := fmt.Sprintf(format, msg...)

	return &ConflictErr{
		&appErr{
			code:  http.StatusConflict,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}

func ConflictWrapf(err2 error, format string, msg ...any) *ConflictErr {
	message := fmt.Sprintf(format, msg...)

	return &ConflictErr{
		&appErr{
			code:  http.StatusConflict,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}

// 500番台
func Internal(msg string) *InternalErr {
	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   msg,
			trace: errors.New(msg),
		},
	}
}

func Internalf(format string, msg ...any) *InternalErr {
	message := fmt.Sprintf(format, msg...)

	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   message,
			trace: errors.Errorf(format, msg...),
		},
	}
}

func InternalWrapf(err2 error, format string, msg ...any) *InternalErr {
	message := fmt.Sprintf(format, msg...)

	return &InternalErr{
		&appErr{
			code:  http.StatusInternalServerError,
			msg:   message,
			trace: errors.Wrapf(err2, format, msg...),
		},
	}
}
