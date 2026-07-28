package model

import "time"

type Workspace struct {
	ID          int64
	Name        string
	Description string
	OwnerID     int64
	CreatedAt   time.Time
	Version     int64
}
