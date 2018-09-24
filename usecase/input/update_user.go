package input

import "github.com/vsanna/go_web/domain/model"

type UpdateUser struct {
	User     *model.User
	Name     string
	Email    string
	Password string
}
