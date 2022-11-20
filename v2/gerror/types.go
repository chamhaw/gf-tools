package gerror

import "encoding/json"

type BizError interface {
	// Code returns the integer number of current error code.
	Code() LocalCode

	// Message returns the brief message for current error code.
	Message() string

	// Detail returns the detailed information of current error code,
	// which is mainly designed as an extension field for error code.
	Detail() interface{}
}

type LocalBizError struct {
	message string
	code    LocalCode
	detail  interface{}
}

func NewBizError(code LocalCode, message string) *LocalBizError {
	return &LocalBizError{
		message: message,
		code:    code,
	}
}

func NewBizErrorWithDetail(code LocalCode, message string, detail interface{}) *LocalBizError {
	return &LocalBizError{
		message: message,
		code:    code,
		detail:  detail,
	}
}

func (err *LocalBizError) Code() LocalCode {
	return err.code
}

func (err *LocalBizError) Message() string {
	return err.message
}

func (err *LocalBizError) Detail() interface{} {
	return err.detail
}

// Error implements the interface of Error, it returns all the error as string.
func (err *LocalBizError) Error() string {
	b, e := json.Marshal(map[string]any{
		"code":    err.Code(),
		"message": err.Message(),
		"detail":  err.Detail(),
	})
	if e != nil {
		return "Cannot marshal"
	}
	return string(b)
}
