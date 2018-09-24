package usecase

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vsanna/go_web/domain/repository"
	"github.com/vsanna/go_web/usecase/output"
)

type JsonExample struct {
	repo repository.User
}

func NewJsonExample(repo repository.User) *JsonExample {
	return &JsonExample{
		repo: repo,
	}
}

func (j *JsonExample) Fetch(ctx context.Context) (output.Users, error) {
	users, err := j.repo.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get user")
	}

	var data output.Users
	for _, u := range users {
		data = append(data, output.User{
			ID:   u.ID,
			Name: u.Name,
		})
	}
	return data, nil
}
