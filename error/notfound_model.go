package error

import "fmt"

type NotfoundModel struct {
	Message      string
	AddedMessage string
}

func NewNotfoundModel(msg string) *NotfoundModel {
	return &NotfoundModel{
		Message:      "not_found model",
		AddedMessage: msg,
	}
}

func (e NotfoundModel) Error() string {
	return fmt.Sprintf("%s : %s", e.Message, e.AddedMessage)
}
