package model

import "time"

type Task struct {
	Id             int64
	ProjectId      int64
	Titel          string
	Description    string
	Parent_task_id int64
	Status         string
	Priority       string
	AssigneeId     int64
	CreatedBy      int64
	Due_date       time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
