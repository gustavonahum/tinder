package domain

import (
	"time"
)

type Profile struct {
	Name string `json:"name"`
	Birthday time.Time `json:"birthday"`
	Bio string `json:"bio"`
	Pictures []Picture `json:"pictures"`
}