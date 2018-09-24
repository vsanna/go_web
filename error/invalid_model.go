package error

import "fmt"

type InvalidModelError struct {
	Message      string
	AddedMessage string
}

func NewInvalidModelError(msg string) *InvalidModelError {
	return &InvalidModelError{
		Message:      "invalid model",
		AddedMessage: msg,
	}
}

func (e InvalidModelError) Error() string {
	return fmt.Sprintf("%s : %s", e.Message, e.AddedMessage)
}
