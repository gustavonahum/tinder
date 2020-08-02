package domain

import (
	"context"
	"time"
)

type User struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Birthday      time.Time `json:"birthday"`
	LikesReceived []Like    `json:"likes_received"`
	Matches []Match    `json:"matches"`
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (User, error)
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
