package input

import "github.com/vsanna/go_web/domain/model"

type Register struct {
	User     *model.User
	Name     string
	Email    string
	Password string
}
