package repository

import (
	"context"
	"github.com/danangsw/go-restapi-mysql/model"
)

// PostRepo explain...
type PostRepo interface {
	Fetch(ctx context.Context, num int64) ([]*model.Post, error)
	GetByID(ctx context.Context, id int64) (*model.Post, error)
	Create(ctx context.Context, p *model.Post) (int64, error)
	Update(ctx context.Context, p *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

// UserRepo explain...
type UserRepo interface {
	Fetch(ctx context.Context, num int64) ([]*model.Post, error)
	GetByID(ctx context.Context, id int64) (*model.Post, error)
	Create(ctx context.Context, p *model.Post) (int64, error)
	Update(ctx context.Context, p *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
}