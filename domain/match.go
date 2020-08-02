package domain

import (
	"context"
	"time"
)

type Match struct {
	ID        int64     `json:"id"`
	IDUserMatched   int64     `json:"id_user_matched"`
	CreatedAt time.Time `json:"created_at"`
	ChatMessages	[]ChatMessage `json:"chat_messages"`
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
