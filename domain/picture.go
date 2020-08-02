package domain

import (
	"time"
	"image"
)

type Picture struct {
	ID int64 `json:"id"`
	Format string `json:"format"`
	Image image.Image `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}