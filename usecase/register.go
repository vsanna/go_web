package usecase

import (
	"context"
	"fmt"

	"github.com/vsanna/go_web/domain/model"
	"github.com/vsanna/go_web/domain/repository"
	customerr "github.com/vsanna/go_web/error"
	"github.com/vsanna/go_web/usecase/input"
)

type RegisterUsecase struct {
	repo repository.User
}

func NewRegisterUsecase(repo repository.User) *RegisterUsecase {
	return &RegisterUsecase{
		repo: repo,
	}
}

func (u *RegisterUsecase) Register(ctx context.Context, input *input.Register) (*model.User, error) {
	email := input.Email
	password := input.Password
	name := input.Name

	if email == "" || password == "" || name == "" {
		return nil, customerr.NewInvalidParamsError(fmt.Sprintf("some of name, password, email are blank. name: %s, email: %s, len(password): %d", name, email, len(password)))
	}

	user, err := model.NewUser(name, email, password)
	if err != nil {
		return nil, customerr.NewInvalidModelError(fmt.Sprintf("failed to setup user. error: %s", err.Error()))
	}

	// NOTE 手を抜いてerrorチェック省略
	err = u.repo.Create(ctx, user)
	if err != nil {
		return nil, customerr.NewInvalidModelError(fmt.Sprintf("failed to create user. error: %s", err.Error()))
	}

	return user, nil
}
