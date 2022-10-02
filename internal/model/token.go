package model

import "time"

type Blacklist struct {
	Email     string
	Token     string
	CreatedAt time.Time
}
