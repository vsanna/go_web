package usecase

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vsanna/go_web/domain/repository"
	"github.com/vsanna/go_web/usecase/output"
)

type ProfileUsecase struct {
	repo repository.User
}

func NewProfileUsecase(repo repository.User) *ProfileUsecase {
	return &ProfileUsecase{
		repo: repo,
	}
}

func (p *ProfileUsecase) GetProfile(ctx context.Context, id int) (*output.Profile, error) {
	u, err := p.repo.FindById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get user")
	}

	profile := &output.Profile{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
	return profile, nil
}
