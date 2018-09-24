package error

import "fmt"

type InvalidParamsError struct {
	Message      string
	AddedMessage string
}

func NewInvalidParamsError(msg string) *InvalidParamsError {
	return &InvalidParamsError{
		Message:      "invalid parameter",
		AddedMessage: msg,
	}
}

func (e InvalidParamsError) Error() string {
	return fmt.Sprintf("%s : %s", e.Message, e.AddedMessage)
}
