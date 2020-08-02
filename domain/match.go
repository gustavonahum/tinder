package domain

import (
	"context"
	"time"
)

type Match struct {
	ID        int64     `json:"id"`
	IDUser1   int64     `json:"id_user_1"`
	IDUser2   int64     `json:"id_user_2"`
	CreatedAt time.Time `json:"created_at"`
}

type MatchUsecase interface {
	GetByID(ctx context.Context, id int64) (Match, error)
	Store(ctx context.Context, match *Match) error
	Delete(ctx context.Context, id int64) error
}

type MatchRepository interface {
	GetByID(ctx context.Context, id int64) (Match, error)
	Store(ctx context.Context, match *Match) error
	Delete(ctx context.Context, id int64) error
}
