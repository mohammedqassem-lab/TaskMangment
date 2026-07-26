package model

import "time"

type Project struct {
	Id          int64
	WorkspaceId int64
	Name        string
	Description string
	CreatedBy   int64
	createdAt   time.Time
	updatedAt   time.Time
}
