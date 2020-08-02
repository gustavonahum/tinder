package domain

import (
	"context"
	"time"
)

type User struct {
	ID            int64     `json:"id"`
	Profile	Profile `json:"profile"`
	LikesReceived []Like    `json:"likes_received"`
	Matches []Match    `json:"matches"`
	CreatedAt time.Time `json:"created_at"`
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
