package model

import "time"

type RefreshToken struct {
	Id        int64
	UserId    int64
	Token     string
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}
