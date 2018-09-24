package usecase

import (
	"context"
	"fmt"

	"github.com/vsanna/go_web/domain/repository"
	customerr "github.com/vsanna/go_web/error"
	"github.com/vsanna/go_web/lib"
)

type SigninUsecase struct {
	repo repository.User
}

func NewSigninUsecase(repo repository.User) *SigninUsecase {
	return &SigninUsecase{
		repo: repo,
	}
}

func (u *SigninUsecase) Signin(ctx context.Context, email, password string) error {
	if email == "" || password == "" {
		return customerr.NewInvalidParamsError(fmt.Sprintf("name: %s, len(password): %d", email, len(password)))
	}

	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return customerr.NewNotfoundModel(fmt.Sprintf("query: email=%s", email))
	}

	if lib.Compare(user.EncryptedPassword, password) {
		return nil
	} else {
		return customerr.NewInvalidParamsError("password is invalid")
	}
}
