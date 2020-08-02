package domain

import (
	"context"
	"time"
)

type Like struct {
	ID        int64     `json:"id"`
	IDLiker   int64     `json:"id_liker"`
	IsSuper	bool	`json:"is_super"`
	CreatedAt time.Time `json:"created_at"`
}

type LikeUsecase interface {
	GetByID(ctx context.Context, id int64) (Like, error)
	Store(ctx context.Context, like *Like) error
	Delete(ctx context.Context, id int64) error
}

type LikeRepository interface {
	GetByID(ctx context.Context, id int64) (Like, error)
	Store(ctx context.Context, like *Like) error
	Delete(ctx context.Context, id int64) error
}
