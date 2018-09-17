package repository

import (
	"context"

	"github.com/vsanna/go_web/domain/model"
)

type Post interface {
	Find(ctx context.Context, ID int) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
	// Where
	// Create
	// Update
	// Destroy
}
